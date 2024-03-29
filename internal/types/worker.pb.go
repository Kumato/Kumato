// Code generated by protoc-gen-go. DO NOT EDIT.
// source: worker.proto

package types

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Images struct {
	ImageRepoTags        []*ImageRepoTags `protobuf:"bytes,1,rep,name=image_repo_tags,json=imageRepoTags,proto3" json:"image_repo_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Images) Reset()         { *m = Images{} }
func (m *Images) String() string { return proto.CompactTextString(m) }
func (*Images) ProtoMessage()    {}
func (*Images) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{0}
}

func (m *Images) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Images.Unmarshal(m, b)
}
func (m *Images) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Images.Marshal(b, m, deterministic)
}
func (m *Images) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Images.Merge(m, src)
}
func (m *Images) XXX_Size() int {
	return xxx_messageInfo_Images.Size(m)
}
func (m *Images) XXX_DiscardUnknown() {
	xxx_messageInfo_Images.DiscardUnknown(m)
}

var xxx_messageInfo_Images proto.InternalMessageInfo

func (m *Images) GetImageRepoTags() []*ImageRepoTags {
	if m != nil {
		return m.ImageRepoTags
	}
	return nil
}

type ImageRepoTags struct {
	Repo                 string   `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Tags                 []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageRepoTags) Reset()         { *m = ImageRepoTags{} }
func (m *ImageRepoTags) String() string { return proto.CompactTextString(m) }
func (*ImageRepoTags) ProtoMessage()    {}
func (*ImageRepoTags) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{1}
}

func (m *ImageRepoTags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageRepoTags.Unmarshal(m, b)
}
func (m *ImageRepoTags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageRepoTags.Marshal(b, m, deterministic)
}
func (m *ImageRepoTags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageRepoTags.Merge(m, src)
}
func (m *ImageRepoTags) XXX_Size() int {
	return xxx_messageInfo_ImageRepoTags.Size(m)
}
func (m *ImageRepoTags) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageRepoTags.DiscardUnknown(m)
}

var xxx_messageInfo_ImageRepoTags proto.InternalMessageInfo

func (m *ImageRepoTags) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *ImageRepoTags) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Stats struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	MemoryTotal          int64    `protobuf:"varint,2,opt,name=memory_total,json=memoryTotal,proto3" json:"memory_total,omitempty"`
	MemoryUsed           int64    `protobuf:"varint,3,opt,name=memory_used,json=memoryUsed,proto3" json:"memory_used,omitempty"`
	NanoCpuTotal         int64    `protobuf:"varint,4,opt,name=nano_cpu_total,json=nanoCpuTotal,proto3" json:"nano_cpu_total,omitempty"`
	NanoCpuUsed          int64    `protobuf:"varint,5,opt,name=nano_cpu_used,json=nanoCpuUsed,proto3" json:"nano_cpu_used,omitempty"`
	GpuTotal             int64    `protobuf:"varint,6,opt,name=gpu_total,json=gpuTotal,proto3" json:"gpu_total,omitempty"`
	GpuUsed              int64    `protobuf:"varint,7,opt,name=gpu_used,json=gpuUsed,proto3" json:"gpu_used,omitempty"`
	CpuUsage             float64  `protobuf:"fixed64,8,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{2}
}

func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Stats) GetMemoryTotal() int64 {
	if m != nil {
		return m.MemoryTotal
	}
	return 0
}

func (m *Stats) GetMemoryUsed() int64 {
	if m != nil {
		return m.MemoryUsed
	}
	return 0
}

func (m *Stats) GetNanoCpuTotal() int64 {
	if m != nil {
		return m.NanoCpuTotal
	}
	return 0
}

func (m *Stats) GetNanoCpuUsed() int64 {
	if m != nil {
		return m.NanoCpuUsed
	}
	return 0
}

func (m *Stats) GetGpuTotal() int64 {
	if m != nil {
		return m.GpuTotal
	}
	return 0
}

func (m *Stats) GetGpuUsed() int64 {
	if m != nil {
		return m.GpuUsed
	}
	return 0
}

func (m *Stats) GetCpuUsage() float64 {
	if m != nil {
		return m.CpuUsage
	}
	return 0
}

func init() {
	proto.RegisterType((*Images)(nil), "types.Images")
	proto.RegisterType((*ImageRepoTags)(nil), "types.ImageRepoTags")
	proto.RegisterType((*Stats)(nil), "types.Stats")
}

func init() {
	proto.RegisterFile("worker.proto", fileDescriptor_e4ff6184b07e587a)
}

var fileDescriptor_e4ff6184b07e587a = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xc7, 0x93, 0xa6, 0x4d, 0xd2, 0xd3, 0xe4, 0x5e, 0x18, 0xee, 0x22, 0x37, 0x2e, 0x8c, 0x83,
	0xd2, 0xe0, 0xa2, 0x8b, 0xba, 0x70, 0xe3, 0x4e, 0xb4, 0x08, 0xae, 0xd2, 0x8a, 0xcb, 0x30, 0xb6,
	0xc3, 0x58, 0x6a, 0x32, 0x43, 0x66, 0x82, 0xf4, 0x05, 0x7c, 0x45, 0x5f, 0x47, 0x66, 0x26, 0xad,
	0x2d, 0x82, 0xbb, 0xf3, 0xff, 0x38, 0xbf, 0x40, 0xce, 0x40, 0xf4, 0xce, 0x9b, 0x0d, 0x6d, 0x26,
	0xa2, 0xe1, 0x8a, 0xa3, 0x81, 0xda, 0x0a, 0x2a, 0xd3, 0x68, 0xc9, 0xab, 0x8a, 0xd7, 0xd6, 0x4c,
	0x41, 0x11, 0xb9, 0xb1, 0x33, 0xbe, 0x07, 0xff, 0xa1, 0x22, 0x8c, 0x4a, 0x74, 0x03, 0x7f, 0xd7,
	0x7a, 0x2a, 0x1b, 0x2a, 0x78, 0xa9, 0x08, 0x93, 0x89, 0x9b, 0x79, 0xf9, 0x68, 0xfa, 0x6f, 0x62,
	0x20, 0x13, 0xd3, 0x2b, 0xa8, 0xe0, 0x0b, 0xc2, 0x64, 0x11, 0xaf, 0x0f, 0x25, 0xbe, 0x86, 0xf8,
	0x28, 0x47, 0x08, 0xfa, 0x1a, 0x94, 0xb8, 0x99, 0x9b, 0x0f, 0x0b, 0x33, 0x6b, 0xcf, 0x70, 0x7b,
	0x99, 0xa7, 0x3d, 0x3d, 0xe3, 0x8f, 0x1e, 0x0c, 0xe6, 0x8a, 0x28, 0x89, 0x52, 0x08, 0x5f, 0xb9,
	0x54, 0x35, 0xa9, 0x68, 0xb7, 0xb5, 0xd7, 0xe8, 0x0c, 0xa2, 0x8a, 0x56, 0xbc, 0xd9, 0x96, 0x8a,
	0x2b, 0xf2, 0x96, 0xf4, 0x32, 0x37, 0xf7, 0x8a, 0x91, 0xf5, 0x16, 0xda, 0x42, 0xa7, 0xd0, 0xc9,
	0xb2, 0x95, 0x74, 0x95, 0x78, 0xa6, 0x01, 0xd6, 0x7a, 0x92, 0x74, 0x85, 0xce, 0xe1, 0x4f, 0x4d,
	0x6a, 0x5e, 0x2e, 0x45, 0xdb, 0x51, 0xfa, 0xa6, 0x13, 0x69, 0xf7, 0x56, 0xb4, 0x16, 0x83, 0x21,
	0xde, 0xb7, 0x0c, 0x68, 0x60, 0x3f, 0xd5, 0x95, 0x0c, 0xe9, 0x04, 0x86, 0x6c, 0x0f, 0xf1, 0x4d,
	0x1e, 0xb2, 0x1d, 0xe0, 0x3f, 0xe8, 0xd9, 0xee, 0x06, 0x26, 0x0b, 0xd8, 0xf7, 0x9e, 0xc5, 0x12,
	0x46, 0x93, 0x30, 0x73, 0x73, 0xb7, 0x08, 0x97, 0x3a, 0x23, 0x8c, 0x4e, 0x3f, 0x5d, 0xf0, 0x9f,
	0xcd, 0xed, 0xd0, 0x18, 0xc2, 0xb9, 0xe2, 0x62, 0x41, 0xe4, 0x06, 0x8d, 0xba, 0xbf, 0xaf, 0x45,
	0x1a, 0x75, 0xe2, 0xae, 0x12, 0x6a, 0x8b, 0x1d, 0x74, 0x01, 0x41, 0xd1, 0xd6, 0x3f, 0x7b, 0x87,
	0x02, 0x3b, 0xe8, 0x12, 0x86, 0x33, 0xaa, 0xba, 0x3b, 0x1f, 0x31, 0xd2, 0xf8, 0xf0, 0xb8, 0x12,
	0x3b, 0x28, 0x87, 0x70, 0x46, 0x95, 0xbd, 0xc8, 0x71, 0x75, 0xa7, 0x4c, 0x86, 0x1d, 0x34, 0x86,
	0x60, 0x46, 0xd5, 0x23, 0x27, 0xab, 0xdf, 0x8b, 0x2f, 0xbe, 0x79, 0x6a, 0x57, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x50, 0x85, 0x95, 0x1d, 0x9b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerClient interface {
	StopTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Empty, error)
	RunTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	GetImages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Images, error)
	GetStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Stats, error)
	GetLoad(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Stats, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) StopTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/types.Worker/StopTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) RunTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/types.Worker/RunTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetImages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Images, error) {
	out := new(Images)
	err := c.cc.Invoke(ctx, "/types.Worker/GetImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := c.cc.Invoke(ctx, "/types.Worker/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetLoad(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := c.cc.Invoke(ctx, "/types.Worker/GetLoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
type WorkerServer interface {
	StopTask(context.Context, *Task) (*Empty, error)
	RunTask(context.Context, *Task) (*Task, error)
	GetImages(context.Context, *Empty) (*Images, error)
	GetStats(context.Context, *Empty) (*Stats, error)
	GetLoad(context.Context, *Empty) (*Stats, error)
}

// UnimplementedWorkerServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (*UnimplementedWorkerServer) StopTask(ctx context.Context, req *Task) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopTask not implemented")
}
func (*UnimplementedWorkerServer) RunTask(ctx context.Context, req *Task) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunTask not implemented")
}
func (*UnimplementedWorkerServer) GetImages(ctx context.Context, req *Empty) (*Images, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImages not implemented")
}
func (*UnimplementedWorkerServer) GetStats(ctx context.Context, req *Empty) (*Stats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (*UnimplementedWorkerServer) GetLoad(ctx context.Context, req *Empty) (*Stats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoad not implemented")
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_StopTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).StopTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Worker/StopTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).StopTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_RunTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).RunTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Worker/RunTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).RunTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Worker/GetImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetImages(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Worker/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetStats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Worker/GetLoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetLoad(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StopTask",
			Handler:    _Worker_StopTask_Handler,
		},
		{
			MethodName: "RunTask",
			Handler:    _Worker_RunTask_Handler,
		},
		{
			MethodName: "GetImages",
			Handler:    _Worker_GetImages_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _Worker_GetStats_Handler,
		},
		{
			MethodName: "GetLoad",
			Handler:    _Worker_GetLoad_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "worker.proto",
}
