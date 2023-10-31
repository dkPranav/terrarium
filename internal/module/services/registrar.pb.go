// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: pb/terrarium/module/services/registrar.proto

package services

import (
	module "github.com/terrariumcloud/terrarium/pkg/terrarium/module"
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

type ModuleMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string          `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Name         string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Provider     string          `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	Description  string          `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	SourceUrl    string          `protobuf:"bytes,5,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	Maturity     module.Maturity `protobuf:"varint,6,opt,name=maturity,proto3,enum=terrarium.module.Maturity" json:"maturity,omitempty"`
}

func (x *ModuleMetadata) Reset() {
	*x = ModuleMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_terrarium_module_services_registrar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleMetadata) ProtoMessage() {}

func (x *ModuleMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_pb_terrarium_module_services_registrar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleMetadata.ProtoReflect.Descriptor instead.
func (*ModuleMetadata) Descriptor() ([]byte, []int) {
	return file_pb_terrarium_module_services_registrar_proto_rawDescGZIP(), []int{0}
}

func (x *ModuleMetadata) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *ModuleMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModuleMetadata) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ModuleMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ModuleMetadata) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *ModuleMetadata) GetMaturity() module.Maturity {
	if x != nil {
		return x.Maturity
	}
	return module.Maturity(0)
}

type ListModulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListModulesRequest) Reset() {
	*x = ListModulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_terrarium_module_services_registrar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModulesRequest) ProtoMessage() {}

func (x *ListModulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_terrarium_module_services_registrar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModulesRequest.ProtoReflect.Descriptor instead.
func (*ListModulesRequest) Descriptor() ([]byte, []int) {
	return file_pb_terrarium_module_services_registrar_proto_rawDescGZIP(), []int{1}
}

type ListModulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*ModuleMetadata `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *ListModulesResponse) Reset() {
	*x = ListModulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_terrarium_module_services_registrar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModulesResponse) ProtoMessage() {}

func (x *ListModulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_terrarium_module_services_registrar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModulesResponse.ProtoReflect.Descriptor instead.
func (*ListModulesResponse) Descriptor() ([]byte, []int) {
	return file_pb_terrarium_module_services_registrar_proto_rawDescGZIP(), []int{2}
}

func (x *ListModulesResponse) GetModules() []*ModuleMetadata {
	if x != nil {
		return x.Modules
	}
	return nil
}

type GetModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetModuleRequest) Reset() {
	*x = GetModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_terrarium_module_services_registrar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModuleRequest) ProtoMessage() {}

func (x *GetModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_terrarium_module_services_registrar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModuleRequest.ProtoReflect.Descriptor instead.
func (*GetModuleRequest) Descriptor() ([]byte, []int) {
	return file_pb_terrarium_module_services_registrar_proto_rawDescGZIP(), []int{3}
}

func (x *GetModuleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetModuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module *ModuleMetadata `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
}

func (x *GetModuleResponse) Reset() {
	*x = GetModuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_terrarium_module_services_registrar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModuleResponse) ProtoMessage() {}

func (x *GetModuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_terrarium_module_services_registrar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModuleResponse.ProtoReflect.Descriptor instead.
func (*GetModuleResponse) Descriptor() ([]byte, []int) {
	return file_pb_terrarium_module_services_registrar_proto_rawDescGZIP(), []int{4}
}

func (x *GetModuleResponse) GetModule() *ModuleMetadata {
	if x != nil {
		return x.Module
	}
	return nil
}

var File_pb_terrarium_module_services_registrar_proto protoreflect.FileDescriptor

var file_pb_terrarium_module_services_registrar_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x20, 0x70, 0x62, 0x2f, 0x74, 0x65,
	0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x0e,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22,
	0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75,
	0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x5a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x72, 0x72,
	0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x26, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x32, 0xb2, 0x02,
	0x0a, 0x09, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x12, 0x4f, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72,
	0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x74, 0x65,
	0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72,
	0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d,
	0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_terrarium_module_services_registrar_proto_rawDescOnce sync.Once
	file_pb_terrarium_module_services_registrar_proto_rawDescData = file_pb_terrarium_module_services_registrar_proto_rawDesc
)

func file_pb_terrarium_module_services_registrar_proto_rawDescGZIP() []byte {
	file_pb_terrarium_module_services_registrar_proto_rawDescOnce.Do(func() {
		file_pb_terrarium_module_services_registrar_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_terrarium_module_services_registrar_proto_rawDescData)
	})
	return file_pb_terrarium_module_services_registrar_proto_rawDescData
}

var file_pb_terrarium_module_services_registrar_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_terrarium_module_services_registrar_proto_goTypes = []interface{}{
	(*ModuleMetadata)(nil),               // 0: terrarium.module.services.ModuleMetadata
	(*ListModulesRequest)(nil),           // 1: terrarium.module.services.ListModulesRequest
	(*ListModulesResponse)(nil),          // 2: terrarium.module.services.ListModulesResponse
	(*GetModuleRequest)(nil),             // 3: terrarium.module.services.GetModuleRequest
	(*GetModuleResponse)(nil),            // 4: terrarium.module.services.GetModuleResponse
	(module.Maturity)(0),                 // 5: terrarium.module.Maturity
	(*module.RegisterModuleRequest)(nil), // 6: terrarium.module.RegisterModuleRequest
	(*module.Response)(nil),              // 7: terrarium.module.Response
}
var file_pb_terrarium_module_services_registrar_proto_depIdxs = []int32{
	5, // 0: terrarium.module.services.ModuleMetadata.maturity:type_name -> terrarium.module.Maturity
	0, // 1: terrarium.module.services.ListModulesResponse.modules:type_name -> terrarium.module.services.ModuleMetadata
	0, // 2: terrarium.module.services.GetModuleResponse.module:type_name -> terrarium.module.services.ModuleMetadata
	6, // 3: terrarium.module.services.Registrar.Register:input_type -> terrarium.module.RegisterModuleRequest
	1, // 4: terrarium.module.services.Registrar.ListModules:input_type -> terrarium.module.services.ListModulesRequest
	3, // 5: terrarium.module.services.Registrar.GetModule:input_type -> terrarium.module.services.GetModuleRequest
	7, // 6: terrarium.module.services.Registrar.Register:output_type -> terrarium.module.Response
	2, // 7: terrarium.module.services.Registrar.ListModules:output_type -> terrarium.module.services.ListModulesResponse
	4, // 8: terrarium.module.services.Registrar.GetModule:output_type -> terrarium.module.services.GetModuleResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_terrarium_module_services_registrar_proto_init() }
func file_pb_terrarium_module_services_registrar_proto_init() {
	if File_pb_terrarium_module_services_registrar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_terrarium_module_services_registrar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleMetadata); i {
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
		file_pb_terrarium_module_services_registrar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModulesRequest); i {
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
		file_pb_terrarium_module_services_registrar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModulesResponse); i {
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
		file_pb_terrarium_module_services_registrar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModuleRequest); i {
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
		file_pb_terrarium_module_services_registrar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModuleResponse); i {
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
			RawDescriptor: file_pb_terrarium_module_services_registrar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_terrarium_module_services_registrar_proto_goTypes,
		DependencyIndexes: file_pb_terrarium_module_services_registrar_proto_depIdxs,
		MessageInfos:      file_pb_terrarium_module_services_registrar_proto_msgTypes,
	}.Build()
	File_pb_terrarium_module_services_registrar_proto = out.File
	file_pb_terrarium_module_services_registrar_proto_rawDesc = nil
	file_pb_terrarium_module_services_registrar_proto_goTypes = nil
	file_pb_terrarium_module_services_registrar_proto_depIdxs = nil
}
