// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: proto/order/order.proto

package order

import (
	management "canaanadvisors-test/proto/management"
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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description  string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	RestaurantId int64                  `protobuf:"varint,3,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Items        []*management.Item     `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Order) GetRestaurantId() int64 {
	if x != nil {
		return x.RestaurantId
	}
	return 0
}

func (x *Order) GetItems() []*management.Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type OrderCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId int64              `protobuf:"varint,1,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Description  string             `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Items        []*management.Item `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *OrderCreateRequest) Reset() {
	*x = OrderCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateRequest) ProtoMessage() {}

func (x *OrderCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateRequest.ProtoReflect.Descriptor instead.
func (*OrderCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderCreateRequest) GetRestaurantId() int64 {
	if x != nil {
		return x.RestaurantId
	}
	return 0
}

func (x *OrderCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OrderCreateRequest) GetItems() []*management.Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type OrderCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *Order `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *OrderCreateResponse) Reset() {
	*x = OrderCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateResponse) ProtoMessage() {}

func (x *OrderCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateResponse.ProtoReflect.Descriptor instead.
func (*OrderCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderCreateResponse) GetData() *Order {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OrderCreateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_order_order_proto protoreflect.FileDescriptor

var file_proto_order_order_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x61, 0x6e, 0x61, 0x61,
	0x6e, 0x61, 0x64, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x61,
	0x6e, 0x61, 0x61, 0x6e, 0x61, 0x64, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x61, 0x6e, 0x61, 0x64, 0x76, 0x69, 0x73, 0x6f, 0x72,
	0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x56, 0x0a,
	0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x61, 0x6e, 0x61, 0x64, 0x76, 0x69, 0x73,
	0x6f, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x7a, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x61, 0x6e, 0x61, 0x64, 0x76,
	0x69, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x61,
	0x6e, 0x61, 0x64, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x42, 0x27, 0x5a, 0x25, 0x63, 0x61, 0x6e, 0x61, 0x61, 0x6e, 0x61, 0x64, 0x76, 0x69, 0x73,
	0x6f, 0x72, 0x73, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_order_order_proto_rawDescOnce sync.Once
	file_proto_order_order_proto_rawDescData = file_proto_order_order_proto_rawDesc
)

func file_proto_order_order_proto_rawDescGZIP() []byte {
	file_proto_order_order_proto_rawDescOnce.Do(func() {
		file_proto_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_order_proto_rawDescData)
	})
	return file_proto_order_order_proto_rawDescData
}

var file_proto_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_order_order_proto_goTypes = []interface{}{
	(*Order)(nil),                 // 0: canaanadvisors.Order
	(*OrderCreateRequest)(nil),    // 1: canaanadvisors.OrderCreateRequest
	(*OrderCreateResponse)(nil),   // 2: canaanadvisors.OrderCreateResponse
	(*management.Item)(nil),       // 3: canaanadvisors.Item
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_proto_order_order_proto_depIdxs = []int32{
	3, // 0: canaanadvisors.Order.items:type_name -> canaanadvisors.Item
	4, // 1: canaanadvisors.Order.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: canaanadvisors.OrderCreateRequest.items:type_name -> canaanadvisors.Item
	0, // 3: canaanadvisors.OrderCreateResponse.data:type_name -> canaanadvisors.Order
	1, // 4: canaanadvisors.OrderService.CreateOrder:input_type -> canaanadvisors.OrderCreateRequest
	2, // 5: canaanadvisors.OrderService.CreateOrder:output_type -> canaanadvisors.OrderCreateResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_order_order_proto_init() }
func file_proto_order_order_proto_init() {
	if File_proto_order_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_order_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_proto_order_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateRequest); i {
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
		file_proto_order_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateResponse); i {
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
			RawDescriptor: file_proto_order_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_order_proto_goTypes,
		DependencyIndexes: file_proto_order_order_proto_depIdxs,
		MessageInfos:      file_proto_order_order_proto_msgTypes,
	}.Build()
	File_proto_order_order_proto = out.File
	file_proto_order_order_proto_rawDesc = nil
	file_proto_order_order_proto_goTypes = nil
	file_proto_order_order_proto_depIdxs = nil
}
