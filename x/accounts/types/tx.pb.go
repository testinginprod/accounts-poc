// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: accounts/accounts/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgDeploy struct {
	Sender      string       `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Kind        string       `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	InitMessage []byte       `protobuf:"bytes,3,opt,name=init_message,json=initMessage,proto3" json:"init_message,omitempty"`
	Funds       []types.Coin `protobuf:"bytes,4,rep,name=funds,proto3" json:"funds"`
}

func (m *MsgDeploy) Reset()         { *m = MsgDeploy{} }
func (m *MsgDeploy) String() string { return proto.CompactTextString(m) }
func (*MsgDeploy) ProtoMessage()    {}
func (*MsgDeploy) Descriptor() ([]byte, []int) {
	return fileDescriptor_1acf5eaefb0cd3e7, []int{0}
}
func (m *MsgDeploy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeploy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeploy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeploy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeploy.Merge(m, src)
}
func (m *MsgDeploy) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeploy) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeploy.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeploy proto.InternalMessageInfo

func (m *MsgDeploy) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgDeploy) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *MsgDeploy) GetInitMessage() []byte {
	if m != nil {
		return m.InitMessage
	}
	return nil
}

func (m *MsgDeploy) GetFunds() []types.Coin {
	if m != nil {
		return m.Funds
	}
	return nil
}

type MsgDeployResponse struct {
	Address string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Id      uint64      `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Data    *types1.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgDeployResponse) Reset()         { *m = MsgDeployResponse{} }
func (m *MsgDeployResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeployResponse) ProtoMessage()    {}
func (*MsgDeployResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1acf5eaefb0cd3e7, []int{1}
}
func (m *MsgDeployResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeployResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeployResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeployResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeployResponse.Merge(m, src)
}
func (m *MsgDeployResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeployResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeployResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeployResponse proto.InternalMessageInfo

func (m *MsgDeployResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgDeployResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgDeployResponse) GetData() *types1.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type MsgExecute struct {
	Sender  string       `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Address string       `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Message []byte       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Funds   []types.Coin `protobuf:"bytes,4,rep,name=funds,proto3" json:"funds"`
}

func (m *MsgExecute) Reset()         { *m = MsgExecute{} }
func (m *MsgExecute) String() string { return proto.CompactTextString(m) }
func (*MsgExecute) ProtoMessage()    {}
func (*MsgExecute) Descriptor() ([]byte, []int) {
	return fileDescriptor_1acf5eaefb0cd3e7, []int{2}
}
func (m *MsgExecute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecute.Merge(m, src)
}
func (m *MsgExecute) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecute) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecute.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecute proto.InternalMessageInfo

func (m *MsgExecute) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgExecute) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgExecute) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *MsgExecute) GetFunds() []types.Coin {
	if m != nil {
		return m.Funds
	}
	return nil
}

type MsgExecuteResponse struct {
	Data *types1.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgExecuteResponse) Reset()         { *m = MsgExecuteResponse{} }
func (m *MsgExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteResponse) ProtoMessage()    {}
func (*MsgExecuteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1acf5eaefb0cd3e7, []int{3}
}
func (m *MsgExecuteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteResponse.Merge(m, src)
}
func (m *MsgExecuteResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteResponse proto.InternalMessageInfo

func (m *MsgExecuteResponse) GetData() *types1.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgDeploy)(nil), "accounts.accounts.MsgDeploy")
	proto.RegisterType((*MsgDeployResponse)(nil), "accounts.accounts.MsgDeployResponse")
	proto.RegisterType((*MsgExecute)(nil), "accounts.accounts.MsgExecute")
	proto.RegisterType((*MsgExecuteResponse)(nil), "accounts.accounts.MsgExecuteResponse")
}

func init() { proto.RegisterFile("accounts/accounts/tx.proto", fileDescriptor_1acf5eaefb0cd3e7) }

var fileDescriptor_1acf5eaefb0cd3e7 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xce, 0xa4, 0xb1, 0x65, 0x5f, 0x17, 0x61, 0x87, 0x45, 0xd2, 0xa0, 0xb1, 0x16, 0x85, 0x9c,
	0x26, 0x6c, 0x17, 0xaf, 0x82, 0xab, 0xde, 0x0c, 0x42, 0x8e, 0x5e, 0x64, 0x92, 0xcc, 0x0e, 0xc1,
	0xdd, 0x99, 0xd0, 0x97, 0x48, 0xf3, 0x2b, 0xd4, 0xff, 0xe0, 0x8f, 0xe9, 0xb1, 0x47, 0x4f, 0x22,
	0xed, 0x1f, 0x91, 0x26, 0x99, 0xd8, 0xa2, 0x56, 0xd8, 0xdb, 0xf7, 0xbe, 0xf7, 0x92, 0xf7, 0x7d,
	0x5f, 0xf2, 0xc0, 0xe3, 0x69, 0xaa, 0x2b, 0x55, 0x62, 0xd8, 0x83, 0x72, 0xc9, 0x8a, 0x85, 0x2e,
	0x35, 0x3d, 0x33, 0x14, 0x33, 0xc0, 0x9b, 0x48, 0xad, 0xe5, 0x8d, 0x08, 0x9b, 0x81, 0xa4, 0xba,
	0x0e, 0xb9, 0xaa, 0xdb, 0x69, 0xcf, 0x4f, 0x35, 0xde, 0x6a, 0x0c, 0x13, 0x8e, 0x22, 0xfc, 0x74,
	0x91, 0x88, 0x92, 0x5f, 0x84, 0xa9, 0xce, 0x55, 0xd7, 0x3f, 0x97, 0x5a, 0xea, 0x06, 0x86, 0x3b,
	0xd4, 0xb2, 0xb3, 0xaf, 0x04, 0x4e, 0x22, 0x94, 0xaf, 0x45, 0x71, 0xa3, 0x6b, 0xfa, 0x00, 0x86,
	0x28, 0x54, 0x26, 0x16, 0x2e, 0x99, 0x92, 0xe0, 0x24, 0xee, 0x2a, 0x4a, 0xc1, 0xf9, 0x98, 0xab,
	0xcc, 0xb5, 0x1b, 0xb6, 0xc1, 0xf4, 0x09, 0x9c, 0xe6, 0x2a, 0x2f, 0x3f, 0xdc, 0x0a, 0x44, 0x2e,
	0x85, 0x3b, 0x98, 0x92, 0xe0, 0x34, 0x1e, 0xef, 0xb8, 0xa8, 0xa5, 0xe8, 0x73, 0xb8, 0x77, 0x5d,
	0xa9, 0x0c, 0x5d, 0x67, 0x3a, 0x08, 0xc6, 0xf3, 0x09, 0x6b, 0x25, 0xb2, 0x9d, 0x44, 0xd6, 0x49,
	0x64, 0xaf, 0x74, 0xae, 0xae, 0x9c, 0xd5, 0x8f, 0xc7, 0x56, 0xdc, 0x4e, 0xcf, 0x24, 0x9c, 0xf5,
	0x92, 0x62, 0x81, 0x85, 0x56, 0x28, 0xa8, 0x0b, 0x23, 0x9e, 0x65, 0x0b, 0x81, 0xd8, 0x69, 0x33,
	0x25, 0xbd, 0x0f, 0x76, 0xde, 0x4a, 0x73, 0x62, 0x3b, 0xcf, 0x68, 0x00, 0x4e, 0xc6, 0x4b, 0xde,
	0x08, 0x1a, 0xcf, 0xcf, 0x59, 0x1b, 0x19, 0x33, 0x91, 0xb1, 0x97, 0xaa, 0x8e, 0x9b, 0x89, 0xd9,
	0x67, 0x02, 0x10, 0xa1, 0x7c, 0xb3, 0x14, 0x69, 0x55, 0x8a, 0x7f, 0xba, 0xdf, 0x5b, 0x6d, 0x1f,
	0xae, 0x76, 0x61, 0x74, 0x68, 0xdf, 0x94, 0x77, 0xb5, 0xfe, 0x02, 0xe8, 0x6f, 0x41, 0xbd, 0x77,
	0xe3, 0x88, 0xfc, 0xcf, 0xd1, 0xfc, 0x1b, 0x81, 0x41, 0x84, 0x92, 0xbe, 0x85, 0x61, 0xf7, 0x49,
	0x1f, 0xb2, 0x3f, 0xfe, 0x22, 0xd6, 0xa7, 0xeb, 0x3d, 0x3d, 0xd6, 0xed, 0xf7, 0xbf, 0x83, 0x91,
	0xc9, 0xe8, 0xd1, 0xdf, 0x1f, 0xe8, 0xda, 0xde, 0xb3, 0xa3, 0x6d, 0xf3, 0xc2, 0xab, 0xcb, 0xd5,
	0xc6, 0x27, 0xeb, 0x8d, 0x4f, 0x7e, 0x6e, 0x7c, 0xf2, 0x65, 0xeb, 0x5b, 0xeb, 0xad, 0x6f, 0x7d,
	0xdf, 0xfa, 0xd6, 0xfb, 0x49, 0x7f, 0x06, 0xcb, 0xbd, 0x8b, 0xa8, 0x0b, 0x81, 0xc9, 0xb0, 0xf1,
	0x7b, 0xf9, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x82, 0x39, 0x3e, 0xc4, 0x33, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Deploy(ctx context.Context, in *MsgDeploy, opts ...grpc.CallOption) (*MsgDeployResponse, error)
	Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Deploy(ctx context.Context, in *MsgDeploy, opts ...grpc.CallOption) (*MsgDeployResponse, error) {
	out := new(MsgDeployResponse)
	err := c.cc.Invoke(ctx, "/accounts.accounts.Msg/Deploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error) {
	out := new(MsgExecuteResponse)
	err := c.cc.Invoke(ctx, "/accounts.accounts.Msg/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Deploy(context.Context, *MsgDeploy) (*MsgDeployResponse, error)
	Execute(context.Context, *MsgExecute) (*MsgExecuteResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Deploy(ctx context.Context, req *MsgDeploy) (*MsgDeployResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (*UnimplementedMsgServer) Execute(ctx context.Context, req *MsgExecute) (*MsgExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeploy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.accounts.Msg/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deploy(ctx, req.(*MsgDeploy))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.accounts.Msg/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Execute(ctx, req.(*MsgExecute))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.accounts.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deploy",
			Handler:    _Msg_Deploy_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _Msg_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts/accounts/tx.proto",
}

func (m *MsgDeploy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeploy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeploy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Funds) > 0 {
		for iNdEx := len(m.Funds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Funds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InitMessage) > 0 {
		i -= len(m.InitMessage)
		copy(dAtA[i:], m.InitMessage)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InitMessage)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(dAtA[i:], m.Kind)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Kind)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeployResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeployResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeployResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Funds) > 0 {
		for iNdEx := len(m.Funds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Funds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecuteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgDeploy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Kind)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.InitMessage)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Funds) > 0 {
		for _, e := range m.Funds {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgDeployResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgExecute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Funds) > 0 {
		for _, e := range m.Funds {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgExecuteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDeploy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeploy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeploy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitMessage", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitMessage = append(m.InitMessage[:0], dAtA[iNdEx:postIndex]...)
			if m.InitMessage == nil {
				m.InitMessage = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funds = append(m.Funds, types.Coin{})
			if err := m.Funds[len(m.Funds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeployResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeployResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeployResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types1.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgExecute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgExecute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message[:0], dAtA[iNdEx:postIndex]...)
			if m.Message == nil {
				m.Message = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funds = append(m.Funds, types.Coin{})
			if err := m.Funds[len(m.Funds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgExecuteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgExecuteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types1.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
