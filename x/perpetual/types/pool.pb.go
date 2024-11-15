// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/perpetual/pool.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type PoolAsset struct {
	Liabilities           cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=liabilities,proto3,customtype=cosmossdk.io/math.Int" json:"liabilities"`
	Custody               cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=custody,proto3,customtype=cosmossdk.io/math.Int" json:"custody"`
	TakeProfitLiabilities cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=take_profit_liabilities,json=takeProfitLiabilities,proto3,customtype=cosmossdk.io/math.Int" json:"take_profit_liabilities"`
	TakeProfitCustody     cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=take_profit_custody,json=takeProfitCustody,proto3,customtype=cosmossdk.io/math.Int" json:"take_profit_custody"`
	AssetDenom            string                `protobuf:"bytes,5,opt,name=asset_denom,json=assetDenom,proto3" json:"asset_denom,omitempty"`
	Collateral            cosmossdk_io_math.Int `protobuf:"bytes,6,opt,name=collateral,proto3,customtype=cosmossdk.io/math.Int" json:"collateral"`
}

func (m *PoolAsset) Reset()         { *m = PoolAsset{} }
func (m *PoolAsset) String() string { return proto.CompactTextString(m) }
func (*PoolAsset) ProtoMessage()    {}
func (*PoolAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2020803d9775cac, []int{0}
}
func (m *PoolAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolAsset.Merge(m, src)
}
func (m *PoolAsset) XXX_Size() int {
	return m.Size()
}
func (m *PoolAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolAsset.DiscardUnknown(m)
}

var xxx_messageInfo_PoolAsset proto.InternalMessageInfo

func (m *PoolAsset) GetAssetDenom() string {
	if m != nil {
		return m.AssetDenom
	}
	return ""
}

type Pool struct {
	AmmPoolId                            uint64                      `protobuf:"varint,1,opt,name=amm_pool_id,json=ammPoolId,proto3" json:"amm_pool_id,omitempty"`
	Health                               cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=health,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"health"`
	BorrowInterestRate                   cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=borrow_interest_rate,json=borrowInterestRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"borrow_interest_rate"`
	PoolAssetsLong                       []PoolAsset                 `protobuf:"bytes,4,rep,name=pool_assets_long,json=poolAssetsLong,proto3" json:"pool_assets_long"`
	PoolAssetsShort                      []PoolAsset                 `protobuf:"bytes,5,rep,name=pool_assets_short,json=poolAssetsShort,proto3" json:"pool_assets_short"`
	LastHeightBorrowInterestRateComputed int64                       `protobuf:"varint,6,opt,name=last_height_borrow_interest_rate_computed,json=lastHeightBorrowInterestRateComputed,proto3" json:"last_height_borrow_interest_rate_computed,omitempty"`
	// funding rate, if positive longs pay shorts, if negative shorts pay longs
	FundingRate   cosmossdk_io_math.LegacyDec `protobuf:"bytes,7,opt,name=funding_rate,json=fundingRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"funding_rate"`
	FeesCollected []types.Coin                `protobuf:"bytes,8,rep,name=fees_collected,json=feesCollected,proto3" json:"fees_collected"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2020803d9775cac, []int{1}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetAmmPoolId() uint64 {
	if m != nil {
		return m.AmmPoolId
	}
	return 0
}

func (m *Pool) GetPoolAssetsLong() []PoolAsset {
	if m != nil {
		return m.PoolAssetsLong
	}
	return nil
}

func (m *Pool) GetPoolAssetsShort() []PoolAsset {
	if m != nil {
		return m.PoolAssetsShort
	}
	return nil
}

func (m *Pool) GetLastHeightBorrowInterestRateComputed() int64 {
	if m != nil {
		return m.LastHeightBorrowInterestRateComputed
	}
	return 0
}

func (m *Pool) GetFeesCollected() []types.Coin {
	if m != nil {
		return m.FeesCollected
	}
	return nil
}

func init() {
	proto.RegisterType((*PoolAsset)(nil), "elys.perpetual.PoolAsset")
	proto.RegisterType((*Pool)(nil), "elys.perpetual.Pool")
}

func init() { proto.RegisterFile("elys/perpetual/pool.proto", fileDescriptor_f2020803d9775cac) }

var fileDescriptor_f2020803d9775cac = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x4f, 0xd4, 0x4e,
	0x14, 0xdf, 0xfd, 0x6f, 0x81, 0x3f, 0xb3, 0x8a, 0x32, 0x42, 0x2c, 0x18, 0x0b, 0x41, 0x0f, 0x78,
	0xb0, 0x0d, 0x7a, 0xf0, 0x60, 0x8c, 0x71, 0x97, 0x10, 0x36, 0x62, 0x42, 0x6a, 0x88, 0x89, 0x97,
	0xc9, 0xb4, 0x7d, 0xb4, 0x13, 0xa6, 0x7d, 0x4d, 0x67, 0x56, 0xdc, 0xcf, 0xe0, 0xc5, 0xef, 0xe3,
	0x17, 0xe0, 0xc8, 0xd1, 0x78, 0x20, 0x66, 0xf7, 0x8b, 0x98, 0x69, 0xbb, 0x50, 0x83, 0x87, 0xde,
	0xda, 0x79, 0xef, 0xf7, 0x7b, 0xbf, 0xf7, 0x7b, 0x6f, 0x86, 0x6c, 0x80, 0x9c, 0x28, 0x2f, 0x87,
	0x22, 0x07, 0x3d, 0xe6, 0xd2, 0xcb, 0x11, 0xa5, 0x9b, 0x17, 0xa8, 0x91, 0xae, 0x98, 0x90, 0x7b,
	0x1d, 0xda, 0x5c, 0x8b, 0x31, 0xc6, 0x32, 0xe4, 0x99, 0xaf, 0x2a, 0x6b, 0xd3, 0x09, 0x51, 0xa5,
	0xa8, 0xbc, 0x80, 0x2b, 0xf0, 0xbe, 0xec, 0x05, 0xa0, 0xf9, 0x9e, 0x17, 0xa2, 0xc8, 0xaa, 0xf8,
	0xce, 0xb7, 0x1e, 0x59, 0x3e, 0x46, 0x94, 0xef, 0x94, 0x02, 0x4d, 0xdf, 0x92, 0xbe, 0x14, 0x3c,
	0x10, 0x52, 0x68, 0x01, 0xca, 0xee, 0x6e, 0x77, 0x77, 0x97, 0x07, 0x8f, 0x2f, 0xae, 0xb6, 0x3a,
	0xbf, 0xae, 0xb6, 0xd6, 0x2b, 0x2a, 0x15, 0x9d, 0xb9, 0x02, 0xbd, 0x94, 0xeb, 0xc4, 0x1d, 0x65,
	0xda, 0x6f, 0x22, 0xe8, 0x2b, 0xb2, 0x14, 0x8e, 0x95, 0xc6, 0x68, 0x62, 0xff, 0xd7, 0x06, 0x3c,
	0xcf, 0xa6, 0x27, 0xe4, 0xa1, 0xe6, 0x67, 0xc0, 0xf2, 0x02, 0x4f, 0x85, 0x66, 0x4d, 0x15, 0xbd,
	0x36, 0x44, 0xeb, 0x06, 0x7d, 0x5c, 0x82, 0x8f, 0x1a, 0x7a, 0x3e, 0x90, 0x07, 0x4d, 0xda, 0xb9,
	0x36, 0xab, 0x0d, 0xe5, 0xea, 0x0d, 0xe5, 0xb0, 0x56, 0xb9, 0x45, 0xfa, 0xdc, 0x18, 0xc5, 0x22,
	0xc8, 0x30, 0xb5, 0x17, 0x0c, 0x8d, 0x4f, 0xca, 0xa3, 0x7d, 0x73, 0x42, 0xdf, 0x10, 0x12, 0xa2,
	0x94, 0x5c, 0x43, 0xc1, 0xa5, 0xbd, 0xd8, 0xa6, 0x4c, 0x03, 0xb0, 0xf3, 0xc3, 0x22, 0x96, 0x99,
	0x06, 0x75, 0x48, 0x9f, 0xa7, 0x29, 0x33, 0xe3, 0x66, 0x22, 0x2a, 0x07, 0x61, 0xf9, 0xcb, 0x3c,
	0x4d, 0x4d, 0x74, 0x14, 0xd1, 0xd7, 0x64, 0x31, 0x01, 0x2e, 0x75, 0x52, 0xdb, 0xfc, 0xa4, 0xae,
	0xf1, 0xe8, 0x76, 0x8d, 0x23, 0x88, 0x79, 0x38, 0xd9, 0x87, 0xd0, 0xaf, 0x21, 0xf4, 0x84, 0xac,
	0x05, 0x58, 0x14, 0x78, 0xce, 0x44, 0xa6, 0xa1, 0x00, 0xa5, 0x59, 0xc1, 0x35, 0xd4, 0x46, 0xb7,
	0xa2, 0xa2, 0x15, 0xc1, 0xa8, 0xc6, 0xfb, 0x5c, 0x03, 0x1d, 0x91, 0xfb, 0xa5, 0xde, 0xd2, 0x0e,
	0xc5, 0x24, 0x66, 0xb1, 0x6d, 0x6d, 0xf7, 0x76, 0xfb, 0x2f, 0x36, 0xdc, 0xbf, 0x77, 0xd5, 0xbd,
	0xde, 0xb8, 0x81, 0x65, 0xaa, 0xf9, 0x2b, 0xf9, 0xfc, 0x40, 0x1d, 0x61, 0x16, 0xd3, 0xf7, 0x64,
	0xb5, 0x49, 0xa5, 0x12, 0x2c, 0xb4, 0xbd, 0xd0, 0x8e, 0xeb, 0xde, 0x0d, 0xd7, 0x47, 0x83, 0xa3,
	0x9f, 0xc8, 0x33, 0xc9, 0x95, 0x66, 0x09, 0x88, 0x38, 0xd1, 0xec, 0x5f, 0xad, 0xb3, 0x10, 0xd3,
	0x7c, 0xac, 0x21, 0x2a, 0x47, 0xd6, 0xf3, 0x9f, 0x1a, 0xc0, 0x61, 0x99, 0x3f, 0xb8, 0xd5, 0xe8,
	0xb0, 0xce, 0xa5, 0x07, 0xe4, 0xce, 0xe9, 0x38, 0x8b, 0x44, 0x16, 0x57, 0xfe, 0x2d, 0xb5, 0xf7,
	0xaf, 0x5f, 0x03, 0x4b, 0xe3, 0x0e, 0xc8, 0xca, 0x29, 0x80, 0x62, 0x66, 0x11, 0x20, 0x34, 0x2a,
	0xfe, 0xaf, 0x5b, 0xad, 0x28, 0x5c, 0x73, 0x79, 0xdd, 0xfa, 0xf2, 0xba, 0x43, 0x14, 0x59, 0xdd,
	0xea, 0x5d, 0x03, 0x1b, 0xce, 0x51, 0x83, 0xc3, 0x8b, 0xa9, 0xd3, 0xbd, 0x9c, 0x3a, 0xdd, 0xdf,
	0x53, 0xa7, 0xfb, 0x7d, 0xe6, 0x74, 0x2e, 0x67, 0x4e, 0xe7, 0xe7, 0xcc, 0xe9, 0x7c, 0x76, 0x63,
	0xa1, 0x93, 0x71, 0xe0, 0x86, 0x98, 0x7a, 0xc6, 0xbe, 0xe7, 0x19, 0xe8, 0x73, 0x2c, 0xce, 0xca,
	0x1f, 0xef, 0x6b, 0xe3, 0x81, 0xd1, 0x93, 0x1c, 0x54, 0xb0, 0x58, 0x3e, 0x0e, 0x2f, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x84, 0x55, 0x21, 0x22, 0x7f, 0x04, 0x00, 0x00,
}

func (m *PoolAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Collateral.Size()
		i -= size
		if _, err := m.Collateral.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.AssetDenom) > 0 {
		i -= len(m.AssetDenom)
		copy(dAtA[i:], m.AssetDenom)
		i = encodeVarintPool(dAtA, i, uint64(len(m.AssetDenom)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.TakeProfitCustody.Size()
		i -= size
		if _, err := m.TakeProfitCustody.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TakeProfitLiabilities.Size()
		i -= size
		if _, err := m.TakeProfitLiabilities.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Custody.Size()
		i -= size
		if _, err := m.Custody.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Liabilities.Size()
		i -= size
		if _, err := m.Liabilities.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeesCollected) > 0 {
		for iNdEx := len(m.FeesCollected) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeesCollected[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	{
		size := m.FundingRate.Size()
		i -= size
		if _, err := m.FundingRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.LastHeightBorrowInterestRateComputed != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.LastHeightBorrowInterestRateComputed))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PoolAssetsShort) > 0 {
		for iNdEx := len(m.PoolAssetsShort) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssetsShort[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PoolAssetsLong) > 0 {
		for iNdEx := len(m.PoolAssetsLong) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssetsLong[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.BorrowInterestRate.Size()
		i -= size
		if _, err := m.BorrowInterestRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Health.Size()
		i -= size
		if _, err := m.Health.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AmmPoolId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.AmmPoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Liabilities.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.Custody.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.TakeProfitLiabilities.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.TakeProfitCustody.Size()
	n += 1 + l + sovPool(uint64(l))
	l = len(m.AssetDenom)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.Collateral.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AmmPoolId != 0 {
		n += 1 + sovPool(uint64(m.AmmPoolId))
	}
	l = m.Health.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.BorrowInterestRate.Size()
	n += 1 + l + sovPool(uint64(l))
	if len(m.PoolAssetsLong) > 0 {
		for _, e := range m.PoolAssetsLong {
			l = e.Size()
			n += 1 + l + sovPool(uint64(l))
		}
	}
	if len(m.PoolAssetsShort) > 0 {
		for _, e := range m.PoolAssetsShort {
			l = e.Size()
			n += 1 + l + sovPool(uint64(l))
		}
	}
	if m.LastHeightBorrowInterestRateComputed != 0 {
		n += 1 + sovPool(uint64(m.LastHeightBorrowInterestRateComputed))
	}
	l = m.FundingRate.Size()
	n += 1 + l + sovPool(uint64(l))
	if len(m.FeesCollected) > 0 {
		for _, e := range m.FeesCollected {
			l = e.Size()
			n += 1 + l + sovPool(uint64(l))
		}
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: PoolAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liabilities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Liabilities.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Custody", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Custody.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakeProfitLiabilities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TakeProfitLiabilities.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakeProfitCustody", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TakeProfitCustody.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Collateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmmPoolId", wireType)
			}
			m.AmmPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AmmPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Health", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Health.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowInterestRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BorrowInterestRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssetsLong", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssetsLong = append(m.PoolAssetsLong, PoolAsset{})
			if err := m.PoolAssetsLong[len(m.PoolAssetsLong)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssetsShort", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssetsShort = append(m.PoolAssetsShort, PoolAsset{})
			if err := m.PoolAssetsShort[len(m.PoolAssetsShort)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeightBorrowInterestRateComputed", wireType)
			}
			m.LastHeightBorrowInterestRateComputed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastHeightBorrowInterestRateComputed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundingRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FundingRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeesCollected", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeesCollected = append(m.FeesCollected, types.Coin{})
			if err := m.FeesCollected[len(m.FeesCollected)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
