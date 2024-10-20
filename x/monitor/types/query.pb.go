// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/monitor/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// QueryEndedEpochBtcHeightRequest defines a query type for EndedEpochBtcHeight
// RPC method
type QueryEndedEpochBtcHeightRequest struct {
	EpochNum uint64 `protobuf:"varint,1,opt,name=epoch_num,json=epochNum,proto3" json:"epoch_num,omitempty"`
}

func (m *QueryEndedEpochBtcHeightRequest) Reset()         { *m = QueryEndedEpochBtcHeightRequest{} }
func (m *QueryEndedEpochBtcHeightRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEndedEpochBtcHeightRequest) ProtoMessage()    {}
func (*QueryEndedEpochBtcHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8aafb034c55a8f2, []int{0}
}
func (m *QueryEndedEpochBtcHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEndedEpochBtcHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEndedEpochBtcHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEndedEpochBtcHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEndedEpochBtcHeightRequest.Merge(m, src)
}
func (m *QueryEndedEpochBtcHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEndedEpochBtcHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEndedEpochBtcHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEndedEpochBtcHeightRequest proto.InternalMessageInfo

func (m *QueryEndedEpochBtcHeightRequest) GetEpochNum() uint64 {
	if m != nil {
		return m.EpochNum
	}
	return 0
}

// QueryEndedEpochBtcHeightResponse defines a response type for
// EndedEpochBtcHeight RPC method
type QueryEndedEpochBtcHeightResponse struct {
	// height of btc light client when epoch ended
	BtcLightClientHeight uint32 `protobuf:"varint,1,opt,name=btc_light_client_height,json=btcLightClientHeight,proto3" json:"btc_light_client_height,omitempty"`
}

func (m *QueryEndedEpochBtcHeightResponse) Reset()         { *m = QueryEndedEpochBtcHeightResponse{} }
func (m *QueryEndedEpochBtcHeightResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEndedEpochBtcHeightResponse) ProtoMessage()    {}
func (*QueryEndedEpochBtcHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8aafb034c55a8f2, []int{1}
}
func (m *QueryEndedEpochBtcHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEndedEpochBtcHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEndedEpochBtcHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEndedEpochBtcHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEndedEpochBtcHeightResponse.Merge(m, src)
}
func (m *QueryEndedEpochBtcHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEndedEpochBtcHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEndedEpochBtcHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEndedEpochBtcHeightResponse proto.InternalMessageInfo

func (m *QueryEndedEpochBtcHeightResponse) GetBtcLightClientHeight() uint32 {
	if m != nil {
		return m.BtcLightClientHeight
	}
	return 0
}

// QueryReportedCheckpointBtcHeightRequest defines a query type for
// ReportedCheckpointBtcHeight RPC method
type QueryReportedCheckpointBtcHeightRequest struct {
	// ckpt_hash is hex encoded byte string of the hash of the checkpoint
	CkptHash string `protobuf:"bytes,1,opt,name=ckpt_hash,json=ckptHash,proto3" json:"ckpt_hash,omitempty"`
}

func (m *QueryReportedCheckpointBtcHeightRequest) Reset() {
	*m = QueryReportedCheckpointBtcHeightRequest{}
}
func (m *QueryReportedCheckpointBtcHeightRequest) String() string { return proto.CompactTextString(m) }
func (*QueryReportedCheckpointBtcHeightRequest) ProtoMessage()    {}
func (*QueryReportedCheckpointBtcHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8aafb034c55a8f2, []int{2}
}
func (m *QueryReportedCheckpointBtcHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryReportedCheckpointBtcHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryReportedCheckpointBtcHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryReportedCheckpointBtcHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReportedCheckpointBtcHeightRequest.Merge(m, src)
}
func (m *QueryReportedCheckpointBtcHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryReportedCheckpointBtcHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReportedCheckpointBtcHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReportedCheckpointBtcHeightRequest proto.InternalMessageInfo

func (m *QueryReportedCheckpointBtcHeightRequest) GetCkptHash() string {
	if m != nil {
		return m.CkptHash
	}
	return ""
}

// QueryReportedCheckpointBtcHeightResponse defines a response type for
// ReportedCheckpointBtcHeight RPC method
type QueryReportedCheckpointBtcHeightResponse struct {
	// height of btc light client when checkpoint is reported
	BtcLightClientHeight uint32 `protobuf:"varint,1,opt,name=btc_light_client_height,json=btcLightClientHeight,proto3" json:"btc_light_client_height,omitempty"`
}

func (m *QueryReportedCheckpointBtcHeightResponse) Reset() {
	*m = QueryReportedCheckpointBtcHeightResponse{}
}
func (m *QueryReportedCheckpointBtcHeightResponse) String() string { return proto.CompactTextString(m) }
func (*QueryReportedCheckpointBtcHeightResponse) ProtoMessage()    {}
func (*QueryReportedCheckpointBtcHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8aafb034c55a8f2, []int{3}
}
func (m *QueryReportedCheckpointBtcHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryReportedCheckpointBtcHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryReportedCheckpointBtcHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryReportedCheckpointBtcHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReportedCheckpointBtcHeightResponse.Merge(m, src)
}
func (m *QueryReportedCheckpointBtcHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryReportedCheckpointBtcHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReportedCheckpointBtcHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReportedCheckpointBtcHeightResponse proto.InternalMessageInfo

func (m *QueryReportedCheckpointBtcHeightResponse) GetBtcLightClientHeight() uint32 {
	if m != nil {
		return m.BtcLightClientHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryEndedEpochBtcHeightRequest)(nil), "babylon.monitor.v1.QueryEndedEpochBtcHeightRequest")
	proto.RegisterType((*QueryEndedEpochBtcHeightResponse)(nil), "babylon.monitor.v1.QueryEndedEpochBtcHeightResponse")
	proto.RegisterType((*QueryReportedCheckpointBtcHeightRequest)(nil), "babylon.monitor.v1.QueryReportedCheckpointBtcHeightRequest")
	proto.RegisterType((*QueryReportedCheckpointBtcHeightResponse)(nil), "babylon.monitor.v1.QueryReportedCheckpointBtcHeightResponse")
}

func init() { proto.RegisterFile("babylon/monitor/v1/query.proto", fileDescriptor_a8aafb034c55a8f2) }

var fileDescriptor_a8aafb034c55a8f2 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcf, 0x8b, 0x13, 0x31,
	0x18, 0xed, 0xf8, 0x8b, 0x36, 0xe0, 0x25, 0x0a, 0x4a, 0x2b, 0x63, 0x99, 0x83, 0x16, 0xa4, 0x13,
	0x6a, 0xf5, 0xa4, 0x78, 0x68, 0xa9, 0x14, 0x14, 0xc1, 0xb9, 0xe9, 0x65, 0xc8, 0xa4, 0x61, 0x12,
	0x3a, 0x93, 0xa4, 0x93, 0x4c, 0xb1, 0x94, 0x5e, 0x3c, 0x7a, 0x12, 0xfc, 0x47, 0xfc, 0x33, 0xbc,
	0x08, 0x05, 0x2f, 0x1e, 0x97, 0x76, 0xff, 0x90, 0x65, 0xd2, 0xee, 0x5c, 0x76, 0xba, 0x65, 0x77,
	0x8f, 0xf9, 0x5e, 0xde, 0xfb, 0xbe, 0xf7, 0xbd, 0x04, 0xb8, 0x11, 0x8e, 0x16, 0x89, 0x14, 0x28,
	0x95, 0x82, 0x1b, 0x99, 0xa1, 0x79, 0x0f, 0xcd, 0x72, 0x9a, 0x2d, 0x7c, 0x95, 0x49, 0x23, 0x21,
	0xdc, 0xe3, 0xfe, 0x1e, 0xf7, 0xe7, 0xbd, 0xe6, 0x93, 0x58, 0xca, 0x38, 0xa1, 0x08, 0x2b, 0x8e,
	0xb0, 0x10, 0xd2, 0x60, 0xc3, 0xa5, 0xd0, 0x3b, 0x86, 0xf7, 0x0e, 0x3c, 0xfd, 0x5c, 0x08, 0x8c,
	0xc4, 0x84, 0x4e, 0x46, 0x4a, 0x12, 0x36, 0x30, 0x64, 0x4c, 0x79, 0xcc, 0x4c, 0x40, 0x67, 0x39,
	0xd5, 0x06, 0xb6, 0x40, 0x83, 0x16, 0x40, 0x28, 0xf2, 0xf4, 0xb1, 0xd3, 0x76, 0x3a, 0x77, 0x82,
	0xba, 0x2d, 0x7c, 0xca, 0x53, 0xef, 0x0b, 0x68, 0x1f, 0xe6, 0x6b, 0x25, 0x85, 0xa6, 0xf0, 0x35,
	0x78, 0x14, 0x19, 0x12, 0x26, 0x45, 0x31, 0x24, 0x09, 0xa7, 0xc2, 0x84, 0xcc, 0x5e, 0xb1, 0x72,
	0xf7, 0x83, 0x87, 0x91, 0x21, 0x1f, 0x8b, 0xf3, 0xd0, 0x82, 0x3b, 0xba, 0xf7, 0x1e, 0x3c, 0xb7,
	0xd2, 0x01, 0x55, 0x32, 0x33, 0x74, 0x32, 0x64, 0x94, 0x4c, 0x95, 0xe4, 0xc2, 0x54, 0x8d, 0x48,
	0xa6, 0xca, 0x84, 0x0c, 0x6b, 0x66, 0x35, 0x1b, 0x41, 0xbd, 0x28, 0x8c, 0xb1, 0x66, 0x1e, 0x06,
	0x9d, 0xe3, 0x3a, 0x37, 0x1a, 0xf5, 0xe5, 0x8f, 0xdb, 0xe0, 0xae, 0xed, 0x01, 0x7f, 0x3b, 0xe0,
	0x41, 0xc5, 0x2e, 0x60, 0xdf, 0xbf, 0x18, 0x8d, 0x7f, 0x64, 0xf3, 0xcd, 0x57, 0x57, 0x23, 0xed,
	0x3c, 0x78, 0xfe, 0xf7, 0x7f, 0xa7, 0xbf, 0x6e, 0x75, 0xe0, 0x33, 0x54, 0xf1, 0x5a, 0x6c, 0x70,
	0x1a, 0x2d, 0xcb, 0x44, 0x57, 0xf0, 0xaf, 0x03, 0x5a, 0x97, 0xec, 0x06, 0xbe, 0x39, 0x38, 0xc5,
	0xf1, 0x64, 0x9a, 0x6f, 0xaf, 0x47, 0xde, 0x5b, 0xe9, 0x5b, 0x2b, 0x5d, 0xf8, 0xa2, 0xca, 0x0a,
	0x29, 0x89, 0x1a, 0x2d, 0xcb, 0xf8, 0x57, 0x83, 0x0f, 0x7f, 0x36, 0xae, 0xb3, 0xde, 0xb8, 0xce,
	0xc9, 0xc6, 0x75, 0x7e, 0x6e, 0xdd, 0xda, 0x7a, 0xeb, 0xd6, 0xfe, 0x6f, 0xdd, 0xda, 0xd7, 0x5e,
	0xcc, 0x0d, 0xcb, 0x23, 0x9f, 0xc8, 0xf4, 0x5c, 0x30, 0xc1, 0x91, 0xee, 0x72, 0x59, 0xea, 0x7f,
	0x2b, 0x3b, 0x98, 0x85, 0xa2, 0x3a, 0xba, 0x67, 0xbf, 0x49, 0xff, 0x2c, 0x00, 0x00, 0xff, 0xff,
	0x85, 0xb8, 0x2a, 0x58, 0x7a, 0x03, 0x00, 0x00,
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
	// EndedEpochBtcHeight returns the BTC light client height at provided epoch
	// finish
	EndedEpochBtcHeight(ctx context.Context, in *QueryEndedEpochBtcHeightRequest, opts ...grpc.CallOption) (*QueryEndedEpochBtcHeightResponse, error)
	// ReportedCheckpointBtcHeight returns the BTC light client height at which
	// the checkpoint with the given hash is reported back to Babylon
	ReportedCheckpointBtcHeight(ctx context.Context, in *QueryReportedCheckpointBtcHeightRequest, opts ...grpc.CallOption) (*QueryReportedCheckpointBtcHeightResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) EndedEpochBtcHeight(ctx context.Context, in *QueryEndedEpochBtcHeightRequest, opts ...grpc.CallOption) (*QueryEndedEpochBtcHeightResponse, error) {
	out := new(QueryEndedEpochBtcHeightResponse)
	err := c.cc.Invoke(ctx, "/babylon.monitor.v1.Query/EndedEpochBtcHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ReportedCheckpointBtcHeight(ctx context.Context, in *QueryReportedCheckpointBtcHeightRequest, opts ...grpc.CallOption) (*QueryReportedCheckpointBtcHeightResponse, error) {
	out := new(QueryReportedCheckpointBtcHeightResponse)
	err := c.cc.Invoke(ctx, "/babylon.monitor.v1.Query/ReportedCheckpointBtcHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// EndedEpochBtcHeight returns the BTC light client height at provided epoch
	// finish
	EndedEpochBtcHeight(context.Context, *QueryEndedEpochBtcHeightRequest) (*QueryEndedEpochBtcHeightResponse, error)
	// ReportedCheckpointBtcHeight returns the BTC light client height at which
	// the checkpoint with the given hash is reported back to Babylon
	ReportedCheckpointBtcHeight(context.Context, *QueryReportedCheckpointBtcHeightRequest) (*QueryReportedCheckpointBtcHeightResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) EndedEpochBtcHeight(ctx context.Context, req *QueryEndedEpochBtcHeightRequest) (*QueryEndedEpochBtcHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndedEpochBtcHeight not implemented")
}
func (*UnimplementedQueryServer) ReportedCheckpointBtcHeight(ctx context.Context, req *QueryReportedCheckpointBtcHeightRequest) (*QueryReportedCheckpointBtcHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportedCheckpointBtcHeight not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_EndedEpochBtcHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEndedEpochBtcHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EndedEpochBtcHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.monitor.v1.Query/EndedEpochBtcHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EndedEpochBtcHeight(ctx, req.(*QueryEndedEpochBtcHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ReportedCheckpointBtcHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReportedCheckpointBtcHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ReportedCheckpointBtcHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.monitor.v1.Query/ReportedCheckpointBtcHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ReportedCheckpointBtcHeight(ctx, req.(*QueryReportedCheckpointBtcHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babylon.monitor.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EndedEpochBtcHeight",
			Handler:    _Query_EndedEpochBtcHeight_Handler,
		},
		{
			MethodName: "ReportedCheckpointBtcHeight",
			Handler:    _Query_ReportedCheckpointBtcHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "babylon/monitor/v1/query.proto",
}

func (m *QueryEndedEpochBtcHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEndedEpochBtcHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEndedEpochBtcHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNum != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.EpochNum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryEndedEpochBtcHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEndedEpochBtcHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEndedEpochBtcHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BtcLightClientHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BtcLightClientHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryReportedCheckpointBtcHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryReportedCheckpointBtcHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryReportedCheckpointBtcHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CkptHash) > 0 {
		i -= len(m.CkptHash)
		copy(dAtA[i:], m.CkptHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CkptHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryReportedCheckpointBtcHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryReportedCheckpointBtcHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryReportedCheckpointBtcHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BtcLightClientHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BtcLightClientHeight))
		i--
		dAtA[i] = 0x8
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
func (m *QueryEndedEpochBtcHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNum != 0 {
		n += 1 + sovQuery(uint64(m.EpochNum))
	}
	return n
}

func (m *QueryEndedEpochBtcHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcLightClientHeight != 0 {
		n += 1 + sovQuery(uint64(m.BtcLightClientHeight))
	}
	return n
}

func (m *QueryReportedCheckpointBtcHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CkptHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryReportedCheckpointBtcHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcLightClientHeight != 0 {
		n += 1 + sovQuery(uint64(m.BtcLightClientHeight))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryEndedEpochBtcHeightRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEndedEpochBtcHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEndedEpochBtcHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNum", wireType)
			}
			m.EpochNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryEndedEpochBtcHeightResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEndedEpochBtcHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEndedEpochBtcHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcLightClientHeight", wireType)
			}
			m.BtcLightClientHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BtcLightClientHeight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryReportedCheckpointBtcHeightRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryReportedCheckpointBtcHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryReportedCheckpointBtcHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CkptHash", wireType)
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
			m.CkptHash = string(dAtA[iNdEx:postIndex])
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
func (m *QueryReportedCheckpointBtcHeightResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryReportedCheckpointBtcHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryReportedCheckpointBtcHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcLightClientHeight", wireType)
			}
			m.BtcLightClientHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BtcLightClientHeight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
