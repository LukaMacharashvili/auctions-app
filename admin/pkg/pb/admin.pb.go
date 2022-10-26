// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin.proto

package pb

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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Role                 string   `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	Balance              float64  `protobuf:"fixed64,6,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type LoadRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadRequest) Reset()         { *m = LoadRequest{} }
func (m *LoadRequest) String() string { return proto.CompactTextString(m) }
func (*LoadRequest) ProtoMessage()    {}
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{1}
}

func (m *LoadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadRequest.Unmarshal(m, b)
}
func (m *LoadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadRequest.Marshal(b, m, deterministic)
}
func (m *LoadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadRequest.Merge(m, src)
}
func (m *LoadRequest) XXX_Size() int {
	return xxx_messageInfo_LoadRequest.Size(m)
}
func (m *LoadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadRequest proto.InternalMessageInfo

type LoadResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadResponse) Reset()         { *m = LoadResponse{} }
func (m *LoadResponse) String() string { return proto.CompactTextString(m) }
func (*LoadResponse) ProtoMessage()    {}
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{2}
}

func (m *LoadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadResponse.Unmarshal(m, b)
}
func (m *LoadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadResponse.Marshal(b, m, deterministic)
}
func (m *LoadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadResponse.Merge(m, src)
}
func (m *LoadResponse) XXX_Size() int {
	return xxx_messageInfo_LoadResponse.Size(m)
}
func (m *LoadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadResponse proto.InternalMessageInfo

func (m *LoadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{3}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{4}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DeleteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{5}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteResponse struct {
	Deleted              bool     `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{6}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *DeleteResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "admin.User")
	proto.RegisterType((*LoadRequest)(nil), "admin.LoadRequest")
	proto.RegisterType((*LoadResponse)(nil), "admin.LoadResponse")
	proto.RegisterType((*CreateRequest)(nil), "admin.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "admin.CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "admin.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "admin.DeleteResponse")
}

func init() {
	proto.RegisterFile("admin.proto", fileDescriptor_73a7fc70dcc2027c)
}

var fileDescriptor_73a7fc70dcc2027c = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x3f, 0x6f, 0xc2, 0x30,
	0x10, 0xc5, 0x71, 0x08, 0x81, 0x5e, 0x80, 0xe1, 0x4a, 0x2b, 0x8b, 0x05, 0xe4, 0x89, 0x89, 0x22,
	0x18, 0xba, 0x55, 0xfd, 0x37, 0x76, 0x4a, 0xd5, 0xa5, 0x9b, 0x21, 0x37, 0x44, 0x0a, 0x71, 0x6a,
	0x87, 0xf6, 0x7b, 0xf4, 0xbb, 0xf4, 0xfb, 0x55, 0xb1, 0x1d, 0xda, 0xa0, 0x4a, 0xdd, 0xf8, 0x9d,
	0xef, 0xf1, 0xee, 0x3d, 0x05, 0x62, 0x99, 0xee, 0xb3, 0x62, 0x59, 0x6a, 0x55, 0x29, 0xec, 0x59,
	0x10, 0x9f, 0x0c, 0xc2, 0x17, 0x43, 0x1a, 0xc7, 0x10, 0x64, 0x29, 0x67, 0x73, 0xb6, 0xe8, 0x26,
	0x41, 0x96, 0xe2, 0x14, 0x06, 0x07, 0x43, 0xba, 0x90, 0x7b, 0xe2, 0xc1, 0x9c, 0x2d, 0xce, 0x92,
	0x23, 0xe3, 0x04, 0x7a, 0xb4, 0x97, 0x59, 0xce, 0xbb, 0xf6, 0xc1, 0x41, 0xad, 0x28, 0xa5, 0x31,
	0x1f, 0x4a, 0xa7, 0x3c, 0x74, 0x8a, 0x86, 0x11, 0x21, 0xd4, 0x2a, 0x27, 0xde, 0xb3, 0x73, 0xfb,
	0x1b, 0x39, 0xf4, 0xb7, 0x32, 0x97, 0xc5, 0x8e, 0x78, 0x34, 0x67, 0x0b, 0x96, 0x34, 0x28, 0x46,
	0x10, 0x3f, 0x29, 0x99, 0x26, 0xf4, 0x76, 0x20, 0x53, 0x89, 0x2b, 0x18, 0x3a, 0x34, 0xa5, 0x2a,
	0x0c, 0xe1, 0x0c, 0xc2, 0xfa, 0x14, 0x7b, 0x6c, 0xbc, 0x8e, 0x97, 0x2e, 0x56, 0x9d, 0x22, 0xb1,
	0x0f, 0x62, 0x05, 0xa3, 0x07, 0x4d, 0xb2, 0x22, 0xff, 0x0f, 0xff, 0x2b, 0x6e, 0x60, 0xdc, 0x28,
	0xbc, 0xc9, 0x25, 0x44, 0xa6, 0x92, 0xd5, 0xc1, 0xf8, 0x4e, 0x3c, 0xd9, 0xec, 0x5a, 0x2b, 0xed,
	0x4b, 0x71, 0x20, 0x66, 0x30, 0x7a, 0xa4, 0x9c, 0x7e, 0x1c, 0x4f, 0xea, 0x14, 0xb7, 0x30, 0x6e,
	0x16, 0xbc, 0x01, 0x87, 0x7e, 0x6a, 0x27, 0x6e, 0x6d, 0x90, 0x34, 0xf8, 0xb7, 0xc5, 0xfa, 0x8b,
	0xc1, 0xf0, 0xae, 0xbe, 0xfb, 0x99, 0xf4, 0x7b, 0xb6, 0x23, 0xdc, 0x40, 0x58, 0xd7, 0x82, 0xe8,
	0xe3, 0xfc, 0xaa, 0x6c, 0x7a, 0xde, 0x9a, 0x39, 0x47, 0xd1, 0x59, 0x31, 0xbc, 0x86, 0xc8, 0x05,
	0xc5, 0x89, 0x5f, 0x69, 0x35, 0x35, 0xbd, 0x38, 0x99, 0x36, 0xd2, 0x5a, 0xe8, 0x02, 0x1c, 0x85,
	0xad, 0xc0, 0x47, 0x61, 0x3b, 0xa5, 0xe8, 0xdc, 0x87, 0xaf, 0x41, 0xb9, 0xdd, 0x46, 0xf6, 0xab,
	0xdb, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x16, 0x99, 0x26, 0x6d, 0x84, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (AdminService_LoadClient, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (AdminService_LoadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AdminService_serviceDesc.Streams[0], "/admin.AdminService/Load", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminServiceLoadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AdminService_LoadClient interface {
	Recv() (*LoadResponse, error)
	grpc.ClientStream
}

type adminServiceLoadClient struct {
	grpc.ClientStream
}

func (x *adminServiceLoadClient) Recv() (*LoadResponse, error) {
	m := new(LoadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	Load(*LoadRequest, AdminService_LoadServer) error
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedAdminServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (*UnimplementedAdminServiceServer) Load(req *LoadRequest, srv AdminService_LoadServer) error {
	return status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (*UnimplementedAdminServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAdminServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_Load_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServiceServer).Load(m, &adminServiceLoadServer{stream})
}

type AdminService_LoadServer interface {
	Send(*LoadResponse) error
	grpc.ServerStream
}

type adminServiceLoadServer struct {
	grpc.ServerStream
}

func (x *adminServiceLoadServer) Send(m *LoadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AdminService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AdminService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdminService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Load",
			Handler:       _AdminService_Load_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "admin.proto",
}