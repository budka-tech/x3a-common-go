// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: news/news.proto

package newsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	News_Init_FullMethodName              = "/assist.News/Init"
	News_GetArticles_FullMethodName       = "/assist.News/GetArticles"
	News_GetArticle_FullMethodName        = "/assist.News/GetArticle"
	News_GetArticleContent_FullMethodName = "/assist.News/GetArticleContent"
	News_GetFilters_FullMethodName        = "/assist.News/GetFilters"
	News_CreateFilter_FullMethodName      = "/assist.News/CreateFilter"
	News_UpdateFilter_FullMethodName      = "/assist.News/UpdateFilter"
	News_DeleteFilter_FullMethodName      = "/assist.News/DeleteFilter"
	News_GetTags_FullMethodName           = "/assist.News/GetTags"
	News_GetSources_FullMethodName        = "/assist.News/GetSources"
	News_GetFavorites_FullMethodName      = "/assist.News/GetFavorites"
	News_SetFavorite_FullMethodName       = "/assist.News/SetFavorite"
)

// NewsClient is the client API for News service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsClient interface {
	Init(ctx context.Context, in *InitParams, opts ...grpc.CallOption) (*InitResponse, error)
	GetArticles(ctx context.Context, in *GetArticlesParams, opts ...grpc.CallOption) (*GetArticlesResponse, error)
	GetArticle(ctx context.Context, in *GetArticleParams, opts ...grpc.CallOption) (*GetArticleResponse, error)
	GetArticleContent(ctx context.Context, in *GetArticleParams, opts ...grpc.CallOption) (*GetArticleContentResponse, error)
	GetFilters(ctx context.Context, in *GetFiltersParams, opts ...grpc.CallOption) (*GetFiltersResponse, error)
	CreateFilter(ctx context.Context, in *CreateFilterParams, opts ...grpc.CallOption) (*ManipulationFilterResponse, error)
	UpdateFilter(ctx context.Context, in *UpdateFilterParams, opts ...grpc.CallOption) (*ManipulationFilterResponse, error)
	DeleteFilter(ctx context.Context, in *DeleteFilterParams, opts ...grpc.CallOption) (*ManipulationFilterResponse, error)
	GetTags(ctx context.Context, in *GetTagsParams, opts ...grpc.CallOption) (*GetTagsResponse, error)
	GetSources(ctx context.Context, in *GetSourcesParams, opts ...grpc.CallOption) (*GetSourcesResponse, error)
	GetFavorites(ctx context.Context, in *GetFavoritesParams, opts ...grpc.CallOption) (*GetFavoritesResponse, error)
	SetFavorite(ctx context.Context, in *SetFavoriteParams, opts ...grpc.CallOption) (*SetFavoriteResponse, error)
}

type newsClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsClient(cc grpc.ClientConnInterface) NewsClient {
	return &newsClient{cc}
}

func (c *newsClient) Init(ctx context.Context, in *InitParams, opts ...grpc.CallOption) (*InitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, News_Init_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetArticles(ctx context.Context, in *GetArticlesParams, opts ...grpc.CallOption) (*GetArticlesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticlesResponse)
	err := c.cc.Invoke(ctx, News_GetArticles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetArticle(ctx context.Context, in *GetArticleParams, opts ...grpc.CallOption) (*GetArticleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticleResponse)
	err := c.cc.Invoke(ctx, News_GetArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetArticleContent(ctx context.Context, in *GetArticleParams, opts ...grpc.CallOption) (*GetArticleContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticleContentResponse)
	err := c.cc.Invoke(ctx, News_GetArticleContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetFilters(ctx context.Context, in *GetFiltersParams, opts ...grpc.CallOption) (*GetFiltersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFiltersResponse)
	err := c.cc.Invoke(ctx, News_GetFilters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) CreateFilter(ctx context.Context, in *CreateFilterParams, opts ...grpc.CallOption) (*ManipulationFilterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ManipulationFilterResponse)
	err := c.cc.Invoke(ctx, News_CreateFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) UpdateFilter(ctx context.Context, in *UpdateFilterParams, opts ...grpc.CallOption) (*ManipulationFilterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ManipulationFilterResponse)
	err := c.cc.Invoke(ctx, News_UpdateFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) DeleteFilter(ctx context.Context, in *DeleteFilterParams, opts ...grpc.CallOption) (*ManipulationFilterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ManipulationFilterResponse)
	err := c.cc.Invoke(ctx, News_DeleteFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetTags(ctx context.Context, in *GetTagsParams, opts ...grpc.CallOption) (*GetTagsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTagsResponse)
	err := c.cc.Invoke(ctx, News_GetTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetSources(ctx context.Context, in *GetSourcesParams, opts ...grpc.CallOption) (*GetSourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSourcesResponse)
	err := c.cc.Invoke(ctx, News_GetSources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetFavorites(ctx context.Context, in *GetFavoritesParams, opts ...grpc.CallOption) (*GetFavoritesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFavoritesResponse)
	err := c.cc.Invoke(ctx, News_GetFavorites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) SetFavorite(ctx context.Context, in *SetFavoriteParams, opts ...grpc.CallOption) (*SetFavoriteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetFavoriteResponse)
	err := c.cc.Invoke(ctx, News_SetFavorite_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServer is the server API for News service.
// All implementations must embed UnimplementedNewsServer
// for forward compatibility.
type NewsServer interface {
	Init(context.Context, *InitParams) (*InitResponse, error)
	GetArticles(context.Context, *GetArticlesParams) (*GetArticlesResponse, error)
	GetArticle(context.Context, *GetArticleParams) (*GetArticleResponse, error)
	GetArticleContent(context.Context, *GetArticleParams) (*GetArticleContentResponse, error)
	GetFilters(context.Context, *GetFiltersParams) (*GetFiltersResponse, error)
	CreateFilter(context.Context, *CreateFilterParams) (*ManipulationFilterResponse, error)
	UpdateFilter(context.Context, *UpdateFilterParams) (*ManipulationFilterResponse, error)
	DeleteFilter(context.Context, *DeleteFilterParams) (*ManipulationFilterResponse, error)
	GetTags(context.Context, *GetTagsParams) (*GetTagsResponse, error)
	GetSources(context.Context, *GetSourcesParams) (*GetSourcesResponse, error)
	GetFavorites(context.Context, *GetFavoritesParams) (*GetFavoritesResponse, error)
	SetFavorite(context.Context, *SetFavoriteParams) (*SetFavoriteResponse, error)
	mustEmbedUnimplementedNewsServer()
}

// UnimplementedNewsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNewsServer struct{}

func (UnimplementedNewsServer) Init(context.Context, *InitParams) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedNewsServer) GetArticles(context.Context, *GetArticlesParams) (*GetArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedNewsServer) GetArticle(context.Context, *GetArticleParams) (*GetArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedNewsServer) GetArticleContent(context.Context, *GetArticleParams) (*GetArticleContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleContent not implemented")
}
func (UnimplementedNewsServer) GetFilters(context.Context, *GetFiltersParams) (*GetFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilters not implemented")
}
func (UnimplementedNewsServer) CreateFilter(context.Context, *CreateFilterParams) (*ManipulationFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFilter not implemented")
}
func (UnimplementedNewsServer) UpdateFilter(context.Context, *UpdateFilterParams) (*ManipulationFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFilter not implemented")
}
func (UnimplementedNewsServer) DeleteFilter(context.Context, *DeleteFilterParams) (*ManipulationFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFilter not implemented")
}
func (UnimplementedNewsServer) GetTags(context.Context, *GetTagsParams) (*GetTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (UnimplementedNewsServer) GetSources(context.Context, *GetSourcesParams) (*GetSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSources not implemented")
}
func (UnimplementedNewsServer) GetFavorites(context.Context, *GetFavoritesParams) (*GetFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavorites not implemented")
}
func (UnimplementedNewsServer) SetFavorite(context.Context, *SetFavoriteParams) (*SetFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFavorite not implemented")
}
func (UnimplementedNewsServer) mustEmbedUnimplementedNewsServer() {}
func (UnimplementedNewsServer) testEmbeddedByValue()              {}

// UnsafeNewsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsServer will
// result in compilation errors.
type UnsafeNewsServer interface {
	mustEmbedUnimplementedNewsServer()
}

func RegisterNewsServer(s grpc.ServiceRegistrar, srv NewsServer) {
	// If the following call pancis, it indicates UnimplementedNewsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&News_ServiceDesc, srv)
}

func _News_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).Init(ctx, req.(*InitParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticles(ctx, req.(*GetArticlesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticle(ctx, req.(*GetArticleParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetArticleContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticleContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticleContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticleContent(ctx, req.(*GetArticleParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFiltersParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetFilters(ctx, req.(*GetFiltersParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_CreateFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFilterParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).CreateFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_CreateFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).CreateFilter(ctx, req.(*CreateFilterParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_UpdateFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFilterParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).UpdateFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_UpdateFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).UpdateFilter(ctx, req.(*UpdateFilterParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_DeleteFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFilterParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).DeleteFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_DeleteFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).DeleteFilter(ctx, req.(*DeleteFilterParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetTags(ctx, req.(*GetTagsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSourcesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetSources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetSources(ctx, req.(*GetSourcesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoritesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetFavorites(ctx, req.(*GetFavoritesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_SetFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFavoriteParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).SetFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_SetFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).SetFavorite(ctx, req.(*SetFavoriteParams))
	}
	return interceptor(ctx, in, info, handler)
}

// News_ServiceDesc is the grpc.ServiceDesc for News service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var News_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assist.News",
	HandlerType: (*NewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _News_Init_Handler,
		},
		{
			MethodName: "GetArticles",
			Handler:    _News_GetArticles_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _News_GetArticle_Handler,
		},
		{
			MethodName: "GetArticleContent",
			Handler:    _News_GetArticleContent_Handler,
		},
		{
			MethodName: "GetFilters",
			Handler:    _News_GetFilters_Handler,
		},
		{
			MethodName: "CreateFilter",
			Handler:    _News_CreateFilter_Handler,
		},
		{
			MethodName: "UpdateFilter",
			Handler:    _News_UpdateFilter_Handler,
		},
		{
			MethodName: "DeleteFilter",
			Handler:    _News_DeleteFilter_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _News_GetTags_Handler,
		},
		{
			MethodName: "GetSources",
			Handler:    _News_GetSources_Handler,
		},
		{
			MethodName: "GetFavorites",
			Handler:    _News_GetFavorites_Handler,
		},
		{
			MethodName: "SetFavorite",
			Handler:    _News_SetFavorite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news/news.proto",
}
