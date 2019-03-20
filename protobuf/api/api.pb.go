// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	ptype "google.golang.org/genproto/protobuf/ptype"
	source_context "google.golang.org/genproto/protobuf/source_context"
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

// Api is a light-weight descriptor for an API Interface.
//
// Interfaces are also described as "protocol buffer services" in some contexts,
// such as by the "service" keyword in a .proto file, but they are different
// from API Services, which represent a concrete implementation of an interface
// as opposed to simply a description of methods and bindings. They are also
// sometimes simply referred to as "APIs" in other contexts, such as the name of
// this message itself. See https://cloud.google.com/apis/design/glossary for
// detailed terminology.
type Api struct {
	// The fully qualified name of this interface, including package name
	// followed by the interface's simple name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The methods of this interface, in unspecified order.
	Methods []*Method `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
	// Any metadata attached to the interface.
	Options []*ptype.Option `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	// A version string for this interface. If specified, must have the form
	// `major-version.minor-version`, as in `1.10`. If the minor version is
	// omitted, it defaults to zero. If the entire version field is empty, the
	// major version is derived from the package name, as outlined below. If the
	// field is not empty, the version in the package name will be verified to be
	// consistent with what is provided here.
	//
	// The versioning schema uses [semantic
	// versioning](http://semver.org) where the major version number
	// indicates a breaking change and the minor version an additive,
	// non-breaking change. Both version numbers are signals to users
	// what to expect from different versions, and should be carefully
	// chosen based on the product plan.
	//
	// The major version is also reflected in the package name of the
	// interface, which must end in `v<major-version>`, as in
	// `google.feature.v1`. For major versions 0 and 1, the suffix can
	// be omitted. Zero major versions must only be used for
	// experimental, non-GA interfaces.
	//
	//
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Source context for the protocol buffer service represented by this
	// message.
	SourceContext *source_context.SourceContext `protobuf:"bytes,5,opt,name=source_context,json=sourceContext,proto3" json:"source_context,omitempty"`
	// Included interfaces. See [Mixin][].
	Mixins []*Mixin `protobuf:"bytes,6,rep,name=mixins,proto3" json:"mixins,omitempty"`
	// The source syntax of the service.
	Syntax               ptype.Syntax `protobuf:"varint,7,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Api) Reset()         { *m = Api{} }
func (m *Api) String() string { return proto.CompactTextString(m) }
func (*Api) ProtoMessage()    {}
func (*Api) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2ec32096296c143, []int{0}
}

func (m *Api) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Api.Unmarshal(m, b)
}
func (m *Api) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Api.Marshal(b, m, deterministic)
}
func (m *Api) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Api.Merge(m, src)
}
func (m *Api) XXX_Size() int {
	return xxx_messageInfo_Api.Size(m)
}
func (m *Api) XXX_DiscardUnknown() {
	xxx_messageInfo_Api.DiscardUnknown(m)
}

var xxx_messageInfo_Api proto.InternalMessageInfo

func (m *Api) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Api) GetMethods() []*Method {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *Api) GetOptions() []*ptype.Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Api) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Api) GetSourceContext() *source_context.SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

func (m *Api) GetMixins() []*Mixin {
	if m != nil {
		return m.Mixins
	}
	return nil
}

func (m *Api) GetSyntax() ptype.Syntax {
	if m != nil {
		return m.Syntax
	}
	return ptype.Syntax_SYNTAX_PROTO2
}

// Method represents a method of an API interface.
type Method struct {
	// The simple name of this method.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A URL of the input message type.
	RequestTypeUrl string `protobuf:"bytes,2,opt,name=request_type_url,json=requestTypeUrl,proto3" json:"request_type_url,omitempty"`
	// If true, the request is streamed.
	RequestStreaming bool `protobuf:"varint,3,opt,name=request_streaming,json=requestStreaming,proto3" json:"request_streaming,omitempty"`
	// The URL of the output message type.
	ResponseTypeUrl string `protobuf:"bytes,4,opt,name=response_type_url,json=responseTypeUrl,proto3" json:"response_type_url,omitempty"`
	// If true, the response is streamed.
	ResponseStreaming bool `protobuf:"varint,5,opt,name=response_streaming,json=responseStreaming,proto3" json:"response_streaming,omitempty"`
	// Any metadata attached to the method.
	Options []*ptype.Option `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty"`
	// The source syntax of this method.
	Syntax               ptype.Syntax `protobuf:"varint,7,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Method) Reset()         { *m = Method{} }
func (m *Method) String() string { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()    {}
func (*Method) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2ec32096296c143, []int{1}
}

func (m *Method) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Method.Unmarshal(m, b)
}
func (m *Method) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Method.Marshal(b, m, deterministic)
}
func (m *Method) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method.Merge(m, src)
}
func (m *Method) XXX_Size() int {
	return xxx_messageInfo_Method.Size(m)
}
func (m *Method) XXX_DiscardUnknown() {
	xxx_messageInfo_Method.DiscardUnknown(m)
}

var xxx_messageInfo_Method proto.InternalMessageInfo

func (m *Method) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Method) GetRequestTypeUrl() string {
	if m != nil {
		return m.RequestTypeUrl
	}
	return ""
}

func (m *Method) GetRequestStreaming() bool {
	if m != nil {
		return m.RequestStreaming
	}
	return false
}

func (m *Method) GetResponseTypeUrl() string {
	if m != nil {
		return m.ResponseTypeUrl
	}
	return ""
}

func (m *Method) GetResponseStreaming() bool {
	if m != nil {
		return m.ResponseStreaming
	}
	return false
}

func (m *Method) GetOptions() []*ptype.Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Method) GetSyntax() ptype.Syntax {
	if m != nil {
		return m.Syntax
	}
	return ptype.Syntax_SYNTAX_PROTO2
}

// Declares an API Interface to be included in this interface. The including
// interface must redeclare all the methods from the included interface, but
// documentation and options are inherited as follows:
//
// - If after comment and whitespace stripping, the documentation
//   string of the redeclared method is empty, it will be inherited
//   from the original method.
//
// - Each annotation belonging to the service config (http,
//   visibility) which is not set in the redeclared method will be
//   inherited.
//
// - If an http annotation is inherited, the path pattern will be
//   modified as follows. Any version prefix will be replaced by the
//   version of the including interface plus the [root][] path if
//   specified.
//
// Example of a simple mixin:
//
//     package google.acl.v1;
//     service AccessControl {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v1/{resource=**}:getAcl";
//       }
//     }
//
//     package google.storage.v2;
//     service Storage {
//       rpc GetAcl(GetAclRequest) returns (Acl);
//
//       // Get a data record.
//       rpc GetData(GetDataRequest) returns (Data) {
//         option (google.api.http).get = "/v2/{resource=**}";
//       }
//     }
//
// Example of a mixin configuration:
//
//     apis:
//     - name: google.storage.v2.Storage
//       mixins:
//       - name: google.acl.v1.AccessControl
//
// The mixin construct implies that all methods in `AccessControl` are
// also declared with same name and request/response types in
// `Storage`. A documentation generator or annotation processor will
// see the effective `Storage.GetAcl` method after inherting
// documentation and annotations as follows:
//
//     service Storage {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v2/{resource=**}:getAcl";
//       }
//       ...
//     }
//
// Note how the version in the path pattern changed from `v1` to `v2`.
//
// If the `root` field in the mixin is specified, it should be a
// relative path under which inherited HTTP paths are placed. Example:
//
//     apis:
//     - name: google.storage.v2.Storage
//       mixins:
//       - name: google.acl.v1.AccessControl
//         root: acls
//
// This implies the following inherited HTTP annotation:
//
//     service Storage {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v2/acls/{resource=**}:getAcl";
//       }
//       ...
//     }
type Mixin struct {
	// The fully qualified name of the interface which is included.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If non-empty specifies a path under which inherited HTTP paths
	// are rooted.
	Root                 string   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mixin) Reset()         { *m = Mixin{} }
func (m *Mixin) String() string { return proto.CompactTextString(m) }
func (*Mixin) ProtoMessage()    {}
func (*Mixin) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2ec32096296c143, []int{2}
}

func (m *Mixin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mixin.Unmarshal(m, b)
}
func (m *Mixin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mixin.Marshal(b, m, deterministic)
}
func (m *Mixin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mixin.Merge(m, src)
}
func (m *Mixin) XXX_Size() int {
	return xxx_messageInfo_Mixin.Size(m)
}
func (m *Mixin) XXX_DiscardUnknown() {
	xxx_messageInfo_Mixin.DiscardUnknown(m)
}

var xxx_messageInfo_Mixin proto.InternalMessageInfo

func (m *Mixin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Mixin) GetRoot() string {
	if m != nil {
		return m.Root
	}
	return ""
}

func init() {
	proto.RegisterType((*Api)(nil), "google.protobuf.Api")
	proto.RegisterType((*Method)(nil), "google.protobuf.Method")
	proto.RegisterType((*Mixin)(nil), "google.protobuf.Mixin")
}

func init() { proto.RegisterFile("google/protobuf/api.proto", fileDescriptor_a2ec32096296c143) }

var fileDescriptor_a2ec32096296c143 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x95, 0xa4, 0x4d, 0x17, 0xaf, 0xe8, 0x82, 0x91, 0xc0, 0xf4, 0xb0, 0x8a, 0x56, 0x1c,
	0x22, 0x2a, 0x12, 0x51, 0x8e, 0x9c, 0x5a, 0x84, 0x7a, 0x40, 0x88, 0x28, 0x05, 0x21, 0x71, 0xa9,
	0xd2, 0x62, 0x82, 0xa5, 0xc4, 0x63, 0x6c, 0x07, 0xda, 0xd7, 0xe1, 0xc8, 0x91, 0x37, 0xe0, 0xcd,
	0x50, 0x9c, 0xb8, 0x7f, 0xd2, 0x22, 0xb1, 0x37, 0x8f, 0xbf, 0xdf, 0x7c, 0x99, 0xf9, 0xac, 0xa0,
	0xc7, 0x39, 0x40, 0x5e, 0xd0, 0x58, 0x48, 0xd0, 0xb0, 0xaa, 0xbe, 0xc4, 0x99, 0x60, 0x91, 0x29,
	0xf0, 0x55, 0x23, 0x45, 0x56, 0x1a, 0x3d, 0xe9, 0xb2, 0x0a, 0x2a, 0xb9, 0xa6, 0xcb, 0x35, 0x70,
	0x4d, 0x37, 0xba, 0x01, 0x47, 0xa3, 0x2e, 0xa5, 0xb7, 0xa2, 0x35, 0xb9, 0xf9, 0xe3, 0x22, 0x6f,
	0x2a, 0x18, 0xc6, 0xa8, 0xc7, 0xb3, 0x92, 0x12, 0x27, 0x70, 0xc2, 0x3b, 0xa9, 0x39, 0xe3, 0xe7,
	0x68, 0x50, 0x52, 0xfd, 0x15, 0x3e, 0x2b, 0xe2, 0x06, 0x5e, 0x78, 0x39, 0x79, 0x14, 0x75, 0x06,
	0x88, 0xde, 0x1a, 0x3d, 0xb5, 0x5c, 0xdd, 0x02, 0x42, 0x33, 0xe0, 0x8a, 0x78, 0xff, 0x68, 0x79,
	0x67, 0xf4, 0xd4, 0x72, 0x98, 0xa0, 0xc1, 0x77, 0x2a, 0x15, 0x03, 0x4e, 0x7a, 0xe6, 0xe3, 0xb6,
	0xc4, 0xaf, 0xd1, 0xf0, 0x78, 0x1f, 0xd2, 0x0f, 0x9c, 0xf0, 0x72, 0x72, 0x7d, 0xe2, 0xb9, 0x30,
	0xd8, 0xab, 0x86, 0x4a, 0xef, 0xaa, 0xc3, 0x12, 0x47, 0xc8, 0x2f, 0xd9, 0x86, 0x71, 0x45, 0x7c,
	0x33, 0xd2, 0xc3, 0xd3, 0x2d, 0x6a, 0x39, 0x6d, 0x29, 0x1c, 0x23, 0x5f, 0x6d, 0xb9, 0xce, 0x36,
	0x64, 0x10, 0x38, 0xe1, 0xf0, 0xcc, 0x0a, 0x0b, 0x23, 0xa7, 0x2d, 0x76, 0xf3, 0xdb, 0x45, 0x7e,
	0x13, 0xc4, 0xd9, 0x18, 0x43, 0x74, 0x4f, 0xd2, 0x6f, 0x15, 0x55, 0x7a, 0x59, 0x07, 0xbf, 0xac,
	0x64, 0x41, 0x5c, 0xa3, 0x0f, 0xdb, 0xfb, 0xf7, 0x5b, 0x41, 0x3f, 0xc8, 0x02, 0x8f, 0xd1, 0x7d,
	0x4b, 0x2a, 0x2d, 0x69, 0x56, 0x32, 0x9e, 0x13, 0x2f, 0x70, 0xc2, 0x8b, 0xd4, 0x5a, 0x2c, 0xec,
	0x3d, 0x7e, 0x5a, 0xc3, 0x4a, 0x00, 0x57, 0x74, 0xef, 0xdb, 0x24, 0x78, 0x65, 0x05, 0x6b, 0xfc,
	0x0c, 0xe1, 0x1d, 0xbb, 0x77, 0xee, 0x1b, 0xe7, 0x9d, 0xcb, 0xde, 0xfa, 0xe0, 0x15, 0xfd, 0xff,
	0x7c, 0xc5, 0x5b, 0x87, 0x16, 0xa3, 0xbe, 0x89, 0xfd, 0x6c, 0x64, 0x18, 0xf5, 0x24, 0x80, 0x6e,
	0x63, 0x32, 0xe7, 0x59, 0x85, 0x1e, 0xac, 0xa1, 0xec, 0xda, 0xce, 0x2e, 0xa6, 0x82, 0x25, 0x75,
	0x91, 0x38, 0x9f, 0xc6, 0xad, 0x98, 0x43, 0x91, 0xf1, 0x3c, 0x02, 0x99, 0xc7, 0x39, 0xe5, 0x06,
	0x3d, 0xfa, 0x9d, 0x5e, 0x66, 0x82, 0xfd, 0x74, 0xbd, 0x79, 0x32, 0xfb, 0xe5, 0x5e, 0xcf, 0x9b,
	0x9e, 0xc4, 0xce, 0xf9, 0x91, 0x16, 0xc5, 0x1b, 0x0e, 0x3f, 0x78, 0x1d, 0x9e, 0x5a, 0xf9, 0xa6,
	0xf1, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x07, 0x73, 0x11, 0x97, 0x03, 0x00, 0x00,
}
