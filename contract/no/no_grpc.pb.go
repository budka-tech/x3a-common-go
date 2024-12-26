// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: no/no.proto

package nov1

import (
	context "context"
	common "github.com/budka-tech/snip-common-go/contract/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NO_UploadArticle_FullMethodName = "/assist.NO/UploadArticle"
)

// NOClient is the client API for NO service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NOClient interface {
	// Получение монет(ы)
	UploadArticle(ctx context.Context, in *UploadArticleParams, opts ...grpc.CallOption) (*common.Response, error)
}

type nOClient struct {
	cc grpc.ClientConnInterface
}

func NewNOClient(cc grpc.ClientConnInterface) NOClient {
	return &nOClient{cc}
}

func (c *nOClient) UploadArticle(ctx context.Context, in *UploadArticleParams, opts ...grpc.CallOption) (*common.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Response)
	err := c.cc.Invoke(ctx, NO_UploadArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NOServer is the server API for NO service.
// All implementations must embed UnimplementedNOServer
// for forward compatibility.
type NOServer interface {
	// Получение монет(ы)
	UploadArticle(context.Context, *UploadArticleParams) (*common.Response, error)
	mustEmbedUnimplementedNOServer()
}

// UnimplementedNOServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNOServer struct{}

func (UnimplementedNOServer) UploadArticle(context.Context, *UploadArticleParams) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadArticle not implemented")
}
func (UnimplementedNOServer) mustEmbedUnimplementedNOServer() {}
func (UnimplementedNOServer) testEmbeddedByValue()            {}

// UnsafeNOServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NOServer will
// result in compilation errors.
type UnsafeNOServer interface {
	mustEmbedUnimplementedNOServer()
}

func RegisterNOServer(s grpc.ServiceRegistrar, srv NOServer) {
	// If the following call pancis, it indicates UnimplementedNOServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NO_ServiceDesc, srv)
}

func _NO_UploadArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadArticleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NOServer).UploadArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NO_UploadArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NOServer).UploadArticle(ctx, req.(*UploadArticleParams))
	}
	return interceptor(ctx, in, info, handler)
}

// NO_ServiceDesc is the grpc.ServiceDesc for NO service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NO_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assist.NO",
	HandlerType: (*NOServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadArticle",
			Handler:    _NO_UploadArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "no/no.proto",
}
