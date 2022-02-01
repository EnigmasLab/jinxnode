// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/margin/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MTPRequest struct {
	Address         string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CustodyAsset    string `protobuf:"bytes,2,opt,name=custody_asset,json=custodyAsset,proto3" json:"custody_asset,omitempty"`
	CollateralAsset string `protobuf:"bytes,3,opt,name=collateral_asset,json=collateralAsset,proto3" json:"collateral_asset,omitempty"`
	Position        string `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
}

func (m *MTPRequest) Reset()         { *m = MTPRequest{} }
func (m *MTPRequest) String() string { return proto.CompactTextString(m) }
func (*MTPRequest) ProtoMessage()    {}
func (*MTPRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73c14070fed1f663, []int{0}
}
func (m *MTPRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MTPRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MTPRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MTPRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MTPRequest.Merge(m, src)
}
func (m *MTPRequest) XXX_Size() int {
	return m.Size()
}
func (m *MTPRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MTPRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MTPRequest proto.InternalMessageInfo

func (m *MTPRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MTPRequest) GetCustodyAsset() string {
	if m != nil {
		return m.CustodyAsset
	}
	return ""
}

func (m *MTPRequest) GetCollateralAsset() string {
	if m != nil {
		return m.CollateralAsset
	}
	return ""
}

func (m *MTPRequest) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

type MTPResponse struct {
	Mtp *MTP `protobuf:"bytes,1,opt,name=mtp,proto3" json:"mtp,omitempty"`
}

func (m *MTPResponse) Reset()         { *m = MTPResponse{} }
func (m *MTPResponse) String() string { return proto.CompactTextString(m) }
func (*MTPResponse) ProtoMessage()    {}
func (*MTPResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73c14070fed1f663, []int{1}
}
func (m *MTPResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MTPResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MTPResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MTPResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MTPResponse.Merge(m, src)
}
func (m *MTPResponse) XXX_Size() int {
	return m.Size()
}
func (m *MTPResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MTPResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MTPResponse proto.InternalMessageInfo

func (m *MTPResponse) GetMtp() *MTP {
	if m != nil {
		return m.Mtp
	}
	return nil
}

type PositionsForAddressRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *PositionsForAddressRequest) Reset()         { *m = PositionsForAddressRequest{} }
func (m *PositionsForAddressRequest) String() string { return proto.CompactTextString(m) }
func (*PositionsForAddressRequest) ProtoMessage()    {}
func (*PositionsForAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73c14070fed1f663, []int{2}
}
func (m *PositionsForAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PositionsForAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PositionsForAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PositionsForAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionsForAddressRequest.Merge(m, src)
}
func (m *PositionsForAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *PositionsForAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionsForAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PositionsForAddressRequest proto.InternalMessageInfo

func (m *PositionsForAddressRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type PositionsForAddressResponse struct {
	Mtps []*MTP `protobuf:"bytes,1,rep,name=mtps,proto3" json:"mtps,omitempty"`
}

func (m *PositionsForAddressResponse) Reset()         { *m = PositionsForAddressResponse{} }
func (m *PositionsForAddressResponse) String() string { return proto.CompactTextString(m) }
func (*PositionsForAddressResponse) ProtoMessage()    {}
func (*PositionsForAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73c14070fed1f663, []int{3}
}
func (m *PositionsForAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PositionsForAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PositionsForAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PositionsForAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionsForAddressResponse.Merge(m, src)
}
func (m *PositionsForAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *PositionsForAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionsForAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PositionsForAddressResponse proto.InternalMessageInfo

func (m *PositionsForAddressResponse) GetMtps() []*MTP {
	if m != nil {
		return m.Mtps
	}
	return nil
}

func init() {
	proto.RegisterType((*MTPRequest)(nil), "sifnode.margin.v1.MTPRequest")
	proto.RegisterType((*MTPResponse)(nil), "sifnode.margin.v1.MTPResponse")
	proto.RegisterType((*PositionsForAddressRequest)(nil), "sifnode.margin.v1.PositionsForAddressRequest")
	proto.RegisterType((*PositionsForAddressResponse)(nil), "sifnode.margin.v1.PositionsForAddressResponse")
}

func init() { proto.RegisterFile("sifnode/margin/v1/query.proto", fileDescriptor_73c14070fed1f663) }

var fileDescriptor_73c14070fed1f663 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0xe2, 0x40,
	0x14, 0xc7, 0xdb, 0x85, 0x65, 0x77, 0x87, 0xdd, 0xec, 0xee, 0xc4, 0x90, 0xa6, 0x4a, 0x63, 0xea,
	0x05, 0x49, 0xec, 0x04, 0x4c, 0xf4, 0x8c, 0x31, 0x12, 0x0e, 0x24, 0x88, 0x9c, 0xbc, 0x98, 0xa1,
	0x1d, 0xca, 0x24, 0x6d, 0xa7, 0x74, 0xa6, 0xc4, 0x7e, 0x0b, 0xe3, 0xa7, 0xf2, 0x62, 0xc2, 0xd1,
	0xa3, 0x81, 0x2f, 0x62, 0x3a, 0x6d, 0xf5, 0x60, 0x51, 0x6f, 0x9d, 0xf7, 0xff, 0xbd, 0xbe, 0x5f,
	0xa7, 0x0f, 0x34, 0x39, 0x9d, 0x05, 0xcc, 0x21, 0xc8, 0xc7, 0x91, 0x4b, 0x03, 0xb4, 0xec, 0xa0,
	0x45, 0x4c, 0xa2, 0xc4, 0x0a, 0x23, 0x26, 0x18, 0xfc, 0x9f, 0xc7, 0x56, 0x16, 0x5b, 0xcb, 0x8e,
	0xbe, 0xe3, 0x32, 0x97, 0xc9, 0x14, 0xa5, 0x4f, 0x19, 0xa8, 0xef, 0xb9, 0x8c, 0xb9, 0x1e, 0x41,
	0x38, 0xa4, 0x08, 0x07, 0x01, 0x13, 0x58, 0x50, 0x16, 0xf0, 0x3c, 0x2d, 0x99, 0x22, 0x92, 0x90,
	0xe4, 0xb1, 0x79, 0xaf, 0x02, 0x30, 0x9c, 0x8c, 0xc6, 0x64, 0x11, 0x13, 0x2e, 0xa0, 0x06, 0x7e,
	0x60, 0xc7, 0x89, 0x08, 0xe7, 0x9a, 0xba, 0xaf, 0xb6, 0x7e, 0x8d, 0x8b, 0x23, 0x3c, 0x00, 0x7f,
	0xec, 0x98, 0x0b, 0xe6, 0x24, 0x37, 0x98, 0x73, 0x22, 0xb4, 0x6f, 0x32, 0xff, 0x9d, 0x17, 0x7b,
	0x69, 0x0d, 0x1e, 0x82, 0x7f, 0x36, 0xf3, 0x3c, 0x2c, 0x48, 0x84, 0xbd, 0x9c, 0xab, 0x48, 0xee,
	0xef, 0x5b, 0x3d, 0x43, 0x75, 0xf0, 0x33, 0x64, 0x9c, 0xa6, 0xaa, 0x5a, 0x55, 0x22, 0xaf, 0x67,
	0xf3, 0x14, 0xd4, 0xa5, 0x13, 0x0f, 0x59, 0xc0, 0x09, 0x6c, 0x81, 0x8a, 0x2f, 0x42, 0x29, 0x54,
	0xef, 0x36, 0xac, 0x77, 0xf7, 0x62, 0xa5, 0x70, 0x8a, 0x98, 0x27, 0x40, 0x1f, 0xe5, 0x2f, 0xe1,
	0x17, 0x2c, 0xea, 0x65, 0xee, 0x9f, 0x7e, 0x9c, 0x39, 0x00, 0xbb, 0xa5, 0x7d, 0xb9, 0x40, 0x1b,
	0x54, 0x7d, 0x11, 0xa6, 0x5d, 0x95, 0x0f, 0x0c, 0x24, 0xd3, 0x7d, 0x54, 0xc1, 0xf7, 0xcb, 0xf4,
	0x37, 0xc2, 0x01, 0xa8, 0xf5, 0x89, 0x18, 0x4e, 0x46, 0xb0, 0xb9, 0xa5, 0x23, 0xf3, 0xd2, 0x8d,
	0x6d, 0x71, 0x36, 0xde, 0x54, 0x60, 0x02, 0x1a, 0x7d, 0x22, 0x4a, 0x14, 0xe1, 0x51, 0x49, 0xef,
	0xf6, 0x2b, 0xd0, 0xad, 0xaf, 0xe2, 0xc5, 0xe8, 0xb3, 0xf3, 0x87, 0xb5, 0xa1, 0xae, 0xd6, 0x86,
	0xfa, 0xbc, 0x36, 0xd4, 0xbb, 0x8d, 0xa1, 0xac, 0x36, 0x86, 0xf2, 0xb4, 0x31, 0x94, 0xeb, 0xb6,
	0x4b, 0xc5, 0x3c, 0x9e, 0x5a, 0x36, 0xf3, 0xd1, 0x15, 0x9d, 0xd9, 0x73, 0x4c, 0x03, 0x54, 0x6c,
	0xdb, 0x6d, 0xb1, 0x6f, 0x72, 0xd9, 0xa6, 0x35, 0xb9, 0x6d, 0xc7, 0x2f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x1c, 0x4c, 0x4c, 0xf8, 0xf4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	GetMTP(ctx context.Context, in *MTPRequest, opts ...grpc.CallOption) (*MTPResponse, error)
	GetPositionsForAddress(ctx context.Context, in *PositionsForAddressRequest, opts ...grpc.CallOption) (*PositionsForAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetMTP(ctx context.Context, in *MTPRequest, opts ...grpc.CallOption) (*MTPResponse, error) {
	out := new(MTPResponse)
	err := c.cc.Invoke(ctx, "/sifnode.margin.v1.Query/GetMTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetPositionsForAddress(ctx context.Context, in *PositionsForAddressRequest, opts ...grpc.CallOption) (*PositionsForAddressResponse, error) {
	out := new(PositionsForAddressResponse)
	err := c.cc.Invoke(ctx, "/sifnode.margin.v1.Query/GetPositionsForAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	GetMTP(context.Context, *MTPRequest) (*MTPResponse, error)
	GetPositionsForAddress(context.Context, *PositionsForAddressRequest) (*PositionsForAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetMTP(ctx context.Context, req *MTPRequest) (*MTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMTP not implemented")
}
func (*UnimplementedQueryServer) GetPositionsForAddress(ctx context.Context, req *PositionsForAddressRequest) (*PositionsForAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPositionsForAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetMTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetMTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sifnode.margin.v1.Query/GetMTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetMTP(ctx, req.(*MTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetPositionsForAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionsForAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetPositionsForAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sifnode.margin.v1.Query/GetPositionsForAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetPositionsForAddress(ctx, req.(*PositionsForAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sifnode.margin.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMTP",
			Handler:    _Query_GetMTP_Handler,
		},
		{
			MethodName: "GetPositionsForAddress",
			Handler:    _Query_GetPositionsForAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sifnode/margin/v1/query.proto",
}

func (m *MTPRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MTPRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MTPRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Position) > 0 {
		i -= len(m.Position)
		copy(dAtA[i:], m.Position)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Position)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CollateralAsset) > 0 {
		i -= len(m.CollateralAsset)
		copy(dAtA[i:], m.CollateralAsset)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CollateralAsset)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CustodyAsset) > 0 {
		i -= len(m.CustodyAsset)
		copy(dAtA[i:], m.CustodyAsset)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CustodyAsset)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MTPResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MTPResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MTPResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Mtp != nil {
		{
			size, err := m.Mtp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PositionsForAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PositionsForAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PositionsForAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PositionsForAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PositionsForAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PositionsForAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Mtps) > 0 {
		for iNdEx := len(m.Mtps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Mtps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MTPRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.CustodyAsset)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.CollateralAsset)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Position)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *MTPResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mtp != nil {
		l = m.Mtp.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *PositionsForAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *PositionsForAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Mtps) > 0 {
		for _, e := range m.Mtps {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MTPRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: MTPRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MTPRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustodyAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CustodyAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Position = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *MTPResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: MTPResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MTPResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mtp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Mtp == nil {
				m.Mtp = &MTP{}
			}
			if err := m.Mtp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *PositionsForAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: PositionsForAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PositionsForAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *PositionsForAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: PositionsForAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PositionsForAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mtps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mtps = append(m.Mtps, &MTP{})
			if err := m.Mtps[len(m.Mtps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
