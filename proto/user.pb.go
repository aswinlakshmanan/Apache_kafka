// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: grpc/proto/user.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Details) Reset() {
	*x = Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *Details) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Details) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Detailsrequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Information *Details `protobuf:"bytes,1,opt,name=information,proto3" json:"information,omitempty"`
}

func (x *Detailsrequest) Reset() {
	*x = Detailsrequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detailsrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detailsrequest) ProtoMessage() {}

func (x *Detailsrequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detailsrequest.ProtoReflect.Descriptor instead.
func (*Detailsrequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *Detailsrequest) GetInformation() *Details {
	if x != nil {
		return x.Information
	}
	return nil
}

type Detailsresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InformationResponse *Details `protobuf:"bytes,1,opt,name=information_response,json=informationResponse,proto3" json:"information_response,omitempty"`
}

func (x *Detailsresponse) Reset() {
	*x = Detailsresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detailsresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detailsresponse) ProtoMessage() {}

func (x *Detailsresponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detailsresponse.ProtoReflect.Descriptor instead.
func (*Detailsresponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *Detailsresponse) GetInformationResponse() *Details {
	if x != nil {
		return x.InformationResponse
	}
	return nil
}

var File_grpc_proto_user_proto protoreflect.FileDescriptor

var file_grpc_proto_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x39, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x0e, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0b, 0x69, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b,
	0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x0f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x14, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x13, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x47, 0x0a, 0x0e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_user_proto_rawDescOnce sync.Once
	file_grpc_proto_user_proto_rawDescData = file_grpc_proto_user_proto_rawDesc
)

func file_grpc_proto_user_proto_rawDescGZIP() []byte {
	file_grpc_proto_user_proto_rawDescOnce.Do(func() {
		file_grpc_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_user_proto_rawDescData)
	})
	return file_grpc_proto_user_proto_rawDescData
}

var file_grpc_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_proto_user_proto_goTypes = []interface{}{
	(*Details)(nil),         // 0: grpc.details
	(*Detailsrequest)(nil),  // 1: grpc.detailsrequest
	(*Detailsresponse)(nil), // 2: grpc.detailsresponse
}
var file_grpc_proto_user_proto_depIdxs = []int32{
	0, // 0: grpc.detailsrequest.information:type_name -> grpc.details
	0, // 1: grpc.detailsresponse.information_response:type_name -> grpc.details
	1, // 2: grpc.detailsservice.info:input_type -> grpc.detailsrequest
	2, // 3: grpc.detailsservice.info:output_type -> grpc.detailsresponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_proto_user_proto_init() }
func file_grpc_proto_user_proto_init() {
	if File_grpc_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Details); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Detailsrequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Detailsresponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_user_proto_goTypes,
		DependencyIndexes: file_grpc_proto_user_proto_depIdxs,
		MessageInfos:      file_grpc_proto_user_proto_msgTypes,
	}.Build()
	File_grpc_proto_user_proto = out.File
	file_grpc_proto_user_proto_rawDesc = nil
	file_grpc_proto_user_proto_goTypes = nil
	file_grpc_proto_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DetailsserviceClient is the client API for Detailsservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DetailsserviceClient interface {
	Info(ctx context.Context, in *Detailsrequest, opts ...grpc.CallOption) (*Detailsresponse, error)
}

type detailsserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewDetailsserviceClient(cc grpc.ClientConnInterface) DetailsserviceClient {
	return &detailsserviceClient{cc}
}

func (c *detailsserviceClient) Info(ctx context.Context, in *Detailsrequest, opts ...grpc.CallOption) (*Detailsresponse, error) {
	out := new(Detailsresponse)
	err := c.cc.Invoke(ctx, "/grpc.detailsservice/info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailsserviceServer is the server API for Detailsservice service.
type DetailsserviceServer interface {
	Info(context.Context, *Detailsrequest) (*Detailsresponse, error)
}

// UnimplementedDetailsserviceServer can be embedded to have forward compatible implementations.
type UnimplementedDetailsserviceServer struct {
}

func (*UnimplementedDetailsserviceServer) Info(context.Context, *Detailsrequest) (*Detailsresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

func RegisterDetailsserviceServer(s *grpc.Server, srv DetailsserviceServer) {
	s.RegisterService(&_Detailsservice_serviceDesc, srv)
}

func _Detailsservice_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Detailsrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailsserviceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.detailsservice/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailsserviceServer).Info(ctx, req.(*Detailsrequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Detailsservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.detailsservice",
	HandlerType: (*DetailsserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "info",
			Handler:    _Detailsservice_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/user.proto",
}
