package gclient

import (
	"context"
	"github.com/budka-tech/envo"
	"github.com/budka-tech/logit-go"
	"github.com/budka-tech/snip-common-go/port"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"net"
	"time"
)

type GClient[T any] struct {
	name          string
	serviceName   port.ServiceName
	listener      net.Listener
	Cli           T
	conn          *grpc.ClientConn
	clientFactory func(*grpc.ClientConn) T
	dialOptions   []grpc.DialOption
	env           *envo.Env
	logger        logit.Logger
}

type GClientParams[T any] struct {
	Name          string
	ServiceName   port.ServiceName
	ClientFactory func(*grpc.ClientConn) T
	Env           *envo.Env
	Logger        logit.Logger
	DialOptions   []grpc.DialOption
}

// NewGClient creates a new GClient instance
func NewGClient[T any](params *GClientParams[T]) *GClient[T] {
	return &GClient[T]{
		name:          params.Name,
		serviceName:   params.ServiceName,
		clientFactory: params.ClientFactory,
		env:           params.Env,
		logger:        params.Logger,
		dialOptions:   append([]grpc.DialOption{grpc.WithInsecure()}, params.DialOptions...),
	}
}

func (n *GClient[T]) Init(ctx context.Context) error {
	backoff := time.Second
	maxBackoff := time.Minute
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			host := port.FormatServiceTCP(n.env, n.serviceName)
			n.logger.Infof(ctx, "Попытка подключения к сервису %s по адресу %s", n.name, host)

			conn, err := grpc.DialContext(ctx, host, n.dialOptions...)
			if err != nil {
				n.logger.Errorf(ctx, "Не удалось подключиться к %s: %v. Повторная попытка через %v...", n.name, err, backoff)
				time.Sleep(backoff)
				backoff = min(backoff*2, maxBackoff)
				continue
			}

			n.conn = conn
			n.Cli = n.clientFactory(conn)

			n.logger.Infof(ctx, "Успешное подключение к сервису %s по адресу %s", n.name, host)

			go n.monitorConnection(ctx)

			return nil
		}
	}
}

func (n *GClient[T]) monitorConnection(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			state := n.conn.GetState()
			if state == connectivity.TransientFailure || state == connectivity.Shutdown {
				n.logger.Errorf(ctx, "Соединение с сервисом %s потеряно. Попытка переподключения", n.name)
				if err := n.Init(ctx); err != nil {
					n.logger.Errorf(ctx, "Не удалось переподключиться к сервису %s: %v", n.name, err)
				}
				return
			}
		}
	}
}

func (n *GClient[T]) Close() error {
	if n.conn != nil {
		return n.conn.Close()
	}
	return nil
}

func min(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}
