package status

import (
	"google.golang.org/grpc/codes"
	gstatus "google.golang.org/grpc/status"
)

var mGrpc = map[Status]codes.Code{
	OK:                  codes.OK,
	Any:                 codes.Unknown,
	NotFound:            codes.NotFound,
	AlreadyExists:       codes.AlreadyExists,
	PermissionDenied:    codes.PermissionDenied,
	TooFrequentRequests: codes.ResourceExhausted,
	InternalError:       codes.Internal,
	ManyConnections:     codes.ResourceExhausted,
	NotEnoughArguments:  codes.InvalidArgument,
	NotAuthorized:       codes.Unauthenticated,
	Authorized:          codes.Unknown,
	IncorrectValue:      codes.Unknown,
	Inactivity:          codes.Unknown,
	Timeout:             codes.DeadlineExceeded,
	ResourceUnavailable: codes.Unavailable,
	OperationFailed:     codes.Internal,
	NotImplemented:      codes.Unimplemented,
}

func GrpcError(st Status) error {
	if val, ok := mGrpc[st]; ok {
		return gstatus.Error(val, Readable(st))
	}

	return gstatus.Error(mGrpc[Any], Readable(st))
}

func Proto(st Status) uint32 {
	return uint32(st)
}

func FromProto(st uint32) Status {
	return Status(st)
}
