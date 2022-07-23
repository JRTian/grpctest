// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: sporting/sporting.proto

package sporting

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *ListSportsRequestFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListEventsRequest) Reset() {
	*x = ListEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sporting_sporting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequest) ProtoMessage() {}

func (x *ListEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sporting_sporting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRequest.ProtoReflect.Descriptor instead.
func (*ListEventsRequest) Descriptor() ([]byte, []int) {
	return file_sporting_sporting_proto_rawDescGZIP(), []int{0}
}

func (x *ListEventsRequest) GetFilter() *ListSportsRequestFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

// Response to ListRaces call.
type ListEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// modified by Gary, keep the original Race
	Sports []*Sport `protobuf:"bytes,1,rep,name=sports,proto3" json:"sports,omitempty"`
}

func (x *ListEventsResponse) Reset() {
	*x = ListEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sporting_sporting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsResponse) ProtoMessage() {}

func (x *ListEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sporting_sporting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsResponse.ProtoReflect.Descriptor instead.
func (*ListEventsResponse) Descriptor() ([]byte, []int) {
	return file_sporting_sporting_proto_rawDescGZIP(), []int{1}
}

func (x *ListEventsResponse) GetSports() []*Sport {
	if x != nil {
		return x.Sports
	}
	return nil
}

// Filter for listing races.
type ListSportsRequestFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *ListSportsRequestFilter) Reset() {
	*x = ListSportsRequestFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sporting_sporting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSportsRequestFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSportsRequestFilter) ProtoMessage() {}

func (x *ListSportsRequestFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sporting_sporting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSportsRequestFilter.ProtoReflect.Descriptor instead.
func (*ListSportsRequestFilter) Descriptor() ([]byte, []int) {
	return file_sporting_sporting_proto_rawDescGZIP(), []int{2}
}

func (x *ListSportsRequestFilter) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

// A race resource.
type Sport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID represents a unique identifier for the race.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name is the official name given to the race.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// AdvertisedStartTime is the time the race is advertised to run.
	AdvertisedStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
}

func (x *Sport) Reset() {
	*x = Sport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sporting_sporting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sport) ProtoMessage() {}

func (x *Sport) ProtoReflect() protoreflect.Message {
	mi := &file_sporting_sporting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sport.ProtoReflect.Descriptor instead.
func (*Sport) Descriptor() ([]byte, []int) {
	return file_sporting_sporting_proto_rawDescGZIP(), []int{3}
}

func (x *Sport) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sport) GetAdvertisedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

type SportWithStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID represents a unique identifier for the race.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name is the official name given to the race.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` // Number represents the number of the race.
	// AdvertisedStartTime is the time the race is advertised to run.
	AdvertisedStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
	// added by Gary as the take 3 requiredß
	Status string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SportWithStatus) Reset() {
	*x = SportWithStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sporting_sporting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportWithStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportWithStatus) ProtoMessage() {}

func (x *SportWithStatus) ProtoReflect() protoreflect.Message {
	mi := &file_sporting_sporting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportWithStatus.ProtoReflect.Descriptor instead.
func (*SportWithStatus) Descriptor() ([]byte, []int) {
	return file_sporting_sporting_proto_rawDescGZIP(), []int{4}
}

func (x *SportWithStatus) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SportWithStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SportWithStatus) GetAdvertisedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

func (x *SportWithStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_sporting_sporting_proto protoreflect.FileDescriptor

var file_sporting_sporting_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x3d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x22, 0x2f, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x22, 0x7b, 0x0a, 0x05, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x4e, 0x0a, 0x15, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x61, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x9d, 0x01, 0x0a, 0x0f, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x61, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0x6f, 0x0a, 0x08, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x63, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a,
	0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sporting_sporting_proto_rawDescOnce sync.Once
	file_sporting_sporting_proto_rawDescData = file_sporting_sporting_proto_rawDesc
)

func file_sporting_sporting_proto_rawDescGZIP() []byte {
	file_sporting_sporting_proto_rawDescOnce.Do(func() {
		file_sporting_sporting_proto_rawDescData = protoimpl.X.CompressGZIP(file_sporting_sporting_proto_rawDescData)
	})
	return file_sporting_sporting_proto_rawDescData
}

var file_sporting_sporting_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sporting_sporting_proto_goTypes = []interface{}{
	(*ListEventsRequest)(nil),       // 0: sporting.ListEventsRequest
	(*ListEventsResponse)(nil),      // 1: sporting.ListEventsResponse
	(*ListSportsRequestFilter)(nil), // 2: sporting.ListSportsRequestFilter
	(*Sport)(nil),                   // 3: sporting.Sport
	(*SportWithStatus)(nil),         // 4: sporting.SportWithStatus
	(*timestamppb.Timestamp)(nil),   // 5: google.protobuf.Timestamp
}
var file_sporting_sporting_proto_depIdxs = []int32{
	2, // 0: sporting.ListEventsRequest.filter:type_name -> sporting.ListSportsRequestFilter
	3, // 1: sporting.ListEventsResponse.sports:type_name -> sporting.Sport
	5, // 2: sporting.Sport.advertised_start_time:type_name -> google.protobuf.Timestamp
	5, // 3: sporting.SportWithStatus.advertised_start_time:type_name -> google.protobuf.Timestamp
	0, // 4: sporting.Sporting.ListEvents:input_type -> sporting.ListEventsRequest
	1, // 5: sporting.Sporting.ListEvents:output_type -> sporting.ListEventsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sporting_sporting_proto_init() }
func file_sporting_sporting_proto_init() {
	if File_sporting_sporting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sporting_sporting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsRequest); i {
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
		file_sporting_sporting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsResponse); i {
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
		file_sporting_sporting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSportsRequestFilter); i {
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
		file_sporting_sporting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sport); i {
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
		file_sporting_sporting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportWithStatus); i {
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
			RawDescriptor: file_sporting_sporting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sporting_sporting_proto_goTypes,
		DependencyIndexes: file_sporting_sporting_proto_depIdxs,
		MessageInfos:      file_sporting_sporting_proto_msgTypes,
	}.Build()
	File_sporting_sporting_proto = out.File
	file_sporting_sporting_proto_rawDesc = nil
	file_sporting_sporting_proto_goTypes = nil
	file_sporting_sporting_proto_depIdxs = nil
}
