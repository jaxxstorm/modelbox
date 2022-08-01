// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/service.proto

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

// ModelStoreClient is the client API for ModelStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelStoreClient interface {
	// Create a new Model under a namespace. If no namespace is specified, models
	// are created under a default namespace.
	CreateModel(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*CreateModelResponse, error)
	// List Models uploaded for a namespace
	ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error)
	// Creates a new model version for a model
	CreateModelVersion(ctx context.Context, in *CreateModelVersionRequest, opts ...grpc.CallOption) (*CreateModelVersionResponse, error)
	// Lists model versions for a model.
	ListModelVersions(ctx context.Context, in *ListModelVersionsRequest, opts ...grpc.CallOption) (*ListModelVersionsResponse, error)
	// Creates a new experiment
	CreateExperiment(ctx context.Context, in *CreateExperimentRequest, opts ...grpc.CallOption) (*CreateExperimentResponse, error)
	// List Experiments
	ListExperiments(ctx context.Context, in *ListExperimentsRequest, opts ...grpc.CallOption) (*ListExperimentsResponse, error)
	// Uploads a new checkpoint for an experiment
	CreateCheckpoint(ctx context.Context, in *CreateCheckpointRequest, opts ...grpc.CallOption) (*CreateCheckpointResponse, error)
	// Lists all the checkpoints for an experiment
	ListCheckpoints(ctx context.Context, in *ListCheckpointsRequest, opts ...grpc.CallOption) (*ListCheckpointsResponse, error)
	// Gets a checkpoint from the modelstore for an experiment
	GetCheckpoint(ctx context.Context, in *GetCheckpointRequest, opts ...grpc.CallOption) (*GetCheckpointResponse, error)
	// UploadFile streams a files to ModelBox and stores the binaries to the condfigured storage
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (ModelStore_UploadFileClient, error)
	// DownloadFile downloads a file from configured storage
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (ModelStore_DownloadFileClient, error)
	// Persists a set of metadata related to objects
	UpdateMetadata(ctx context.Context, in *UpdateMetadataRequest, opts ...grpc.CallOption) (*UpdateMetadataResponse, error)
	// Lists metadata associated with an object
	ListMetadata(ctx context.Context, in *ListMetadataRequest, opts ...grpc.CallOption) (*ListMetadataResponse, error)
}

type modelStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewModelStoreClient(cc grpc.ClientConnInterface) ModelStoreClient {
	return &modelStoreClient{cc}
}

func (c *modelStoreClient) CreateModel(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*CreateModelResponse, error) {
	out := new(CreateModelResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/CreateModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelStoreClient) ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ListModelsResponse, error) {
	out := new(ListModelsResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/ListModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelStoreClient) CreateModelVersion(ctx context.Context, in *CreateModelVersionRequest, opts ...grpc.CallOption) (*CreateModelVersionResponse, error) {
	out := new(CreateModelVersionResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/CreateModelVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelStoreClient) ListModelVersions(ctx context.Context, in *ListModelVersionsRequest, opts ...grpc.CallOption) (*ListModelVersionsResponse, error) {
	out := new(ListModelVersionsResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/ListModelVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelStoreClient) CreateExperiment(ctx context.Context, in *CreateExperimentRequest, opts ...grpc.CallOption) (*CreateExperimentResponse, error) {
	out := new(CreateExperimentResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/CreateExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelStoreClient) ListExperiments(ctx context.Context, in *ListExperimentsRequest, opts ...grpc.CallOption) (*ListExperimentsResponse, error) {
	out := new(ListExperimentsResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/ListExperiments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelStoreClient) CreateCheckpoint(ctx context.Context, in *CreateCheckpointRequest, opts ...grpc.CallOption) (*CreateCheckpointResponse, error) {
	out := new(CreateCheckpointResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/CreateCheckpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelStoreClient) ListCheckpoints(ctx context.Context, in *ListCheckpointsRequest, opts ...grpc.CallOption) (*ListCheckpointsResponse, error) {
	out := new(ListCheckpointsResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/ListCheckpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelStoreClient) GetCheckpoint(ctx context.Context, in *GetCheckpointRequest, opts ...grpc.CallOption) (*GetCheckpointResponse, error) {
	out := new(GetCheckpointResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/GetCheckpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelStoreClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (ModelStore_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &ModelStore_ServiceDesc.Streams[0], "/modelbox.ModelStore/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &modelStoreUploadFileClient{stream}
	return x, nil
}

type ModelStore_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type modelStoreUploadFileClient struct {
	grpc.ClientStream
}

func (x *modelStoreUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *modelStoreUploadFileClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *modelStoreClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (ModelStore_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &ModelStore_ServiceDesc.Streams[1], "/modelbox.ModelStore/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &modelStoreDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ModelStore_DownloadFileClient interface {
	Recv() (*DownloadFileResponse, error)
	grpc.ClientStream
}

type modelStoreDownloadFileClient struct {
	grpc.ClientStream
}

func (x *modelStoreDownloadFileClient) Recv() (*DownloadFileResponse, error) {
	m := new(DownloadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *modelStoreClient) UpdateMetadata(ctx context.Context, in *UpdateMetadataRequest, opts ...grpc.CallOption) (*UpdateMetadataResponse, error) {
	out := new(UpdateMetadataResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/UpdateMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelStoreClient) ListMetadata(ctx context.Context, in *ListMetadataRequest, opts ...grpc.CallOption) (*ListMetadataResponse, error) {
	out := new(ListMetadataResponse)
	err := c.cc.Invoke(ctx, "/modelbox.ModelStore/ListMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelStoreServer is the server API for ModelStore service.
// All implementations must embed UnimplementedModelStoreServer
// for forward compatibility
type ModelStoreServer interface {
	// Create a new Model under a namespace. If no namespace is specified, models
	// are created under a default namespace.
	CreateModel(context.Context, *CreateModelRequest) (*CreateModelResponse, error)
	// List Models uploaded for a namespace
	ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error)
	// Creates a new model version for a model
	CreateModelVersion(context.Context, *CreateModelVersionRequest) (*CreateModelVersionResponse, error)
	// Lists model versions for a model.
	ListModelVersions(context.Context, *ListModelVersionsRequest) (*ListModelVersionsResponse, error)
	// Creates a new experiment
	CreateExperiment(context.Context, *CreateExperimentRequest) (*CreateExperimentResponse, error)
	// List Experiments
	ListExperiments(context.Context, *ListExperimentsRequest) (*ListExperimentsResponse, error)
	// Uploads a new checkpoint for an experiment
	CreateCheckpoint(context.Context, *CreateCheckpointRequest) (*CreateCheckpointResponse, error)
	// Lists all the checkpoints for an experiment
	ListCheckpoints(context.Context, *ListCheckpointsRequest) (*ListCheckpointsResponse, error)
	// Gets a checkpoint from the modelstore for an experiment
	GetCheckpoint(context.Context, *GetCheckpointRequest) (*GetCheckpointResponse, error)
	// UploadFile streams a files to ModelBox and stores the binaries to the condfigured storage
	UploadFile(ModelStore_UploadFileServer) error
	// DownloadFile downloads a file from configured storage
	DownloadFile(*DownloadFileRequest, ModelStore_DownloadFileServer) error
	// Persists a set of metadata related to objects
	UpdateMetadata(context.Context, *UpdateMetadataRequest) (*UpdateMetadataResponse, error)
	// Lists metadata associated with an object
	ListMetadata(context.Context, *ListMetadataRequest) (*ListMetadataResponse, error)
	mustEmbedUnimplementedModelStoreServer()
}

// UnimplementedModelStoreServer must be embedded to have forward compatible implementations.
type UnimplementedModelStoreServer struct {
}

func (UnimplementedModelStoreServer) CreateModel(context.Context, *CreateModelRequest) (*CreateModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModel not implemented")
}
func (UnimplementedModelStoreServer) ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModels not implemented")
}
func (UnimplementedModelStoreServer) CreateModelVersion(context.Context, *CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModelVersion not implemented")
}
func (UnimplementedModelStoreServer) ListModelVersions(context.Context, *ListModelVersionsRequest) (*ListModelVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModelVersions not implemented")
}
func (UnimplementedModelStoreServer) CreateExperiment(context.Context, *CreateExperimentRequest) (*CreateExperimentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExperiment not implemented")
}
func (UnimplementedModelStoreServer) ListExperiments(context.Context, *ListExperimentsRequest) (*ListExperimentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExperiments not implemented")
}
func (UnimplementedModelStoreServer) CreateCheckpoint(context.Context, *CreateCheckpointRequest) (*CreateCheckpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCheckpoint not implemented")
}
func (UnimplementedModelStoreServer) ListCheckpoints(context.Context, *ListCheckpointsRequest) (*ListCheckpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCheckpoints not implemented")
}
func (UnimplementedModelStoreServer) GetCheckpoint(context.Context, *GetCheckpointRequest) (*GetCheckpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckpoint not implemented")
}
func (UnimplementedModelStoreServer) UploadFile(ModelStore_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedModelStoreServer) DownloadFile(*DownloadFileRequest, ModelStore_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedModelStoreServer) UpdateMetadata(context.Context, *UpdateMetadataRequest) (*UpdateMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetadata not implemented")
}
func (UnimplementedModelStoreServer) ListMetadata(context.Context, *ListMetadataRequest) (*ListMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetadata not implemented")
}
func (UnimplementedModelStoreServer) mustEmbedUnimplementedModelStoreServer() {}

// UnsafeModelStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelStoreServer will
// result in compilation errors.
type UnsafeModelStoreServer interface {
	mustEmbedUnimplementedModelStoreServer()
}

func RegisterModelStoreServer(s grpc.ServiceRegistrar, srv ModelStoreServer) {
	s.RegisterService(&ModelStore_ServiceDesc, srv)
}

func _ModelStore_CreateModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).CreateModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/CreateModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).CreateModel(ctx, req.(*CreateModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelStore_ListModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).ListModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/ListModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).ListModels(ctx, req.(*ListModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelStore_CreateModelVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModelVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).CreateModelVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/CreateModelVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).CreateModelVersion(ctx, req.(*CreateModelVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelStore_ListModelVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).ListModelVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/ListModelVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).ListModelVersions(ctx, req.(*ListModelVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelStore_CreateExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).CreateExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/CreateExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).CreateExperiment(ctx, req.(*CreateExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelStore_ListExperiments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExperimentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).ListExperiments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/ListExperiments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).ListExperiments(ctx, req.(*ListExperimentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelStore_CreateCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).CreateCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/CreateCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).CreateCheckpoint(ctx, req.(*CreateCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelStore_ListCheckpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCheckpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).ListCheckpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/ListCheckpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).ListCheckpoints(ctx, req.(*ListCheckpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelStore_GetCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).GetCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/GetCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).GetCheckpoint(ctx, req.(*GetCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelStore_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ModelStoreServer).UploadFile(&modelStoreUploadFileServer{stream})
}

type ModelStore_UploadFileServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type modelStoreUploadFileServer struct {
	grpc.ServerStream
}

func (x *modelStoreUploadFileServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *modelStoreUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ModelStore_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ModelStoreServer).DownloadFile(m, &modelStoreDownloadFileServer{stream})
}

type ModelStore_DownloadFileServer interface {
	Send(*DownloadFileResponse) error
	grpc.ServerStream
}

type modelStoreDownloadFileServer struct {
	grpc.ServerStream
}

func (x *modelStoreDownloadFileServer) Send(m *DownloadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ModelStore_UpdateMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).UpdateMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/UpdateMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).UpdateMetadata(ctx, req.(*UpdateMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelStore_ListMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelStoreServer).ListMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modelbox.ModelStore/ListMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelStoreServer).ListMetadata(ctx, req.(*ListMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelStore_ServiceDesc is the grpc.ServiceDesc for ModelStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "modelbox.ModelStore",
	HandlerType: (*ModelStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateModel",
			Handler:    _ModelStore_CreateModel_Handler,
		},
		{
			MethodName: "ListModels",
			Handler:    _ModelStore_ListModels_Handler,
		},
		{
			MethodName: "CreateModelVersion",
			Handler:    _ModelStore_CreateModelVersion_Handler,
		},
		{
			MethodName: "ListModelVersions",
			Handler:    _ModelStore_ListModelVersions_Handler,
		},
		{
			MethodName: "CreateExperiment",
			Handler:    _ModelStore_CreateExperiment_Handler,
		},
		{
			MethodName: "ListExperiments",
			Handler:    _ModelStore_ListExperiments_Handler,
		},
		{
			MethodName: "CreateCheckpoint",
			Handler:    _ModelStore_CreateCheckpoint_Handler,
		},
		{
			MethodName: "ListCheckpoints",
			Handler:    _ModelStore_ListCheckpoints_Handler,
		},
		{
			MethodName: "GetCheckpoint",
			Handler:    _ModelStore_GetCheckpoint_Handler,
		},
		{
			MethodName: "UpdateMetadata",
			Handler:    _ModelStore_UpdateMetadata_Handler,
		},
		{
			MethodName: "ListMetadata",
			Handler:    _ModelStore_ListMetadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _ModelStore_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _ModelStore_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}