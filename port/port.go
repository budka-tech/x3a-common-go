package port

import (
	"fmt"
	"github.com/budka-tech/envo"
)

type Port = uint16

const (
	Redis      Port = 6379
	UsersGrpc       = 10100
	CCGrpc       = 10100
	S3Grpc          = 10500
)

type ServiceName string

const (
	ServiceUsers  ServiceName = "x3a-users-service"
	ServiceRedis              = "x3a-redis-service"
	ServiceGsm                = "x3a-gsm-service"
	ServicePush               = "x3a-push-service"
	ServiceAssist             = "x3a-assist-service"
	ServiceS3G                = "x3a-s3-service"
)

var m = map[ServiceName]Port{
	ServiceUsers:  UsersGrpc,
	ServiceRedis:  Redis,
	ServiceGsm:    GsmGrpc,
	ServicePush:   PushGrpc,
	ServiceAssist: AssistGrpc,
	ServiceS3G:    S3Grpc,
}

func Format(port Port) string {
	return fmt.Sprintf("%d", port)
}

func FormatTCP(port Port) string {
	return fmt.Sprintf(":%d", port)
}

func FormatServiceTCP(env *envo.Env, name ServiceName) string {
	port := m[name]
	if env.IsLocal() {
		return FormatLocal(port)
	} else {
		return fmt.Sprintf("%v:%d", name, port)
	}
}

func FormatLocal(port Port) string {
	return fmt.Sprintf("localhost:%d", port)
}
