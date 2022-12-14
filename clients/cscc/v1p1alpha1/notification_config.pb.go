// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1p1alpha1/notification_config.proto

package securitycenter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The type of events.
type NotificationConfig_EventType int32

const (
	// Unspecified event type.
	NotificationConfig_EVENT_TYPE_UNSPECIFIED NotificationConfig_EventType = 0
	// Events for findings.
	NotificationConfig_FINDING NotificationConfig_EventType = 1
)

var NotificationConfig_EventType_name = map[int32]string{
	0: "EVENT_TYPE_UNSPECIFIED",
	1: "FINDING",
}

var NotificationConfig_EventType_value = map[string]int32{
	"EVENT_TYPE_UNSPECIFIED": 0,
	"FINDING":                1,
}

func (x NotificationConfig_EventType) String() string {
	return proto.EnumName(NotificationConfig_EventType_name, int32(x))
}

func (NotificationConfig_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cc064b43fa79af45, []int{0, 0}
}

// Cloud Security Command Center (Cloud SCC) notification configs.
//
// A notification config is a Cloud SCC resource that contains the configuration
// to send notifications for create/update events of findings, assets and etc.
type NotificationConfig struct {
	// The relative resource name of this notification config. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/notificationConfigs/notify_public_bucket".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the notification config (max of 1024 characters).
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The type of events the config is for, e.g. FINDING.
	EventType NotificationConfig_EventType `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=google.cloud.securitycenter.v1p1alpha1.NotificationConfig_EventType" json:"event_type,omitempty"`
	// The PubSub topic to send notifications to. Its format is
	// "projects/[project_id]/topics/[topic]".
	PubsubTopic string `protobuf:"bytes,4,opt,name=pubsub_topic,json=pubsubTopic,proto3" json:"pubsub_topic,omitempty"`
	// Output only. The service account that needs "pubsub.topics.publish"
	// permission to publish to the PubSub topic.
	ServiceAccount string `protobuf:"bytes,5,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// The config for triggering notifications.
	//
	// Types that are valid to be assigned to NotifyConfig:
	//	*NotificationConfig_StreamingConfig_
	NotifyConfig         isNotificationConfig_NotifyConfig `protobuf_oneof:"notify_config"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *NotificationConfig) Reset()         { *m = NotificationConfig{} }
func (m *NotificationConfig) String() string { return proto.CompactTextString(m) }
func (*NotificationConfig) ProtoMessage()    {}
func (*NotificationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc064b43fa79af45, []int{0}
}

func (m *NotificationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationConfig.Unmarshal(m, b)
}
func (m *NotificationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationConfig.Marshal(b, m, deterministic)
}
func (m *NotificationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationConfig.Merge(m, src)
}
func (m *NotificationConfig) XXX_Size() int {
	return xxx_messageInfo_NotificationConfig.Size(m)
}
func (m *NotificationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationConfig proto.InternalMessageInfo

func (m *NotificationConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NotificationConfig) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NotificationConfig) GetEventType() NotificationConfig_EventType {
	if m != nil {
		return m.EventType
	}
	return NotificationConfig_EVENT_TYPE_UNSPECIFIED
}

func (m *NotificationConfig) GetPubsubTopic() string {
	if m != nil {
		return m.PubsubTopic
	}
	return ""
}

func (m *NotificationConfig) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

type isNotificationConfig_NotifyConfig interface {
	isNotificationConfig_NotifyConfig()
}

type NotificationConfig_StreamingConfig_ struct {
	StreamingConfig *NotificationConfig_StreamingConfig `protobuf:"bytes,6,opt,name=streaming_config,json=streamingConfig,proto3,oneof"`
}

func (*NotificationConfig_StreamingConfig_) isNotificationConfig_NotifyConfig() {}

func (m *NotificationConfig) GetNotifyConfig() isNotificationConfig_NotifyConfig {
	if m != nil {
		return m.NotifyConfig
	}
	return nil
}

func (m *NotificationConfig) GetStreamingConfig() *NotificationConfig_StreamingConfig {
	if x, ok := m.GetNotifyConfig().(*NotificationConfig_StreamingConfig_); ok {
		return x.StreamingConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NotificationConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NotificationConfig_StreamingConfig_)(nil),
	}
}

// The config for streaming-based notifications, which send each event as soon
// as it is detected.
type NotificationConfig_StreamingConfig struct {
	// Expression that defines the filter to apply across create/update events
	// of assets or findings as specified by the event type. The expression is a
	// list of zero or more restrictions combined via logical operators `AND`
	// and `OR`. Parentheses are supported, and `OR` has higher precedence than
	// `AND`.
	//
	// Restrictions have the form `<field> <operator> <value>` and may have a
	// `-` character in front of them to indicate negation. The fields map to
	// those defined in the corresponding resource.
	//
	// The supported operators are:
	//
	// * `=` for all value types.
	// * `>`, `<`, `>=`, `<=` for integer values.
	// * `:`, meaning substring matching, for strings.
	//
	// The supported value types are:
	//
	// * string literals in quotes.
	// * integer literals without quotes.
	// * boolean literals `true` and `false` without quotes.
	Filter               string   `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationConfig_StreamingConfig) Reset()         { *m = NotificationConfig_StreamingConfig{} }
func (m *NotificationConfig_StreamingConfig) String() string { return proto.CompactTextString(m) }
func (*NotificationConfig_StreamingConfig) ProtoMessage()    {}
func (*NotificationConfig_StreamingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc064b43fa79af45, []int{0, 0}
}

func (m *NotificationConfig_StreamingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationConfig_StreamingConfig.Unmarshal(m, b)
}
func (m *NotificationConfig_StreamingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationConfig_StreamingConfig.Marshal(b, m, deterministic)
}
func (m *NotificationConfig_StreamingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationConfig_StreamingConfig.Merge(m, src)
}
func (m *NotificationConfig_StreamingConfig) XXX_Size() int {
	return xxx_messageInfo_NotificationConfig_StreamingConfig.Size(m)
}
func (m *NotificationConfig_StreamingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationConfig_StreamingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationConfig_StreamingConfig proto.InternalMessageInfo

func (m *NotificationConfig_StreamingConfig) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.securitycenter.v1p1alpha1.NotificationConfig_EventType", NotificationConfig_EventType_name, NotificationConfig_EventType_value)
	proto.RegisterType((*NotificationConfig)(nil), "google.cloud.securitycenter.v1p1alpha1.NotificationConfig")
	proto.RegisterType((*NotificationConfig_StreamingConfig)(nil), "google.cloud.securitycenter.v1p1alpha1.NotificationConfig.StreamingConfig")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1p1alpha1/notification_config.proto", fileDescriptor_cc064b43fa79af45)
}

var fileDescriptor_cc064b43fa79af45 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x6e, 0x14, 0xd5, 0x85, 0xb5, 0xf2, 0x61, 0x8a, 0x2a, 0x0e, 0x65, 0x87, 0xd1,
	0x71, 0x48, 0x94, 0xc1, 0x29, 0x5c, 0xd8, 0xda, 0x6c, 0x94, 0x43, 0x54, 0xb5, 0xa1, 0x12, 0xa8,
	0x52, 0xe4, 0x7a, 0xaf, 0xc1, 0x52, 0x6a, 0x5b, 0x89, 0x53, 0xd4, 0x2b, 0x1f, 0x87, 0x8f, 0xc2,
	0x07, 0xe1, 0xc0, 0xa7, 0x40, 0xb5, 0xdd, 0xb1, 0xc0, 0x81, 0x49, 0xdc, 0xec, 0xff, 0xfb, 0xf9,
	0xff, 0xfe, 0x7e, 0x7a, 0xe8, 0x6d, 0x26, 0x44, 0x96, 0x83, 0x4f, 0x73, 0x51, 0xdd, 0xfa, 0x25,
	0xd0, 0xaa, 0x60, 0x6a, 0x4b, 0x81, 0x2b, 0x28, 0xfc, 0x4d, 0x20, 0x03, 0x92, 0xcb, 0xcf, 0x24,
	0xf0, 0xb9, 0x50, 0x6c, 0xc5, 0x28, 0x51, 0x4c, 0xf0, 0x94, 0x0a, 0xbe, 0x62, 0x99, 0x27, 0x0b,
	0xa1, 0x04, 0x3e, 0x33, 0x0e, 0x9e, 0x76, 0xf0, 0xea, 0x0e, 0xde, 0x6f, 0x87, 0xde, 0x33, 0xdb,
	0x89, 0x48, 0xe6, 0x13, 0xce, 0x85, 0xd2, 0x5e, 0xa5, 0x71, 0x39, 0xfd, 0x71, 0x88, 0x70, 0x7c,
	0xaf, 0xc7, 0x50, 0xb7, 0xc0, 0x18, 0x1d, 0x71, 0xb2, 0x06, 0xd7, 0xe9, 0x3b, 0x83, 0xd6, 0x54,
	0x9f, 0x71, 0x1f, 0xb5, 0x6f, 0xa1, 0xa4, 0x05, 0x93, 0x3b, 0xd0, 0x6d, 0xe8, 0xd2, 0x7d, 0x09,
	0x53, 0x84, 0x60, 0x03, 0x5c, 0xa5, 0x6a, 0x2b, 0xc1, 0x3d, 0xec, 0x3b, 0x83, 0xe3, 0x8b, 0x91,
	0xf7, 0xb0, 0x9c, 0xde, 0xdf, 0x29, 0xbc, 0x68, 0x67, 0x96, 0x6c, 0x25, 0x4c, 0x5b, 0xb0, 0x3f,
	0xe2, 0xe7, 0xe8, 0x89, 0xac, 0x96, 0x65, 0xb5, 0x4c, 0x95, 0x90, 0x8c, 0xba, 0x47, 0x26, 0x87,
	0xd1, 0x92, 0x9d, 0x84, 0x5f, 0xa0, 0x4e, 0x09, 0xc5, 0x86, 0x51, 0x48, 0x09, 0xa5, 0xa2, 0xe2,
	0xca, 0x7d, 0xa4, 0xa9, 0x63, 0x2b, 0x5f, 0x1a, 0x15, 0x7f, 0x41, 0xdd, 0x52, 0x15, 0x40, 0xd6,
	0x8c, 0x67, 0x76, 0xba, 0x6e, 0xb3, 0xef, 0x0c, 0xda, 0x17, 0xef, 0xff, 0x23, 0xf6, 0x6c, 0x6f,
	0x69, 0xee, 0xef, 0x0e, 0xa6, 0x9d, 0xb2, 0x2e, 0xf5, 0xce, 0x51, 0xe7, 0x0f, 0x0a, 0x9f, 0xa0,
	0xe6, 0x8a, 0xe5, 0x0a, 0x0a, 0x3b, 0x74, 0x7b, 0x3b, 0x7d, 0x8d, 0x5a, 0x77, 0x73, 0xc0, 0x3d,
	0x74, 0x12, 0xcd, 0xa3, 0x38, 0x49, 0x93, 0x8f, 0x93, 0x28, 0xfd, 0x10, 0xcf, 0x26, 0xd1, 0x70,
	0x7c, 0x3d, 0x8e, 0x46, 0xdd, 0x03, 0xdc, 0x46, 0x8f, 0xaf, 0xc7, 0xf1, 0x68, 0x1c, 0xdf, 0x74,
	0x9d, 0xab, 0x0e, 0x7a, 0xaa, 0x57, 0x67, 0x6b, 0xbf, 0x75, 0xf5, 0xb5, 0x81, 0x5e, 0x52, 0xb1,
	0x7e, 0xe0, 0xb7, 0x26, 0xce, 0xa7, 0xc4, 0x92, 0x99, 0xc8, 0x09, 0xcf, 0x3c, 0x51, 0x64, 0x7e,
	0x06, 0x5c, 0x6f, 0x8d, 0x6f, 0x4a, 0x44, 0xb2, 0xf2, 0x5f, 0x0b, 0xfc, 0xa6, 0x5e, 0xf9, 0xd6,
	0x38, 0xbb, 0x31, 0xb6, 0x43, 0x1d, 0x60, 0x66, 0xab, 0x43, 0x13, 0x60, 0x1e, 0x4c, 0x82, 0x4b,
	0xfd, 0xee, 0xfb, 0x1e, 0x5c, 0x68, 0x70, 0x51, 0x07, 0x17, 0xf3, 0xbb, 0x06, 0x3f, 0x1b, 0xe7,
	0x06, 0x0c, 0x43, 0x4d, 0x86, 0x61, 0x1d, 0x0d, 0xc3, 0x1d, 0x6b, 0x4c, 0x97, 0x4d, 0x1d, 0xff,
	0xd5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x9f, 0x4c, 0xb3, 0x7e, 0x03, 0x00, 0x00,
}
