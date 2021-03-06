// Code generated by protoc-gen-go. DO NOT EDIT.
// source: analyzer.proto

package pulumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _struct "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// EnforcementLevel indicates the severity of a policy violation.
type EnforcementLevel int32

const (
	EnforcementLevel_ADVISORY  EnforcementLevel = 0
	EnforcementLevel_MANDATORY EnforcementLevel = 1
)

var EnforcementLevel_name = map[int32]string{
	0: "ADVISORY",
	1: "MANDATORY",
}
var EnforcementLevel_value = map[string]int32{
	"ADVISORY":  0,
	"MANDATORY": 1,
}

func (x EnforcementLevel) String() string {
	return proto.EnumName(EnforcementLevel_name, int32(x))
}
func (EnforcementLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_analyzer_a9c30ddfcaef9aa8, []int{0}
}

type AnalyzeRequest struct {
	Type                 string          `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties           *_struct.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
	Urn                  string          `protobuf:"bytes,3,opt,name=urn" json:"urn,omitempty"`
	Name                 string          `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AnalyzeRequest) Reset()         { *m = AnalyzeRequest{} }
func (m *AnalyzeRequest) String() string { return proto.CompactTextString(m) }
func (*AnalyzeRequest) ProtoMessage()    {}
func (*AnalyzeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyzer_a9c30ddfcaef9aa8, []int{0}
}
func (m *AnalyzeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzeRequest.Unmarshal(m, b)
}
func (m *AnalyzeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzeRequest.Marshal(b, m, deterministic)
}
func (dst *AnalyzeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzeRequest.Merge(dst, src)
}
func (m *AnalyzeRequest) XXX_Size() int {
	return xxx_messageInfo_AnalyzeRequest.Size(m)
}
func (m *AnalyzeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzeRequest proto.InternalMessageInfo

func (m *AnalyzeRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AnalyzeRequest) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *AnalyzeRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *AnalyzeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Resource defines the view of a Pulumi-managed resource as sent to Analyzers. The properties
// of the resource are specific to the type of analysis being performed. See the Analyzer
// service definition for more information.
type AnalyzerResource struct {
	Type                 string          `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties           *_struct.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
	Urn                  string          `protobuf:"bytes,3,opt,name=urn" json:"urn,omitempty"`
	Name                 string          `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AnalyzerResource) Reset()         { *m = AnalyzerResource{} }
func (m *AnalyzerResource) String() string { return proto.CompactTextString(m) }
func (*AnalyzerResource) ProtoMessage()    {}
func (*AnalyzerResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyzer_a9c30ddfcaef9aa8, []int{1}
}
func (m *AnalyzerResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzerResource.Unmarshal(m, b)
}
func (m *AnalyzerResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzerResource.Marshal(b, m, deterministic)
}
func (dst *AnalyzerResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzerResource.Merge(dst, src)
}
func (m *AnalyzerResource) XXX_Size() int {
	return xxx_messageInfo_AnalyzerResource.Size(m)
}
func (m *AnalyzerResource) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzerResource.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzerResource proto.InternalMessageInfo

func (m *AnalyzerResource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AnalyzerResource) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *AnalyzerResource) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *AnalyzerResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AnalyzeStackRequest struct {
	Resources            []*AnalyzerResource `protobuf:"bytes,1,rep,name=resources" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AnalyzeStackRequest) Reset()         { *m = AnalyzeStackRequest{} }
func (m *AnalyzeStackRequest) String() string { return proto.CompactTextString(m) }
func (*AnalyzeStackRequest) ProtoMessage()    {}
func (*AnalyzeStackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyzer_a9c30ddfcaef9aa8, []int{2}
}
func (m *AnalyzeStackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzeStackRequest.Unmarshal(m, b)
}
func (m *AnalyzeStackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzeStackRequest.Marshal(b, m, deterministic)
}
func (dst *AnalyzeStackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzeStackRequest.Merge(dst, src)
}
func (m *AnalyzeStackRequest) XXX_Size() int {
	return xxx_messageInfo_AnalyzeStackRequest.Size(m)
}
func (m *AnalyzeStackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzeStackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzeStackRequest proto.InternalMessageInfo

func (m *AnalyzeStackRequest) GetResources() []*AnalyzerResource {
	if m != nil {
		return m.Resources
	}
	return nil
}

type AnalyzeResponse struct {
	Diagnostics          []*AnalyzeDiagnostic `protobuf:"bytes,2,rep,name=diagnostics" json:"diagnostics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AnalyzeResponse) Reset()         { *m = AnalyzeResponse{} }
func (m *AnalyzeResponse) String() string { return proto.CompactTextString(m) }
func (*AnalyzeResponse) ProtoMessage()    {}
func (*AnalyzeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyzer_a9c30ddfcaef9aa8, []int{3}
}
func (m *AnalyzeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzeResponse.Unmarshal(m, b)
}
func (m *AnalyzeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzeResponse.Marshal(b, m, deterministic)
}
func (dst *AnalyzeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzeResponse.Merge(dst, src)
}
func (m *AnalyzeResponse) XXX_Size() int {
	return xxx_messageInfo_AnalyzeResponse.Size(m)
}
func (m *AnalyzeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzeResponse proto.InternalMessageInfo

func (m *AnalyzeResponse) GetDiagnostics() []*AnalyzeDiagnostic {
	if m != nil {
		return m.Diagnostics
	}
	return nil
}

type AnalyzeDiagnostic struct {
	PolicyName           string           `protobuf:"bytes,1,opt,name=policyName" json:"policyName,omitempty"`
	PolicyPackName       string           `protobuf:"bytes,2,opt,name=policyPackName" json:"policyPackName,omitempty"`
	PolicyPackVersion    string           `protobuf:"bytes,3,opt,name=policyPackVersion" json:"policyPackVersion,omitempty"`
	Description          string           `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Message              string           `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
	Tags                 []string         `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	EnforcementLevel     EnforcementLevel `protobuf:"varint,7,opt,name=enforcementLevel,enum=pulumirpc.EnforcementLevel" json:"enforcementLevel,omitempty"`
	Urn                  string           `protobuf:"bytes,8,opt,name=urn" json:"urn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AnalyzeDiagnostic) Reset()         { *m = AnalyzeDiagnostic{} }
func (m *AnalyzeDiagnostic) String() string { return proto.CompactTextString(m) }
func (*AnalyzeDiagnostic) ProtoMessage()    {}
func (*AnalyzeDiagnostic) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyzer_a9c30ddfcaef9aa8, []int{4}
}
func (m *AnalyzeDiagnostic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzeDiagnostic.Unmarshal(m, b)
}
func (m *AnalyzeDiagnostic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzeDiagnostic.Marshal(b, m, deterministic)
}
func (dst *AnalyzeDiagnostic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzeDiagnostic.Merge(dst, src)
}
func (m *AnalyzeDiagnostic) XXX_Size() int {
	return xxx_messageInfo_AnalyzeDiagnostic.Size(m)
}
func (m *AnalyzeDiagnostic) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzeDiagnostic.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzeDiagnostic proto.InternalMessageInfo

func (m *AnalyzeDiagnostic) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *AnalyzeDiagnostic) GetPolicyPackName() string {
	if m != nil {
		return m.PolicyPackName
	}
	return ""
}

func (m *AnalyzeDiagnostic) GetPolicyPackVersion() string {
	if m != nil {
		return m.PolicyPackVersion
	}
	return ""
}

func (m *AnalyzeDiagnostic) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AnalyzeDiagnostic) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AnalyzeDiagnostic) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *AnalyzeDiagnostic) GetEnforcementLevel() EnforcementLevel {
	if m != nil {
		return m.EnforcementLevel
	}
	return EnforcementLevel_ADVISORY
}

func (m *AnalyzeDiagnostic) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

// AnalyzerInfo provides metadata about a PolicyPack inside an analyzer.
type AnalyzerInfo struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DisplayName          string        `protobuf:"bytes,2,opt,name=displayName" json:"displayName,omitempty"`
	Policies             []*PolicyInfo `protobuf:"bytes,3,rep,name=policies" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AnalyzerInfo) Reset()         { *m = AnalyzerInfo{} }
func (m *AnalyzerInfo) String() string { return proto.CompactTextString(m) }
func (*AnalyzerInfo) ProtoMessage()    {}
func (*AnalyzerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyzer_a9c30ddfcaef9aa8, []int{5}
}
func (m *AnalyzerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzerInfo.Unmarshal(m, b)
}
func (m *AnalyzerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzerInfo.Marshal(b, m, deterministic)
}
func (dst *AnalyzerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzerInfo.Merge(dst, src)
}
func (m *AnalyzerInfo) XXX_Size() int {
	return xxx_messageInfo_AnalyzerInfo.Size(m)
}
func (m *AnalyzerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzerInfo proto.InternalMessageInfo

func (m *AnalyzerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AnalyzerInfo) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *AnalyzerInfo) GetPolicies() []*PolicyInfo {
	if m != nil {
		return m.Policies
	}
	return nil
}

// PolicyInfo provides metadata about an individual Policy within a Policy Pack.
type PolicyInfo struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DisplayName          string           `protobuf:"bytes,2,opt,name=displayName" json:"displayName,omitempty"`
	Description          string           `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Message              string           `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	EnforcementLevel     EnforcementLevel `protobuf:"varint,5,opt,name=enforcementLevel,enum=pulumirpc.EnforcementLevel" json:"enforcementLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PolicyInfo) Reset()         { *m = PolicyInfo{} }
func (m *PolicyInfo) String() string { return proto.CompactTextString(m) }
func (*PolicyInfo) ProtoMessage()    {}
func (*PolicyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_analyzer_a9c30ddfcaef9aa8, []int{6}
}
func (m *PolicyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyInfo.Unmarshal(m, b)
}
func (m *PolicyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyInfo.Marshal(b, m, deterministic)
}
func (dst *PolicyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyInfo.Merge(dst, src)
}
func (m *PolicyInfo) XXX_Size() int {
	return xxx_messageInfo_PolicyInfo.Size(m)
}
func (m *PolicyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyInfo proto.InternalMessageInfo

func (m *PolicyInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PolicyInfo) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *PolicyInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PolicyInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PolicyInfo) GetEnforcementLevel() EnforcementLevel {
	if m != nil {
		return m.EnforcementLevel
	}
	return EnforcementLevel_ADVISORY
}

func init() {
	proto.RegisterType((*AnalyzeRequest)(nil), "pulumirpc.AnalyzeRequest")
	proto.RegisterType((*AnalyzerResource)(nil), "pulumirpc.AnalyzerResource")
	proto.RegisterType((*AnalyzeStackRequest)(nil), "pulumirpc.AnalyzeStackRequest")
	proto.RegisterType((*AnalyzeResponse)(nil), "pulumirpc.AnalyzeResponse")
	proto.RegisterType((*AnalyzeDiagnostic)(nil), "pulumirpc.AnalyzeDiagnostic")
	proto.RegisterType((*AnalyzerInfo)(nil), "pulumirpc.AnalyzerInfo")
	proto.RegisterType((*PolicyInfo)(nil), "pulumirpc.PolicyInfo")
	proto.RegisterEnum("pulumirpc.EnforcementLevel", EnforcementLevel_name, EnforcementLevel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Analyzer service

type AnalyzerClient interface {
	// Analyze analyzes a single resource object, and returns any errors that it finds.
	// Called with the "inputs" to the resource, before it is updated.
	Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
	// AnalyzeStack analyzes all resources within a stack, at the end of a successful
	// preview or update. The provided resources are the "outputs", after any mutations
	// have taken place.
	AnalyzeStack(ctx context.Context, in *AnalyzeStackRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
	// GetAnalyzerInfo returns metadata about the analyzer (e.g., list of policies contained).
	GetAnalyzerInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AnalyzerInfo, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PluginInfo, error)
}

type analyzerClient struct {
	cc *grpc.ClientConn
}

func NewAnalyzerClient(cc *grpc.ClientConn) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	out := new(AnalyzeResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.Analyzer/Analyze", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) AnalyzeStack(ctx context.Context, in *AnalyzeStackRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	out := new(AnalyzeResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.Analyzer/AnalyzeStack", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) GetAnalyzerInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AnalyzerInfo, error) {
	out := new(AnalyzerInfo)
	err := grpc.Invoke(ctx, "/pulumirpc.Analyzer/GetAnalyzerInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) GetPluginInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := grpc.Invoke(ctx, "/pulumirpc.Analyzer/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Analyzer service

type AnalyzerServer interface {
	// Analyze analyzes a single resource object, and returns any errors that it finds.
	// Called with the "inputs" to the resource, before it is updated.
	Analyze(context.Context, *AnalyzeRequest) (*AnalyzeResponse, error)
	// AnalyzeStack analyzes all resources within a stack, at the end of a successful
	// preview or update. The provided resources are the "outputs", after any mutations
	// have taken place.
	AnalyzeStack(context.Context, *AnalyzeStackRequest) (*AnalyzeResponse, error)
	// GetAnalyzerInfo returns metadata about the analyzer (e.g., list of policies contained).
	GetAnalyzerInfo(context.Context, *empty.Empty) (*AnalyzerInfo, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(context.Context, *empty.Empty) (*PluginInfo, error)
}

func RegisterAnalyzerServer(s *grpc.Server, srv AnalyzerServer) {
	s.RegisterService(&_Analyzer_serviceDesc, srv)
}

func _Analyzer_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Analyzer/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Analyze(ctx, req.(*AnalyzeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_AnalyzeStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeStackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).AnalyzeStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Analyzer/AnalyzeStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).AnalyzeStack(ctx, req.(*AnalyzeStackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_GetAnalyzerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetAnalyzerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Analyzer/GetAnalyzerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetAnalyzerInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Analyzer/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetPluginInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analyzer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _Analyzer_Analyze_Handler,
		},
		{
			MethodName: "AnalyzeStack",
			Handler:    _Analyzer_AnalyzeStack_Handler,
		},
		{
			MethodName: "GetAnalyzerInfo",
			Handler:    _Analyzer_GetAnalyzerInfo_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _Analyzer_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyzer.proto",
}

func init() { proto.RegisterFile("analyzer.proto", fileDescriptor_analyzer_a9c30ddfcaef9aa8) }

var fileDescriptor_analyzer_a9c30ddfcaef9aa8 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x54, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0x6e, 0x9a, 0x6e, 0x6d, 0x4f, 0xbb, 0xae, 0xf3, 0x4f, 0x3f, 0x16, 0xba, 0x69, 0x8a, 0x72,
	0x81, 0x2a, 0x84, 0x52, 0x51, 0x2e, 0x10, 0x37, 0x88, 0xa2, 0x4e, 0xd5, 0xa4, 0x31, 0x4a, 0x8a,
	0x26, 0x71, 0x99, 0x65, 0x6e, 0x14, 0x2d, 0x8d, 0x8d, 0xed, 0x80, 0xca, 0x2d, 0xe2, 0x4d, 0x78,
	0x04, 0xde, 0x81, 0xd7, 0x42, 0x76, 0xfe, 0x34, 0x4d, 0xaa, 0x21, 0xed, 0x86, 0x3b, 0xfb, 0x9c,
	0xcf, 0x9f, 0x8f, 0xbf, 0xef, 0x1c, 0x43, 0xcf, 0x8d, 0xdc, 0x70, 0xfd, 0x0d, 0x33, 0x9b, 0x32,
	0x22, 0x08, 0x6a, 0xd3, 0x38, 0x8c, 0x57, 0x01, 0xa3, 0xde, 0xa0, 0x4b, 0xc3, 0xd8, 0x0f, 0xa2,
	0x24, 0x31, 0x38, 0xf1, 0x09, 0xf1, 0x43, 0x3c, 0x52, 0xbb, 0x9b, 0x78, 0x39, 0xc2, 0x2b, 0x2a,
	0xd6, 0x69, 0xf2, 0xb4, 0x9c, 0xe4, 0x82, 0xc5, 0x9e, 0x48, 0xb2, 0xd6, 0x77, 0x0d, 0x7a, 0x93,
	0xe4, 0x1a, 0x07, 0x7f, 0x8e, 0x31, 0x17, 0x08, 0x41, 0x43, 0xac, 0x29, 0x36, 0x34, 0x53, 0x1b,
	0xb6, 0x1d, 0xb5, 0x46, 0x2f, 0x01, 0x28, 0x23, 0x14, 0x33, 0x11, 0x60, 0x6e, 0xd4, 0x4d, 0x6d,
	0xd8, 0x19, 0x1f, 0xdb, 0x09, 0xb3, 0x9d, 0x31, 0xdb, 0x0b, 0xc5, 0xec, 0x14, 0xa0, 0xa8, 0x0f,
	0x7a, 0xcc, 0x22, 0x43, 0x57, 0x5c, 0x72, 0x29, 0xe9, 0x23, 0x77, 0x85, 0x8d, 0x46, 0x42, 0x2f,
	0xd7, 0xd6, 0x0f, 0x0d, 0xfa, 0x69, 0x15, 0xcc, 0xc1, 0x9c, 0xc4, 0xcc, 0xc3, 0xff, 0xa2, 0x8e,
	0x39, 0xfc, 0x97, 0x96, 0xb1, 0x10, 0xae, 0x77, 0x97, 0x29, 0xf2, 0x0a, 0xda, 0x2c, 0xad, 0x8a,
	0x1b, 0x9a, 0xa9, 0x0f, 0x3b, 0xe3, 0x13, 0x3b, 0x37, 0xc3, 0x2e, 0x57, 0xee, 0x6c, 0xd0, 0xd6,
	0x07, 0x38, 0xcc, 0xe5, 0xe5, 0x94, 0x44, 0x1c, 0xa3, 0xd7, 0xd0, 0xb9, 0x0d, 0x5c, 0x3f, 0x22,
	0x5c, 0x04, 0x9e, 0x7c, 0x84, 0xe4, 0x3b, 0xad, 0xf2, 0x4d, 0x73, 0x90, 0x53, 0x3c, 0x60, 0xfd,
	0xaa, 0xc3, 0x51, 0x05, 0x82, 0xce, 0x00, 0x28, 0x09, 0x03, 0x6f, 0x7d, 0x25, 0x1f, 0x95, 0x68,
	0x56, 0x88, 0xa0, 0x27, 0xd0, 0x4b, 0x76, 0x73, 0xd7, 0xbb, 0x53, 0x98, 0xba, 0xc2, 0x94, 0xa2,
	0xe8, 0x19, 0x1c, 0x6d, 0x22, 0xd7, 0x98, 0xf1, 0x80, 0x64, 0xb2, 0x55, 0x13, 0xc8, 0x84, 0xce,
	0x2d, 0xe6, 0x1e, 0x0b, 0xa8, 0x90, 0xb8, 0x44, 0xcb, 0x62, 0x08, 0x19, 0xd0, 0x5c, 0x61, 0xce,
	0x5d, 0x1f, 0x1b, 0x7b, 0x2a, 0x9b, 0x6d, 0x95, 0xbf, 0xae, 0xcf, 0x8d, 0x7d, 0x53, 0x57, 0xfe,
	0xba, 0x3e, 0x47, 0x33, 0xe8, 0xe3, 0x68, 0x49, 0x98, 0x87, 0x57, 0x38, 0x12, 0x97, 0xf8, 0x0b,
	0x0e, 0x8d, 0xa6, 0xa9, 0x0d, 0x7b, 0x5b, 0x82, 0x9f, 0x97, 0x20, 0x4e, 0xe5, 0x50, 0xe6, 0x77,
	0x2b, 0xf7, 0xdb, 0xfa, 0x0a, 0xdd, 0xcc, 0xa8, 0x8b, 0x68, 0x49, 0x72, 0xff, 0xb5, 0x8d, 0xff,
	0xea, 0x39, 0x01, 0xa7, 0xa1, 0xbb, 0x2e, 0x28, 0x54, 0x0c, 0xa1, 0xe7, 0xd0, 0x52, 0x2a, 0xc8,
	0xf6, 0xd3, 0x95, 0x73, 0xff, 0x17, 0x0a, 0x9b, 0x2b, 0x81, 0x24, 0xbd, 0x93, 0xc3, 0xac, 0xdf,
	0x1a, 0xc0, 0x26, 0xf1, 0xc0, 0x7b, 0x4b, 0x42, 0xeb, 0xf7, 0x0a, 0xdd, 0xd8, 0x16, 0x7a, 0x97,
	0xa8, 0x7b, 0x0f, 0x10, 0xf5, 0xe9, 0x08, 0xfa, 0x65, 0x14, 0xea, 0x42, 0x6b, 0x32, 0xbd, 0xbe,
	0x58, 0xbc, 0x77, 0x3e, 0xf5, 0x6b, 0xe8, 0x00, 0xda, 0xef, 0x26, 0x57, 0xd3, 0xc9, 0x47, 0xb9,
	0xd5, 0xc6, 0x3f, 0xeb, 0xd0, 0xca, 0x44, 0x47, 0x6f, 0xa1, 0x99, 0xae, 0xd1, 0xe3, 0x6a, 0xb7,
	0xa7, 0xb3, 0x36, 0x18, 0xec, 0x4a, 0x25, 0x93, 0x63, 0xd5, 0xd0, 0x65, 0x6e, 0xa2, 0x1a, 0x50,
	0x74, 0x56, 0x45, 0x17, 0x27, 0xf7, 0x2f, 0x6c, 0x53, 0x38, 0x9c, 0x61, 0xb1, 0xd5, 0x15, 0x8f,
	0x2a, 0x9f, 0xc9, 0xb9, 0xfc, 0x4b, 0x07, 0xc7, 0x3b, 0xe6, 0x5d, 0x1e, 0xb0, 0x6a, 0xe8, 0x0d,
	0x1c, 0xcc, 0xb0, 0x98, 0xab, 0x0f, 0xf9, 0x5e, 0x8e, 0xad, 0x4e, 0xc9, 0xe1, 0x56, 0xed, 0x66,
	0x5f, 0x01, 0x5f, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x63, 0x35, 0xde, 0x7f, 0xf1, 0x05, 0x00,
	0x00,
}
