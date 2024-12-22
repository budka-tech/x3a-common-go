package rdb

import (
	"context"
	"fmt"
	"github.com/budka-tech/envo"
	"github.com/budka-tech/logit-go"
	"github.com/budka-tech/snip-common-go/port"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	defaultConnectTimeout = 5 * time.Second
	defaultPingTimeout    = 2 * time.Second
)

type Rdb struct {
	redis.Client
	intervalMonitorConnection *time.Duration
	serviceName               port.ServiceName
	stopMonitor               chan struct{}
	env                       *envo.Env
	logger                    logit.Logger
}

type Params struct {
	ServiceName               port.ServiceName
	IntervalMonitorConnection *time.Duration
	Env                       *envo.Env
	Logger                    logit.Logger
}

func NewRdb(params Params) *Rdb {
	var intervalMonitorConnection *time.Duration = params.IntervalMonitorConnection
	if intervalMonitorConnection == nil {
		d := time.Duration(15 * time.Second)
		intervalMonitorConnection = &d
	}

	return &Rdb{
		intervalMonitorConnection: intervalMonitorConnection,
		serviceName:               params.ServiceName,
		env:                       params.Env,
		logger:                    params.Logger,
	}
}

func (r *Rdb) Init(ctx context.Context) error {
	const op = "rdb.Init"
	ctx = r.logger.NewOpCtx(ctx, op)

	err := r.connect(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	r.stopMonitor = make(chan struct{})
	go r.monitorConnection(ctx)

	return nil
}

func (r *Rdb) connect(ctx context.Context) error {
	const op = "rdb.connect"
	ctx = r.logger.NewOpCtx(ctx, op)

	addr := port.FormatServiceTCP(r.env, r.serviceName)
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	connectCtx, cancel := context.WithTimeout(ctx, defaultConnectTimeout)
	defer cancel()

	err := client.Ping(connectCtx).Err()
	if err != nil {
		return fmt.Errorf("%s: не удалось подключиться к Redis: %w", op, err)
	}

	r.Client = *client
	r.logger.Info(ctx, "успешное подключение к Redis")
	return nil
}

func (r *Rdb) Ping(ctx context.Context) error {
	const op = "rdb.Ping"
	ctx = r.logger.NewOpCtx(ctx, op)

	pingCtx, cancel := context.WithTimeout(ctx, defaultPingTimeout)
	defer cancel()

	_, err := r.Client.Ping(pingCtx).Result()
	if err != nil {
		return fmt.Errorf("%s: не удалось выполнить пинг Redis: %w", op, err)
	}

	return nil
}

func (r *Rdb) monitorConnection(ctx context.Context) {
	ticker := time.NewTicker(*r.intervalMonitorConnection)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			r.logger.Info(ctx, "мониторинг подключения к Redis остановлен")
			return
		case <-r.stopMonitor:
			r.logger.Info(ctx, "мониторинг подключения к Redis остановлен")
			return
		case <-ticker.C:
			err := r.Ping(ctx)
			if err != nil {
				r.logger.Error(ctx, fmt.Errorf("ошибка подключения к Redis: %v", err))
				if reconnectErr := r.connect(ctx); reconnectErr != nil {
					r.logger.Error(ctx, fmt.Errorf("не удалось переподключиться к Redis: %v", reconnectErr))
				} else {
					r.logger.Info(ctx, "успешное переподключение к Redis")
				}
			}
		}
	}
}

func (r *Rdb) Dispose() error {
	if r.stopMonitor != nil {
		close(r.stopMonitor)
	}
	return r.Client.Close()
}
