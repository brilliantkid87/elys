// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/amm/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

type MsgCreatePool struct {
	Creator         string       `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Weights         []types.Coin `protobuf:"bytes,2,rep,name=weights,proto3" json:"weights"`
	InitialDeposit  []types.Coin `protobuf:"bytes,3,rep,name=initialDeposit,proto3" json:"initialDeposit"`
	SwapFee         uint64       `protobuf:"varint,4,opt,name=swapFee,proto3" json:"swapFee,omitempty"`
	ExitFee         uint64       `protobuf:"varint,5,opt,name=exitFee,proto3" json:"exitFee,omitempty"`
	FutureGovernor  string       `protobuf:"bytes,6,opt,name=futureGovernor,proto3" json:"futureGovernor,omitempty"`
	ScallingFactors []uint64     `protobuf:"varint,7,rep,packed,name=scallingFactors,proto3" json:"scallingFactors,omitempty"`
}

func (m *MsgCreatePool) Reset()         { *m = MsgCreatePool{} }
func (m *MsgCreatePool) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePool) ProtoMessage()    {}
func (*MsgCreatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7ddafd861f6b7f, []int{0}
}
func (m *MsgCreatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePool.Merge(m, src)
}
func (m *MsgCreatePool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePool proto.InternalMessageInfo

func (m *MsgCreatePool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePool) GetWeights() []types.Coin {
	if m != nil {
		return m.Weights
	}
	return nil
}

func (m *MsgCreatePool) GetInitialDeposit() []types.Coin {
	if m != nil {
		return m.InitialDeposit
	}
	return nil
}

func (m *MsgCreatePool) GetSwapFee() uint64 {
	if m != nil {
		return m.SwapFee
	}
	return 0
}

func (m *MsgCreatePool) GetExitFee() uint64 {
	if m != nil {
		return m.ExitFee
	}
	return 0
}

func (m *MsgCreatePool) GetFutureGovernor() string {
	if m != nil {
		return m.FutureGovernor
	}
	return ""
}

func (m *MsgCreatePool) GetScallingFactors() []uint64 {
	if m != nil {
		return m.ScallingFactors
	}
	return nil
}

type MsgCreatePoolResponse struct {
}

func (m *MsgCreatePoolResponse) Reset()         { *m = MsgCreatePoolResponse{} }
func (m *MsgCreatePoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePoolResponse) ProtoMessage()    {}
func (*MsgCreatePoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7ddafd861f6b7f, []int{1}
}
func (m *MsgCreatePoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePoolResponse.Merge(m, src)
}
func (m *MsgCreatePoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePoolResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreatePool)(nil), "elysnetwork.elys.amm.MsgCreatePool")
	proto.RegisterType((*MsgCreatePoolResponse)(nil), "elysnetwork.elys.amm.MsgCreatePoolResponse")
}

func init() { proto.RegisterFile("elys/amm/tx.proto", fileDescriptor_ed7ddafd861f6b7f) }

var fileDescriptor_ed7ddafd861f6b7f = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xca, 0xd3, 0x40,
	0x14, 0xc5, 0x93, 0x26, 0xb6, 0x38, 0x62, 0xc5, 0x50, 0x31, 0x76, 0x11, 0x43, 0x05, 0x09, 0x88,
	0x33, 0xb4, 0xae, 0xdc, 0xb6, 0xd2, 0xae, 0x0a, 0x92, 0xa5, 0x0b, 0x61, 0x12, 0xae, 0xe9, 0x60,
	0x92, 0x1b, 0x66, 0xa6, 0xff, 0xde, 0xc2, 0x77, 0xf1, 0x25, 0xba, 0xec, 0xd2, 0x95, 0x48, 0xfb,
	0x22, 0x32, 0x49, 0x03, 0x5f, 0xcb, 0xb7, 0xe8, 0xee, 0x9e, 0x7b, 0x4e, 0x7e, 0x99, 0xb9, 0x73,
	0xc9, 0x4b, 0xc8, 0xf7, 0x8a, 0xf1, 0xa2, 0x60, 0x7a, 0x47, 0x2b, 0x89, 0x1a, 0xbd, 0x81, 0x69,
	0x95, 0xa0, 0xb7, 0x28, 0x7f, 0x52, 0x53, 0x53, 0x5e, 0x14, 0xc3, 0x41, 0x86, 0x19, 0xd6, 0x01,
	0x66, 0xaa, 0x26, 0x3b, 0x0c, 0x52, 0x54, 0x05, 0x2a, 0x96, 0x70, 0x05, 0x6c, 0x33, 0x4e, 0x40,
	0xf3, 0x31, 0x4b, 0x51, 0x94, 0x8d, 0x3f, 0xfa, 0xdd, 0x21, 0xcf, 0x97, 0x2a, 0x9b, 0x49, 0xe0,
	0x1a, 0xbe, 0x22, 0xe6, 0x9e, 0x4f, 0x7a, 0xa9, 0x51, 0x28, 0x7d, 0x3b, 0xb4, 0xa3, 0xa7, 0x71,
	0x2b, 0xbd, 0xcf, 0xa4, 0xb7, 0x05, 0x91, 0xad, 0xb4, 0xf2, 0x3b, 0xa1, 0x13, 0x3d, 0x9b, 0xbc,
	0xa1, 0x0d, 0x9d, 0x1a, 0x3a, 0xbd, 0xd0, 0xe9, 0x0c, 0x45, 0x39, 0x75, 0x0f, 0x7f, 0xdf, 0x5a,
	0x71, 0x9b, 0xf7, 0x16, 0xa4, 0x2f, 0x4a, 0xa1, 0x05, 0xcf, 0xbf, 0x40, 0x85, 0x4a, 0x68, 0xdf,
	0xb9, 0x8f, 0x70, 0xf3, 0x99, 0x39, 0x9d, 0xda, 0xf2, 0x6a, 0x0e, 0xe0, 0xbb, 0xa1, 0x1d, 0xb9,
	0x71, 0x2b, 0x8d, 0x03, 0x3b, 0xa1, 0x8d, 0xf3, 0xa4, 0x71, 0x2e, 0xd2, 0x7b, 0x4f, 0xfa, 0x3f,
	0xd6, 0x7a, 0x2d, 0x61, 0x81, 0x1b, 0x90, 0x25, 0x4a, 0xbf, 0x5b, 0x5f, 0xec, 0xa6, 0xeb, 0x45,
	0xe4, 0x85, 0x4a, 0x79, 0x9e, 0x8b, 0x32, 0x9b, 0xf3, 0x54, 0xa3, 0x54, 0x7e, 0x2f, 0x74, 0x22,
	0x37, 0xbe, 0x6d, 0x8f, 0x5e, 0x93, 0x57, 0x57, 0x43, 0x8b, 0x41, 0x55, 0x58, 0x2a, 0x98, 0x00,
	0x71, 0x96, 0x2a, 0xf3, 0xbe, 0x13, 0xf2, 0x60, 0xa2, 0xef, 0xe8, 0x63, 0x0f, 0x46, 0xaf, 0x08,
	0xc3, 0x0f, 0x77, 0x84, 0xda, 0xdf, 0x4c, 0xa7, 0x87, 0x53, 0x60, 0x1f, 0x4f, 0x81, 0xfd, 0xef,
	0x14, 0xd8, 0xbf, 0xce, 0x81, 0x75, 0x3c, 0x07, 0xd6, 0x9f, 0x73, 0x60, 0x7d, 0x8b, 0x32, 0xa1,
	0x57, 0xeb, 0x84, 0xa6, 0x58, 0x30, 0x03, 0xf9, 0x78, 0x21, 0xd6, 0x82, 0xed, 0x9a, 0x45, 0xda,
	0x57, 0xa0, 0x92, 0x6e, 0xbd, 0x00, 0x9f, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x58, 0xe6,
	0x35, 0x61, 0x02, 0x00, 0x00,
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
	CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error) {
	out := new(MsgCreatePoolResponse)
	err := c.cc.Invoke(ctx, "/elysnetwork.elys.amm.Msg/CreatePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePool(ctx context.Context, req *MsgCreatePool) (*MsgCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elysnetwork.elys.amm.Msg/CreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePool(ctx, req.(*MsgCreatePool))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elysnetwork.elys.amm.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePool",
			Handler:    _Msg_CreatePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "elys/amm/tx.proto",
}

func (m *MsgCreatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ScallingFactors) > 0 {
		dAtA2 := make([]byte, len(m.ScallingFactors)*10)
		var j1 int
		for _, num := range m.ScallingFactors {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTx(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FutureGovernor) > 0 {
		i -= len(m.FutureGovernor)
		copy(dAtA[i:], m.FutureGovernor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FutureGovernor)))
		i--
		dAtA[i] = 0x32
	}
	if m.ExitFee != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ExitFee))
		i--
		dAtA[i] = 0x28
	}
	if m.SwapFee != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SwapFee))
		i--
		dAtA[i] = 0x20
	}
	if len(m.InitialDeposit) > 0 {
		for iNdEx := len(m.InitialDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Weights) > 0 {
		for iNdEx := len(m.Weights) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Weights[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgCreatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Weights) > 0 {
		for _, e := range m.Weights {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.InitialDeposit) > 0 {
		for _, e := range m.InitialDeposit {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.SwapFee != 0 {
		n += 1 + sovTx(uint64(m.SwapFee))
	}
	if m.ExitFee != 0 {
		n += 1 + sovTx(uint64(m.ExitFee))
	}
	l = len(m.FutureGovernor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ScallingFactors) > 0 {
		l = 0
		for _, e := range m.ScallingFactors {
			l += sovTx(uint64(e))
		}
		n += 1 + sovTx(uint64(l)) + l
	}
	return n
}

func (m *MsgCreatePoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreatePool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weights", wireType)
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
			m.Weights = append(m.Weights, types.Coin{})
			if err := m.Weights[len(m.Weights)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialDeposit", wireType)
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
			m.InitialDeposit = append(m.InitialDeposit, types.Coin{})
			if err := m.InitialDeposit[len(m.InitialDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			m.SwapFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SwapFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitFee", wireType)
			}
			m.ExitFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExitFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FutureGovernor", wireType)
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
			m.FutureGovernor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ScallingFactors = append(m.ScallingFactors, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ScallingFactors) == 0 {
					m.ScallingFactors = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ScallingFactors = append(m.ScallingFactors, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ScallingFactors", wireType)
			}
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
func (m *MsgCreatePoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreatePoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
