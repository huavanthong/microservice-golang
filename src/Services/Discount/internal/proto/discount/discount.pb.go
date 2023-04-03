// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: discount.proto

package discount

import (
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

// DiscountRequest defines the request for a Discount call
type DiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName string `protobuf:"bytes,1,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
}

func (x *DiscountRequest) Reset() {
	*x = DiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountRequest) ProtoMessage() {}

func (x *DiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountRequest.ProtoReflect.Descriptor instead.
func (*DiscountRequest) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{0}
}

func (x *DiscountRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

// DiscountResponse is the response from a DiscountRequest call
type DiscountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName string  `protobuf:"bytes,1,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Amount      float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *DiscountResponse) Reset() {
	*x = DiscountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountResponse) ProtoMessage() {}

func (x *DiscountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountResponse.ProtoReflect.Descriptor instead.
func (*DiscountResponse) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{1}
}

func (x *DiscountResponse) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *DiscountResponse) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DiscountResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_discount_proto protoreflect.FileDescriptor

var file_discount_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x34, 0x0a, 0x0f, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x6f, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0xb4, 0x02, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x1a, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x1a, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x19, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discount_proto_rawDescOnce sync.Once
	file_discount_proto_rawDescData = file_discount_proto_rawDesc
)

func file_discount_proto_rawDescGZIP() []byte {
	file_discount_proto_rawDescOnce.Do(func() {
		file_discount_proto_rawDescData = protoimpl.X.CompressGZIP(file_discount_proto_rawDescData)
	})
	return file_discount_proto_rawDescData
}

var file_discount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_discount_proto_goTypes = []interface{}{
	(*DiscountRequest)(nil),  // 0: discount.DiscountRequest
	(*DiscountResponse)(nil), // 1: discount.DiscountResponse
}
var file_discount_proto_depIdxs = []int32{
	0, // 0: discount.DiscountService.GetDiscount:input_type -> discount.DiscountRequest
	1, // 1: discount.DiscountService.CreateDiscount:input_type -> discount.DiscountResponse
	1, // 2: discount.DiscountService.UpdateDiscount:input_type -> discount.DiscountResponse
	0, // 3: discount.DiscountService.DeleteDiscount:input_type -> discount.DiscountRequest
	1, // 4: discount.DiscountService.GetDiscount:output_type -> discount.DiscountResponse
	1, // 5: discount.DiscountService.CreateDiscount:output_type -> discount.DiscountResponse
	1, // 6: discount.DiscountService.UpdateDiscount:output_type -> discount.DiscountResponse
	1, // 7: discount.DiscountService.DeleteDiscount:output_type -> discount.DiscountResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_discount_proto_init() }
func file_discount_proto_init() {
	if File_discount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountRequest); i {
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
		file_discount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountResponse); i {
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
			RawDescriptor: file_discount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discount_proto_goTypes,
		DependencyIndexes: file_discount_proto_depIdxs,
		MessageInfos:      file_discount_proto_msgTypes,
	}.Build()
	File_discount_proto = out.File
	file_discount_proto_rawDesc = nil
	file_discount_proto_goTypes = nil
	file_discount_proto_depIdxs = nil
}
