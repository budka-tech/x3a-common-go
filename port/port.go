package port

import (
	"github.com/budka-tech/envo"
	"github.com/budka-tech/iport"
)

const (
	Redis      iport.Port = 6379
	UsersGrpc             = 10100
	CCGrpc                = 10200
	NOGrpc                = 10300
	NewsGrpc              = 10400
	SearchGrpc            = 10500
	S3Grpc                = 10600
)

const (
	ServiceUsers  iport.Host = "x3a-users-service"
	ServiceRedis             = "x3a-redis-service"
	ServiceCC                = "x3a-cc-service"
	ServiceNO                = "x3a-no-service"
	ServiceNews              = "x3a-news-service"
	ServiceSearch            = "x3a-search-service"
	ServiceS3G               = "x3a-s3-service"
)

var m = map[iport.Host]iport.Port{
	ServiceUsers:  UsersGrpc,
	ServiceRedis:  Redis,
	ServiceCC:     CCGrpc,
	ServiceNO:     NOGrpc,
	ServiceNews:   NewsGrpc,
	ServiceSearch: SearchGrpc,
	ServiceS3G:    S3Grpc,
}

func NewPorts(env *envo.Env) *iport.Ports {
	return iport.NewPorts(
		&iport.Params{
			Env: env,
			M:   m,
		},
	)
}
