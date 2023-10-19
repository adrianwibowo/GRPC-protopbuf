// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type UserGender int32

const (
	UserGender_UNDEFINED UserGender = 0
	UserGender_MALE      UserGender = 1
	UserGender_FEMALE    UserGender = 2
)

var UserGender_name = map[int32]string{
	0: "UNDEFINED",
	1: "MALE",
	2: "FEMALE",
}

var UserGender_value = map[string]int32{
	"UNDEFINED": 0,
	"MALE":      1,
	"FEMALE":    2,
}

func (x UserGender) String() string {
	return proto.EnumName(UserGender_name, int32(x))
}

func (UserGender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

type User struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Gender               UserGender `protobuf:"varint,4,opt,name=gender,proto3,enum=model.UserGender" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetGender() UserGender {
	if m != nil {
		return m.Gender
	}
	return UserGender_UNDEFINED
}

type UserList struct {
	List                 []*User  `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterEnum("model.UserGender", UserGender_name, UserGender_value)
	proto.RegisterType((*User)(nil), "model.User")
	proto.RegisterType((*UserList)(nil), "model.UserList")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0x74, 0x1b, 0xd2, 0x29, 0xd6, 0x3a, 0x07, 0x09, 0xf1, 0x60, 0xc8, 0x29, 0x2a,
	0x6c, 0x49, 0xfc, 0x04, 0x42, 0x53, 0x11, 0x6a, 0x0f, 0x81, 0x7e, 0x80, 0x96, 0x5d, 0xc3, 0x42,
	0xd2, 0x8d, 0xbb, 0x5b, 0xc4, 0x6f, 0x2f, 0x99, 0xfa, 0x27, 0x17, 0x6f, 0xb3, 0x6f, 0x7e, 0xef,
	0xed, 0xf0, 0x00, 0x4e, 0x56, 0x1a, 0xde, 0x19, 0xed, 0x34, 0x4e, 0x5a, 0x2d, 0x64, 0x13, 0xdf,
	0xd4, 0x5a, 0xd7, 0x8d, 0x5c, 0x92, 0x78, 0x38, 0xbd, 0x2d, 0xcb, 0xb6, 0x73, 0x9f, 0x67, 0x26,
	0x7d, 0x07, 0xb6, 0xb3, 0xd2, 0xe0, 0x1c, 0x7c, 0x25, 0x22, 0x2f, 0xf1, 0xb2, 0x69, 0xe5, 0x2b,
	0x81, 0x08, 0xec, 0xb8, 0x6f, 0x65, 0xe4, 0x93, 0x42, 0x33, 0xc6, 0x10, 0x76, 0x7b, 0x6b, 0x3f,
	0xb4, 0x11, 0xd1, 0x98, 0xf4, 0xdf, 0x37, 0xde, 0x41, 0x50, 0xcb, 0xa3, 0x90, 0x26, 0x62, 0x89,
	0x97, 0xcd, 0x8b, 0x2b, 0x4e, 0x9f, 0xf3, 0x3e, 0xfc, 0x99, 0x16, 0xd5, 0x37, 0x90, 0x3e, 0x40,
	0xd8, 0xab, 0x1b, 0x65, 0x1d, 0xde, 0x02, 0x6b, 0x94, 0x75, 0x91, 0x97, 0x8c, 0xb3, 0x59, 0x31,
	0x1b, 0x98, 0x2a, 0x5a, 0xdc, 0xe7, 0x00, 0x7f, 0x11, 0x78, 0x01, 0xd3, 0xdd, 0x76, 0x55, 0xae,
	0x5f, 0xb6, 0xe5, 0x6a, 0x31, 0xc2, 0x10, 0xd8, 0xeb, 0xd3, 0xa6, 0x5c, 0x78, 0x08, 0x10, 0xac,
	0x4b, 0x9a, 0xfd, 0xa2, 0x85, 0x49, 0x6f, 0xb1, 0x98, 0x43, 0x58, 0xc9, 0x5a, 0x59, 0x27, 0x0d,
	0x0e, 0xa3, 0xe3, 0x6b, 0x7e, 0xae, 0x84, 0xff, 0x54, 0xc2, 0xa9, 0x92, 0x74, 0x84, 0x39, 0x30,
	0xba, 0xeb, 0x1f, 0x22, 0xbe, 0x1c, 0xc4, 0xf4, 0x60, 0x3a, 0x3a, 0x04, 0x84, 0x3c, 0x7e, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x6a, 0x7a, 0xef, 0x84, 0x7a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserList, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/model.Users/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/model.Users/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Register(context.Context, *User) (*emptypb.Empty, error)
	List(context.Context, *emptypb.Empty) (*UserList, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Users/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Users/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Users_Register_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Users_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}