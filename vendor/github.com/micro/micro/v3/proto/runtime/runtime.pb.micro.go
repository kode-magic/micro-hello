// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/runtime/runtime.proto

package runtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Runtime service

func NewRuntimeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Runtime service

type RuntimeService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Logs(ctx context.Context, in *LogsRequest, opts ...client.CallOption) (Runtime_LogsService, error)
}

type runtimeService struct {
	c    client.Client
	name string
}

func NewRuntimeService(name string, c client.Client) RuntimeService {
	return &runtimeService{
		c:    c,
		name: name,
	}
}

func (c *runtimeService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Logs(ctx context.Context, in *LogsRequest, opts ...client.CallOption) (Runtime_LogsService, error) {
	req := c.c.NewRequest(c.name, "Runtime.Logs", &LogsRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &runtimeServiceLogs{stream}, nil
}

type Runtime_LogsService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*LogRecord, error)
}

type runtimeServiceLogs struct {
	stream client.Stream
}

func (x *runtimeServiceLogs) Close() error {
	return x.stream.Close()
}

func (x *runtimeServiceLogs) Context() context.Context {
	return x.stream.Context()
}

func (x *runtimeServiceLogs) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeServiceLogs) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeServiceLogs) Recv() (*LogRecord, error) {
	m := new(LogRecord)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Runtime service

type RuntimeHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Logs(context.Context, *LogsRequest, Runtime_LogsStream) error
}

func RegisterRuntimeHandler(s server.Server, hdlr RuntimeHandler, opts ...server.HandlerOption) error {
	type runtime interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Logs(ctx context.Context, stream server.Stream) error
	}
	type Runtime struct {
		runtime
	}
	h := &runtimeHandler{hdlr}
	return s.Handle(s.NewHandler(&Runtime{h}, opts...))
}

type runtimeHandler struct {
	RuntimeHandler
}

func (h *runtimeHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.RuntimeHandler.Create(ctx, in, out)
}

func (h *runtimeHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.RuntimeHandler.Read(ctx, in, out)
}

func (h *runtimeHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.RuntimeHandler.Delete(ctx, in, out)
}

func (h *runtimeHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.RuntimeHandler.Update(ctx, in, out)
}

func (h *runtimeHandler) Logs(ctx context.Context, stream server.Stream) error {
	m := new(LogsRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RuntimeHandler.Logs(ctx, m, &runtimeLogsStream{stream})
}

type Runtime_LogsStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*LogRecord) error
}

type runtimeLogsStream struct {
	stream server.Stream
}

func (x *runtimeLogsStream) Close() error {
	return x.stream.Close()
}

func (x *runtimeLogsStream) Context() context.Context {
	return x.stream.Context()
}

func (x *runtimeLogsStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeLogsStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeLogsStream) Send(m *LogRecord) error {
	return x.stream.Send(m)
}

// Api Endpoints for Source service

func NewSourceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Source service

type SourceService interface {
	Upload(ctx context.Context, opts ...client.CallOption) (Source_UploadService, error)
}

type sourceService struct {
	c    client.Client
	name string
}

func NewSourceService(name string, c client.Client) SourceService {
	return &sourceService{
		c:    c,
		name: name,
	}
}

func (c *sourceService) Upload(ctx context.Context, opts ...client.CallOption) (Source_UploadService, error) {
	req := c.c.NewRequest(c.name, "Source.Upload", &UploadRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &sourceServiceUpload{stream}, nil
}

type Source_UploadService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseAndRecv() (*UploadResponse, error)
	Send(*UploadRequest) error
}

type sourceServiceUpload struct {
	stream client.Stream
}

func (x *sourceServiceUpload) CloseAndRecv() (*UploadResponse, error) {
	if err := x.stream.Close(); err != nil {
		return nil, err
	}
	r := new(UploadResponse)
	err := x.RecvMsg(r)
	return r, err
}

func (x *sourceServiceUpload) Context() context.Context {
	return x.stream.Context()
}

func (x *sourceServiceUpload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceServiceUpload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceServiceUpload) Send(m *UploadRequest) error {
	return x.stream.Send(m)
}

// Server API for Source service

type SourceHandler interface {
	Upload(context.Context, Source_UploadStream) error
}

func RegisterSourceHandler(s server.Server, hdlr SourceHandler, opts ...server.HandlerOption) error {
	type source interface {
		Upload(ctx context.Context, stream server.Stream) error
	}
	type Source struct {
		source
	}
	h := &sourceHandler{hdlr}
	return s.Handle(s.NewHandler(&Source{h}, opts...))
}

type sourceHandler struct {
	SourceHandler
}

func (h *sourceHandler) Upload(ctx context.Context, stream server.Stream) error {
	return h.SourceHandler.Upload(ctx, &sourceUploadStream{stream})
}

type Source_UploadStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
}

type sourceUploadStream struct {
	stream server.Stream
}

func (x *sourceUploadStream) SendAndClose(in *UploadResponse) error {
	if err := x.SendMsg(in); err != nil {
		return err
	}
	return x.stream.Close()
}

func (x *sourceUploadStream) Context() context.Context {
	return x.stream.Context()
}

func (x *sourceUploadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceUploadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceUploadStream) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Api Endpoints for Build service

func NewBuildEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Build service

type BuildService interface {
	Read(ctx context.Context, in *Service, opts ...client.CallOption) (Build_ReadService, error)
}

type buildService struct {
	c    client.Client
	name string
}

func NewBuildService(name string, c client.Client) BuildService {
	return &buildService{
		c:    c,
		name: name,
	}
}

func (c *buildService) Read(ctx context.Context, in *Service, opts ...client.CallOption) (Build_ReadService, error) {
	req := c.c.NewRequest(c.name, "Build.Read", &Service{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &buildServiceRead{stream}, nil
}

type Build_ReadService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*BuildReadResponse, error)
}

type buildServiceRead struct {
	stream client.Stream
}

func (x *buildServiceRead) Close() error {
	return x.stream.Close()
}

func (x *buildServiceRead) Context() context.Context {
	return x.stream.Context()
}

func (x *buildServiceRead) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *buildServiceRead) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *buildServiceRead) Recv() (*BuildReadResponse, error) {
	m := new(BuildReadResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Build service

type BuildHandler interface {
	Read(context.Context, *Service, Build_ReadStream) error
}

func RegisterBuildHandler(s server.Server, hdlr BuildHandler, opts ...server.HandlerOption) error {
	type build interface {
		Read(ctx context.Context, stream server.Stream) error
	}
	type Build struct {
		build
	}
	h := &buildHandler{hdlr}
	return s.Handle(s.NewHandler(&Build{h}, opts...))
}

type buildHandler struct {
	BuildHandler
}

func (h *buildHandler) Read(ctx context.Context, stream server.Stream) error {
	m := new(Service)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BuildHandler.Read(ctx, m, &buildReadStream{stream})
}

type Build_ReadStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BuildReadResponse) error
}

type buildReadStream struct {
	stream server.Stream
}

func (x *buildReadStream) Close() error {
	return x.stream.Close()
}

func (x *buildReadStream) Context() context.Context {
	return x.stream.Context()
}

func (x *buildReadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *buildReadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *buildReadStream) Send(m *BuildReadResponse) error {
	return x.stream.Send(m)
}
