// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KeywordsClient is the client API for Keywords service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeywordsClient interface {
	StreamingKeyword(ctx context.Context, in *Input, opts ...grpc.CallOption) (Keywords_StreamingKeywordClient, error)
}

type keywordsClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordsClient(cc grpc.ClientConnInterface) KeywordsClient {
	return &keywordsClient{cc}
}

func (c *keywordsClient) StreamingKeyword(ctx context.Context, in *Input, opts ...grpc.CallOption) (Keywords_StreamingKeywordClient, error) {
	stream, err := c.cc.NewStream(ctx, &Keywords_ServiceDesc.Streams[0], "/Keywords/StreamingKeyword", opts...)
	if err != nil {
		return nil, err
	}
	x := &keywordsStreamingKeywordClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Keywords_StreamingKeywordClient interface {
	Recv() (*KeywordsData, error)
	grpc.ClientStream
}

type keywordsStreamingKeywordClient struct {
	grpc.ClientStream
}

func (x *keywordsStreamingKeywordClient) Recv() (*KeywordsData, error) {
	m := new(KeywordsData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KeywordsServer is the server API for Keywords service.
// All implementations must embed UnimplementedKeywordsServer
// for forward compatibility
type KeywordsServer interface {
	StreamingKeyword(*Input, Keywords_StreamingKeywordServer) error
	mustEmbedUnimplementedKeywordsServer()
}

// UnimplementedKeywordsServer must be embedded to have forward compatible implementations.
type UnimplementedKeywordsServer struct {
}

func (UnimplementedKeywordsServer) StreamingKeyword(*Input, Keywords_StreamingKeywordServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingKeyword not implemented")
}
func (UnimplementedKeywordsServer) mustEmbedUnimplementedKeywordsServer() {}

// UnsafeKeywordsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordsServer will
// result in compilation errors.
type UnsafeKeywordsServer interface {
	mustEmbedUnimplementedKeywordsServer()
}

func RegisterKeywordsServer(s grpc.ServiceRegistrar, srv KeywordsServer) {
	s.RegisterService(&Keywords_ServiceDesc, srv)
}

func _Keywords_StreamingKeyword_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Input)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeywordsServer).StreamingKeyword(m, &keywordsStreamingKeywordServer{stream})
}

type Keywords_StreamingKeywordServer interface {
	Send(*KeywordsData) error
	grpc.ServerStream
}

type keywordsStreamingKeywordServer struct {
	grpc.ServerStream
}

func (x *keywordsStreamingKeywordServer) Send(m *KeywordsData) error {
	return x.ServerStream.SendMsg(m)
}

// Keywords_ServiceDesc is the grpc.ServiceDesc for Keywords service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Keywords_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Keywords",
	HandlerType: (*KeywordsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingKeyword",
			Handler:       _Keywords_StreamingKeyword_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "streaming.proto",
}
