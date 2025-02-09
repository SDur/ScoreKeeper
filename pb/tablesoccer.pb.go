// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tablesoccer.proto

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

type MatchResult struct {
	Teams                []*MatchResult_Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MatchResult) Reset()         { *m = MatchResult{} }
func (m *MatchResult) String() string { return proto.CompactTextString(m) }
func (*MatchResult) ProtoMessage()    {}
func (*MatchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_479a75e1ae1fe396, []int{0}
}

func (m *MatchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchResult.Unmarshal(m, b)
}
func (m *MatchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchResult.Marshal(b, m, deterministic)
}
func (m *MatchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchResult.Merge(m, src)
}
func (m *MatchResult) XXX_Size() int {
	return xxx_messageInfo_MatchResult.Size(m)
}
func (m *MatchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchResult.DiscardUnknown(m)
}

var xxx_messageInfo_MatchResult proto.InternalMessageInfo

func (m *MatchResult) GetTeams() []*MatchResult_Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type MatchResult_Team struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	Score                int32     `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MatchResult_Team) Reset()         { *m = MatchResult_Team{} }
func (m *MatchResult_Team) String() string { return proto.CompactTextString(m) }
func (*MatchResult_Team) ProtoMessage()    {}
func (*MatchResult_Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_479a75e1ae1fe396, []int{0, 0}
}

func (m *MatchResult_Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchResult_Team.Unmarshal(m, b)
}
func (m *MatchResult_Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchResult_Team.Marshal(b, m, deterministic)
}
func (m *MatchResult_Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchResult_Team.Merge(m, src)
}
func (m *MatchResult_Team) XXX_Size() int {
	return xxx_messageInfo_MatchResult_Team.Size(m)
}
func (m *MatchResult_Team) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchResult_Team.DiscardUnknown(m)
}

var xxx_messageInfo_MatchResult_Team proto.InternalMessageInfo

func (m *MatchResult_Team) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

func (m *MatchResult_Team) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type Person struct {
	Firstname            string   `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Wins                 int32    `protobuf:"varint,3,opt,name=wins,proto3" json:"wins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_479a75e1ae1fe396, []int{1}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Person) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *Person) GetWins() int32 {
	if m != nil {
		return m.Wins
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_479a75e1ae1fe396, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MatchResult)(nil), "MatchResult")
	proto.RegisterType((*MatchResult_Team)(nil), "MatchResult.Team")
	proto.RegisterType((*Person)(nil), "Person")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("tablesoccer.proto", fileDescriptor_479a75e1ae1fe396) }

var fileDescriptor_479a75e1ae1fe396 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0x87, 0x49, 0x9a, 0x3f, 0xed, 0xf4, 0xbd, 0x74, 0x78, 0x0f, 0x21, 0xf4, 0x10, 0x73, 0x31,
	0xa7, 0x3d, 0xd4, 0x0f, 0x20, 0x08, 0xe2, 0x41, 0x04, 0xd9, 0x48, 0x0f, 0xde, 0x36, 0xcb, 0x88,
	0x81, 0x24, 0xbb, 0xec, 0x8e, 0x16, 0xbf, 0xbd, 0x74, 0x63, 0xb4, 0xb7, 0x79, 0xe6, 0x99, 0xd9,
	0xe5, 0x37, 0xb0, 0x63, 0xd5, 0x0d, 0xe4, 0x8d, 0xd6, 0xe4, 0x84, 0x75, 0x86, 0x4d, 0x7d, 0x82,
	0xed, 0x93, 0x62, 0xfd, 0x2e, 0xc9, 0x7f, 0x0c, 0x8c, 0xd7, 0x90, 0x32, 0xa9, 0xd1, 0x17, 0x51,
	0xb5, 0x6a, 0xb6, 0x87, 0x9d, 0xb8, 0x90, 0xe2, 0x85, 0xd4, 0x28, 0x67, 0x5f, 0xde, 0x42, 0x72,
	0x46, 0xbc, 0x82, 0xdc, 0x92, 0xf3, 0x66, 0x5a, 0x56, 0x72, 0xf1, 0x1c, 0x58, 0x2e, 0x7d, 0xfc,
	0x0f, 0xa9, 0xd7, 0xc6, 0x51, 0x11, 0x57, 0x51, 0x93, 0xca, 0x19, 0xea, 0x23, 0x64, 0xf3, 0x20,
	0xee, 0x61, 0xf3, 0xd6, 0x3b, 0xcf, 0x93, 0x1a, 0xa9, 0x88, 0xaa, 0xa8, 0xd9, 0xc8, 0xbf, 0x06,
	0x96, 0xb0, 0x1e, 0xd4, 0x8f, 0x8c, 0x83, 0xfc, 0x65, 0x44, 0x48, 0x4e, 0xfd, 0xe4, 0x8b, 0x55,
	0x78, 0x38, 0xd4, 0x75, 0x0e, 0xe9, 0xfd, 0x68, 0xf9, 0xeb, 0x70, 0x04, 0x6c, 0xcf, 0x3f, 0x3d,
	0x12, 0x59, 0x72, 0x2d, 0xb9, 0xcf, 0x5e, 0x13, 0xd6, 0x00, 0x2d, 0x1b, 0x47, 0x41, 0xe1, 0xbf,
	0xcb, 0x7c, 0x65, 0x26, 0xc2, 0x26, 0xee, 0x61, 0xfd, 0x40, 0x3c, 0x4f, 0x2c, 0x71, 0xca, 0xa5,
	0xb8, 0x4b, 0x5e, 0x63, 0xdb, 0x75, 0x59, 0x38, 0xdf, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2b, 0xc9, 0xfb, 0x3e, 0x53, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScoreKeeperServiceClient is the client API for ScoreKeeperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScoreKeeperServiceClient interface {
	StoreScore(ctx context.Context, in *MatchResult, opts ...grpc.CallOption) (*Empty, error)
	GetScore(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
}

type scoreKeeperServiceClient struct {
	cc *grpc.ClientConn
}

func NewScoreKeeperServiceClient(cc *grpc.ClientConn) ScoreKeeperServiceClient {
	return &scoreKeeperServiceClient{cc}
}

func (c *scoreKeeperServiceClient) StoreScore(ctx context.Context, in *MatchResult, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ScoreKeeperService/StoreScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreKeeperServiceClient) GetScore(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/ScoreKeeperService/GetScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoreKeeperServiceServer is the server API for ScoreKeeperService service.
type ScoreKeeperServiceServer interface {
	StoreScore(context.Context, *MatchResult) (*Empty, error)
	GetScore(context.Context, *Person) (*Person, error)
}

// UnimplementedScoreKeeperServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScoreKeeperServiceServer struct {
}

func (*UnimplementedScoreKeeperServiceServer) StoreScore(ctx context.Context, req *MatchResult) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreScore not implemented")
}
func (*UnimplementedScoreKeeperServiceServer) GetScore(ctx context.Context, req *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScore not implemented")
}

func RegisterScoreKeeperServiceServer(s *grpc.Server, srv ScoreKeeperServiceServer) {
	s.RegisterService(&_ScoreKeeperService_serviceDesc, srv)
}

func _ScoreKeeperService_StoreScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreKeeperServiceServer).StoreScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ScoreKeeperService/StoreScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreKeeperServiceServer).StoreScore(ctx, req.(*MatchResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreKeeperService_GetScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreKeeperServiceServer).GetScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ScoreKeeperService/GetScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreKeeperServiceServer).GetScore(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScoreKeeperService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ScoreKeeperService",
	HandlerType: (*ScoreKeeperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreScore",
			Handler:    _ScoreKeeperService_StoreScore_Handler,
		},
		{
			MethodName: "GetScore",
			Handler:    _ScoreKeeperService_GetScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tablesoccer.proto",
}
