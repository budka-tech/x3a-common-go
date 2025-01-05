package port

import (
	"fmt"
	"github.com/budka-tech/envo"
)

type Port = uint16

const (
	Redis      Port = 6379
	UsersGrpc       = 10100
	CCGrpc          = 10200
	NOGrpc          = 10300
	NewsGrpc        = 10400
	SearchGrpc      = 10500
	S3Grpc          = 10600
)

type ServiceName string

const (
	ServiceUsers  ServiceName = "x3a-users-service"
	ServiceRedis              = "x3a-redis-service"
	ServiceCC                 = "x3a-cc-service"
	ServiceNO                 = "x3a-no-service"
	ServiceNews               = "x3a-news-service"
	ServiceSearch             = "x3a-search-service"
	ServiceS3G                = "x3a-s3-service"
)

var m = map[ServiceName]Port{
	ServiceUsers:  UsersGrpc,
	ServiceRedis:  Redis,
	ServiceCC:     CCGrpc,
	ServiceNO:     NOGrpc,
	ServiceNews:   NewsGrpc,
	ServiceSearch: SearchGrpc,
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
