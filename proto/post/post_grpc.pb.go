// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: proto/post/post.proto

package post

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
	LendingPostService_CreateLendingPost_FullMethodName    = "/LendingPostService/CreateLendingPost"
	LendingPostService_GetLendingPostDetail_FullMethodName = "/LendingPostService/GetLendingPostDetail"
	LendingPostService_SearchLendingPost_FullMethodName    = "/LendingPostService/SearchLendingPost"
)

// LendingPostServiceClient is the client API for LendingPostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// LendingPostService definition
type LendingPostServiceClient interface {
	CreateLendingPost(ctx context.Context, in *CreateLendingPostRequest, opts ...grpc.CallOption) (*CreateLendingPostResponse, error)
	GetLendingPostDetail(ctx context.Context, in *GetLendingPostDetailRequest, opts ...grpc.CallOption) (*LendingPost, error)
	SearchLendingPost(ctx context.Context, in *SearchLendingPostRequest, opts ...grpc.CallOption) (*LendingPostList, error)
}

type lendingPostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLendingPostServiceClient(cc grpc.ClientConnInterface) LendingPostServiceClient {
	return &lendingPostServiceClient{cc}
}

func (c *lendingPostServiceClient) CreateLendingPost(ctx context.Context, in *CreateLendingPostRequest, opts ...grpc.CallOption) (*CreateLendingPostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLendingPostResponse)
	err := c.cc.Invoke(ctx, LendingPostService_CreateLendingPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lendingPostServiceClient) GetLendingPostDetail(ctx context.Context, in *GetLendingPostDetailRequest, opts ...grpc.CallOption) (*LendingPost, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LendingPost)
	err := c.cc.Invoke(ctx, LendingPostService_GetLendingPostDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lendingPostServiceClient) SearchLendingPost(ctx context.Context, in *SearchLendingPostRequest, opts ...grpc.CallOption) (*LendingPostList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LendingPostList)
	err := c.cc.Invoke(ctx, LendingPostService_SearchLendingPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LendingPostServiceServer is the server API for LendingPostService service.
// All implementations must embed UnimplementedLendingPostServiceServer
// for forward compatibility.
//
// LendingPostService definition
type LendingPostServiceServer interface {
	CreateLendingPost(context.Context, *CreateLendingPostRequest) (*CreateLendingPostResponse, error)
	GetLendingPostDetail(context.Context, *GetLendingPostDetailRequest) (*LendingPost, error)
	SearchLendingPost(context.Context, *SearchLendingPostRequest) (*LendingPostList, error)
	mustEmbedUnimplementedLendingPostServiceServer()
}

// UnimplementedLendingPostServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLendingPostServiceServer struct{}

func (UnimplementedLendingPostServiceServer) CreateLendingPost(context.Context, *CreateLendingPostRequest) (*CreateLendingPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLendingPost not implemented")
}
func (UnimplementedLendingPostServiceServer) GetLendingPostDetail(context.Context, *GetLendingPostDetailRequest) (*LendingPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLendingPostDetail not implemented")
}
func (UnimplementedLendingPostServiceServer) SearchLendingPost(context.Context, *SearchLendingPostRequest) (*LendingPostList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchLendingPost not implemented")
}
func (UnimplementedLendingPostServiceServer) mustEmbedUnimplementedLendingPostServiceServer() {}
func (UnimplementedLendingPostServiceServer) testEmbeddedByValue()                            {}

// UnsafeLendingPostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LendingPostServiceServer will
// result in compilation errors.
type UnsafeLendingPostServiceServer interface {
	mustEmbedUnimplementedLendingPostServiceServer()
}

func RegisterLendingPostServiceServer(s grpc.ServiceRegistrar, srv LendingPostServiceServer) {
	// If the following call pancis, it indicates UnimplementedLendingPostServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LendingPostService_ServiceDesc, srv)
}

func _LendingPostService_CreateLendingPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLendingPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendingPostServiceServer).CreateLendingPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LendingPostService_CreateLendingPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendingPostServiceServer).CreateLendingPost(ctx, req.(*CreateLendingPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LendingPostService_GetLendingPostDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLendingPostDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendingPostServiceServer).GetLendingPostDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LendingPostService_GetLendingPostDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendingPostServiceServer).GetLendingPostDetail(ctx, req.(*GetLendingPostDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LendingPostService_SearchLendingPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchLendingPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendingPostServiceServer).SearchLendingPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LendingPostService_SearchLendingPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendingPostServiceServer).SearchLendingPost(ctx, req.(*SearchLendingPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LendingPostService_ServiceDesc is the grpc.ServiceDesc for LendingPostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LendingPostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LendingPostService",
	HandlerType: (*LendingPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLendingPost",
			Handler:    _LendingPostService_CreateLendingPost_Handler,
		},
		{
			MethodName: "GetLendingPostDetail",
			Handler:    _LendingPostService_GetLendingPostDetail_Handler,
		},
		{
			MethodName: "SearchLendingPost",
			Handler:    _LendingPostService_SearchLendingPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/post/post.proto",
}

const (
	BorrowingPostService_CreateBorrowingPost_FullMethodName    = "/BorrowingPostService/CreateBorrowingPost"
	BorrowingPostService_GetBorrowingPostDetail_FullMethodName = "/BorrowingPostService/GetBorrowingPostDetail"
	BorrowingPostService_SearchBorrowingPost_FullMethodName    = "/BorrowingPostService/SearchBorrowingPost"
)

// BorrowingPostServiceClient is the client API for BorrowingPostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BorrowingPostServiceClient interface {
	// RPC method to create a new BorrowingPost
	CreateBorrowingPost(ctx context.Context, in *CreateBorrowingPostRequest, opts ...grpc.CallOption) (*CreateBorrowingPostResponse, error)
	GetBorrowingPostDetail(ctx context.Context, in *GetBorrowingPostDetailRequest, opts ...grpc.CallOption) (*BorrowingPost, error)
	SearchBorrowingPost(ctx context.Context, in *SearchBorrowingPostRequest, opts ...grpc.CallOption) (*BorrowingPostList, error)
}

type borrowingPostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBorrowingPostServiceClient(cc grpc.ClientConnInterface) BorrowingPostServiceClient {
	return &borrowingPostServiceClient{cc}
}

func (c *borrowingPostServiceClient) CreateBorrowingPost(ctx context.Context, in *CreateBorrowingPostRequest, opts ...grpc.CallOption) (*CreateBorrowingPostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBorrowingPostResponse)
	err := c.cc.Invoke(ctx, BorrowingPostService_CreateBorrowingPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowingPostServiceClient) GetBorrowingPostDetail(ctx context.Context, in *GetBorrowingPostDetailRequest, opts ...grpc.CallOption) (*BorrowingPost, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BorrowingPost)
	err := c.cc.Invoke(ctx, BorrowingPostService_GetBorrowingPostDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowingPostServiceClient) SearchBorrowingPost(ctx context.Context, in *SearchBorrowingPostRequest, opts ...grpc.CallOption) (*BorrowingPostList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BorrowingPostList)
	err := c.cc.Invoke(ctx, BorrowingPostService_SearchBorrowingPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BorrowingPostServiceServer is the server API for BorrowingPostService service.
// All implementations must embed UnimplementedBorrowingPostServiceServer
// for forward compatibility.
type BorrowingPostServiceServer interface {
	// RPC method to create a new BorrowingPost
	CreateBorrowingPost(context.Context, *CreateBorrowingPostRequest) (*CreateBorrowingPostResponse, error)
	GetBorrowingPostDetail(context.Context, *GetBorrowingPostDetailRequest) (*BorrowingPost, error)
	SearchBorrowingPost(context.Context, *SearchBorrowingPostRequest) (*BorrowingPostList, error)
	mustEmbedUnimplementedBorrowingPostServiceServer()
}

// UnimplementedBorrowingPostServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBorrowingPostServiceServer struct{}

func (UnimplementedBorrowingPostServiceServer) CreateBorrowingPost(context.Context, *CreateBorrowingPostRequest) (*CreateBorrowingPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBorrowingPost not implemented")
}
func (UnimplementedBorrowingPostServiceServer) GetBorrowingPostDetail(context.Context, *GetBorrowingPostDetailRequest) (*BorrowingPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBorrowingPostDetail not implemented")
}
func (UnimplementedBorrowingPostServiceServer) SearchBorrowingPost(context.Context, *SearchBorrowingPostRequest) (*BorrowingPostList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBorrowingPost not implemented")
}
func (UnimplementedBorrowingPostServiceServer) mustEmbedUnimplementedBorrowingPostServiceServer() {}
func (UnimplementedBorrowingPostServiceServer) testEmbeddedByValue()                              {}

// UnsafeBorrowingPostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BorrowingPostServiceServer will
// result in compilation errors.
type UnsafeBorrowingPostServiceServer interface {
	mustEmbedUnimplementedBorrowingPostServiceServer()
}

func RegisterBorrowingPostServiceServer(s grpc.ServiceRegistrar, srv BorrowingPostServiceServer) {
	// If the following call pancis, it indicates UnimplementedBorrowingPostServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BorrowingPostService_ServiceDesc, srv)
}

func _BorrowingPostService_CreateBorrowingPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBorrowingPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowingPostServiceServer).CreateBorrowingPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BorrowingPostService_CreateBorrowingPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowingPostServiceServer).CreateBorrowingPost(ctx, req.(*CreateBorrowingPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BorrowingPostService_GetBorrowingPostDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBorrowingPostDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowingPostServiceServer).GetBorrowingPostDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BorrowingPostService_GetBorrowingPostDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowingPostServiceServer).GetBorrowingPostDetail(ctx, req.(*GetBorrowingPostDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BorrowingPostService_SearchBorrowingPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBorrowingPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowingPostServiceServer).SearchBorrowingPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BorrowingPostService_SearchBorrowingPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowingPostServiceServer).SearchBorrowingPost(ctx, req.(*SearchBorrowingPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BorrowingPostService_ServiceDesc is the grpc.ServiceDesc for BorrowingPostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BorrowingPostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BorrowingPostService",
	HandlerType: (*BorrowingPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBorrowingPost",
			Handler:    _BorrowingPostService_CreateBorrowingPost_Handler,
		},
		{
			MethodName: "GetBorrowingPostDetail",
			Handler:    _BorrowingPostService_GetBorrowingPostDetail_Handler,
		},
		{
			MethodName: "SearchBorrowingPost",
			Handler:    _BorrowingPostService_SearchBorrowingPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/post/post.proto",
}
