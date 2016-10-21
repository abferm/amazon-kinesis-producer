// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package kinesis is a generated protocol buffer package.

It is generated from these files:
	config.proto
	messages.proto

It has these top-level messages:
	AdditionalDimension
	Configuration
	Tag
	Record
	AggregatedRecord
	Message
	PutRecord
	Flush
	Attempt
	PutRecordResult
	Credentials
	SetCredentials
	Dimension
	Stats
	Metric
	MetricsRequest
	MetricsResponse
*/
package kinesis

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AdditionalDimension struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	Granularity      *string `protobuf:"bytes,3,req,name=granularity" json:"granularity,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AdditionalDimension) Reset()                    { *m = AdditionalDimension{} }
func (m *AdditionalDimension) String() string            { return proto.CompactTextString(m) }
func (*AdditionalDimension) ProtoMessage()               {}
func (*AdditionalDimension) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AdditionalDimension) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *AdditionalDimension) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *AdditionalDimension) GetGranularity() string {
	if m != nil && m.Granularity != nil {
		return *m.Granularity
	}
	return ""
}

type Configuration struct {
	AdditionalMetricDims  []*AdditionalDimension `protobuf:"bytes,128,rep,name=additional_metric_dims" json:"additional_metric_dims,omitempty"`
	AggregationEnabled    *bool                  `protobuf:"varint,1,opt,name=aggregation_enabled,def=1" json:"aggregation_enabled,omitempty"`
	AggregationMaxCount   *uint64                `protobuf:"varint,2,opt,name=aggregation_max_count,def=4294967295" json:"aggregation_max_count,omitempty"`
	AggregationMaxSize    *uint64                `protobuf:"varint,3,opt,name=aggregation_max_size,def=51200" json:"aggregation_max_size,omitempty"`
	CloudwatchEndpoint    *string                `protobuf:"bytes,4,opt,name=cloudwatch_endpoint,def=" json:"cloudwatch_endpoint,omitempty"`
	CloudwatchPort        *uint64                `protobuf:"varint,5,opt,name=cloudwatch_port,def=443" json:"cloudwatch_port,omitempty"`
	CollectionMaxCount    *uint64                `protobuf:"varint,6,opt,name=collection_max_count,def=500" json:"collection_max_count,omitempty"`
	CollectionMaxSize     *uint64                `protobuf:"varint,7,opt,name=collection_max_size,def=5242880" json:"collection_max_size,omitempty"`
	ConnectTimeout        *uint64                `protobuf:"varint,8,opt,name=connect_timeout,def=6000" json:"connect_timeout,omitempty"`
	EnableCoreDumps       *bool                  `protobuf:"varint,9,opt,name=enable_core_dumps,def=0" json:"enable_core_dumps,omitempty"`
	FailIfThrottled       *bool                  `protobuf:"varint,10,opt,name=fail_if_throttled,def=0" json:"fail_if_throttled,omitempty"`
	KinesisEndpoint       *string                `protobuf:"bytes,11,opt,name=kinesis_endpoint,def=" json:"kinesis_endpoint,omitempty"`
	KinesisPort           *uint64                `protobuf:"varint,12,opt,name=kinesis_port,def=443" json:"kinesis_port,omitempty"`
	LogLevel              *string                `protobuf:"bytes,13,opt,name=log_level,def=info" json:"log_level,omitempty"`
	MaxConnections        *uint64                `protobuf:"varint,14,opt,name=max_connections,def=24" json:"max_connections,omitempty"`
	MetricsGranularity    *string                `protobuf:"bytes,15,opt,name=metrics_granularity,def=shard" json:"metrics_granularity,omitempty"`
	MetricsLevel          *string                `protobuf:"bytes,16,opt,name=metrics_level,def=detailed" json:"metrics_level,omitempty"`
	MetricsNamespace      *string                `protobuf:"bytes,17,opt,name=metrics_namespace,def=KinesisProducerLibrary" json:"metrics_namespace,omitempty"`
	MetricsUploadDelay    *uint64                `protobuf:"varint,18,opt,name=metrics_upload_delay,def=60000" json:"metrics_upload_delay,omitempty"`
	MinConnections        *uint64                `protobuf:"varint,19,opt,name=min_connections,def=1" json:"min_connections,omitempty"`
	RateLimit             *uint64                `protobuf:"varint,20,opt,name=rate_limit,def=150" json:"rate_limit,omitempty"`
	RecordMaxBufferedTime *uint64                `protobuf:"varint,21,opt,name=record_max_buffered_time,def=100" json:"record_max_buffered_time,omitempty"`
	RecordTtl             *uint64                `protobuf:"varint,22,opt,name=record_ttl,def=30000" json:"record_ttl,omitempty"`
	Region                *string                `protobuf:"bytes,23,opt,name=region,def=" json:"region,omitempty"`
	RequestTimeout        *uint64                `protobuf:"varint,24,opt,name=request_timeout,def=6000" json:"request_timeout,omitempty"`
	VerifyCertificate     *bool                  `protobuf:"varint,25,opt,name=verify_certificate,def=1" json:"verify_certificate,omitempty"`
	XXX_unrecognized      []byte                 `json:"-"`
}

func (m *Configuration) Reset()                    { *m = Configuration{} }
func (m *Configuration) String() string            { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()               {}
func (*Configuration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_Configuration_AggregationEnabled bool = true
const Default_Configuration_AggregationMaxCount uint64 = 4294967295
const Default_Configuration_AggregationMaxSize uint64 = 51200
const Default_Configuration_CloudwatchPort uint64 = 443
const Default_Configuration_CollectionMaxCount uint64 = 500
const Default_Configuration_CollectionMaxSize uint64 = 5242880
const Default_Configuration_ConnectTimeout uint64 = 6000
const Default_Configuration_EnableCoreDumps bool = false
const Default_Configuration_FailIfThrottled bool = false
const Default_Configuration_KinesisPort uint64 = 443
const Default_Configuration_LogLevel string = "info"
const Default_Configuration_MaxConnections uint64 = 24
const Default_Configuration_MetricsGranularity string = "shard"
const Default_Configuration_MetricsLevel string = "detailed"
const Default_Configuration_MetricsNamespace string = "KinesisProducerLibrary"
const Default_Configuration_MetricsUploadDelay uint64 = 60000
const Default_Configuration_MinConnections uint64 = 1
const Default_Configuration_RateLimit uint64 = 150
const Default_Configuration_RecordMaxBufferedTime uint64 = 100
const Default_Configuration_RecordTtl uint64 = 30000
const Default_Configuration_RequestTimeout uint64 = 6000
const Default_Configuration_VerifyCertificate bool = true

func (m *Configuration) GetAdditionalMetricDims() []*AdditionalDimension {
	if m != nil {
		return m.AdditionalMetricDims
	}
	return nil
}

func (m *Configuration) GetAggregationEnabled() bool {
	if m != nil && m.AggregationEnabled != nil {
		return *m.AggregationEnabled
	}
	return Default_Configuration_AggregationEnabled
}

func (m *Configuration) GetAggregationMaxCount() uint64 {
	if m != nil && m.AggregationMaxCount != nil {
		return *m.AggregationMaxCount
	}
	return Default_Configuration_AggregationMaxCount
}

func (m *Configuration) GetAggregationMaxSize() uint64 {
	if m != nil && m.AggregationMaxSize != nil {
		return *m.AggregationMaxSize
	}
	return Default_Configuration_AggregationMaxSize
}

func (m *Configuration) GetCloudwatchEndpoint() string {
	if m != nil && m.CloudwatchEndpoint != nil {
		return *m.CloudwatchEndpoint
	}
	return ""
}

func (m *Configuration) GetCloudwatchPort() uint64 {
	if m != nil && m.CloudwatchPort != nil {
		return *m.CloudwatchPort
	}
	return Default_Configuration_CloudwatchPort
}

func (m *Configuration) GetCollectionMaxCount() uint64 {
	if m != nil && m.CollectionMaxCount != nil {
		return *m.CollectionMaxCount
	}
	return Default_Configuration_CollectionMaxCount
}

func (m *Configuration) GetCollectionMaxSize() uint64 {
	if m != nil && m.CollectionMaxSize != nil {
		return *m.CollectionMaxSize
	}
	return Default_Configuration_CollectionMaxSize
}

func (m *Configuration) GetConnectTimeout() uint64 {
	if m != nil && m.ConnectTimeout != nil {
		return *m.ConnectTimeout
	}
	return Default_Configuration_ConnectTimeout
}

func (m *Configuration) GetEnableCoreDumps() bool {
	if m != nil && m.EnableCoreDumps != nil {
		return *m.EnableCoreDumps
	}
	return Default_Configuration_EnableCoreDumps
}

func (m *Configuration) GetFailIfThrottled() bool {
	if m != nil && m.FailIfThrottled != nil {
		return *m.FailIfThrottled
	}
	return Default_Configuration_FailIfThrottled
}

func (m *Configuration) GetKinesisEndpoint() string {
	if m != nil && m.KinesisEndpoint != nil {
		return *m.KinesisEndpoint
	}
	return ""
}

func (m *Configuration) GetKinesisPort() uint64 {
	if m != nil && m.KinesisPort != nil {
		return *m.KinesisPort
	}
	return Default_Configuration_KinesisPort
}

func (m *Configuration) GetLogLevel() string {
	if m != nil && m.LogLevel != nil {
		return *m.LogLevel
	}
	return Default_Configuration_LogLevel
}

func (m *Configuration) GetMaxConnections() uint64 {
	if m != nil && m.MaxConnections != nil {
		return *m.MaxConnections
	}
	return Default_Configuration_MaxConnections
}

func (m *Configuration) GetMetricsGranularity() string {
	if m != nil && m.MetricsGranularity != nil {
		return *m.MetricsGranularity
	}
	return Default_Configuration_MetricsGranularity
}

func (m *Configuration) GetMetricsLevel() string {
	if m != nil && m.MetricsLevel != nil {
		return *m.MetricsLevel
	}
	return Default_Configuration_MetricsLevel
}

func (m *Configuration) GetMetricsNamespace() string {
	if m != nil && m.MetricsNamespace != nil {
		return *m.MetricsNamespace
	}
	return Default_Configuration_MetricsNamespace
}

func (m *Configuration) GetMetricsUploadDelay() uint64 {
	if m != nil && m.MetricsUploadDelay != nil {
		return *m.MetricsUploadDelay
	}
	return Default_Configuration_MetricsUploadDelay
}

func (m *Configuration) GetMinConnections() uint64 {
	if m != nil && m.MinConnections != nil {
		return *m.MinConnections
	}
	return Default_Configuration_MinConnections
}

func (m *Configuration) GetRateLimit() uint64 {
	if m != nil && m.RateLimit != nil {
		return *m.RateLimit
	}
	return Default_Configuration_RateLimit
}

func (m *Configuration) GetRecordMaxBufferedTime() uint64 {
	if m != nil && m.RecordMaxBufferedTime != nil {
		return *m.RecordMaxBufferedTime
	}
	return Default_Configuration_RecordMaxBufferedTime
}

func (m *Configuration) GetRecordTtl() uint64 {
	if m != nil && m.RecordTtl != nil {
		return *m.RecordTtl
	}
	return Default_Configuration_RecordTtl
}

func (m *Configuration) GetRegion() string {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return ""
}

func (m *Configuration) GetRequestTimeout() uint64 {
	if m != nil && m.RequestTimeout != nil {
		return *m.RequestTimeout
	}
	return Default_Configuration_RequestTimeout
}

func (m *Configuration) GetVerifyCertificate() bool {
	if m != nil && m.VerifyCertificate != nil {
		return *m.VerifyCertificate
	}
	return Default_Configuration_VerifyCertificate
}

func init() {
	proto.RegisterType((*AdditionalDimension)(nil), "aws.kinesis.protobuf.AdditionalDimension")
	proto.RegisterType((*Configuration)(nil), "aws.kinesis.protobuf.Configuration")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0x4f, 0x53, 0xdb, 0x3a,
	0x14, 0xc5, 0x5f, 0xfe, 0x01, 0xb9, 0x21, 0x2f, 0xa0, 0x04, 0x10, 0xbc, 0xc7, 0x34, 0xd0, 0x76,
	0x26, 0x6c, 0x32, 0x4e, 0x48, 0x28, 0xb0, 0xeb, 0xb4, 0xd3, 0x45, 0xdb, 0x45, 0xbf, 0x81, 0x46,
	0x91, 0xae, 0x8d, 0x06, 0x59, 0x4a, 0x65, 0x19, 0x9a, 0xae, 0xba, 0xef, 0x97, 0xee, 0x58, 0x4e,
	0x9a, 0xc0, 0xb0, 0xb4, 0xee, 0xef, 0x5e, 0x9d, 0x7b, 0x8e, 0x0c, 0xbb, 0xc2, 0x9a, 0x58, 0x25,
	0xc3, 0xb9, 0xb3, 0xde, 0x92, 0x1e, 0x7f, 0xcc, 0x86, 0xf7, 0xca, 0x60, 0xa6, 0xb2, 0xf2, 0x68,
	0x96, 0xc7, 0xe7, 0x9f, 0xa0, 0xfb, 0x5e, 0x4a, 0xe5, 0x95, 0x35, 0x5c, 0x7f, 0x54, 0x29, 0x9a,
	0x4c, 0x59, 0x43, 0x5a, 0x50, 0xbb, 0xc7, 0x05, 0xad, 0xf4, 0xab, 0x83, 0x26, 0x69, 0x43, 0xe3,
	0x81, 0xeb, 0x1c, 0x69, 0x35, 0x7c, 0x76, 0xa1, 0x95, 0x38, 0x6e, 0x72, 0xcd, 0x9d, 0xf2, 0x0b,
	0x5a, 0x2b, 0x0e, 0xcf, 0x7f, 0x6f, 0x43, 0xfb, 0x43, 0xb8, 0x2e, 0x77, 0xbc, 0x98, 0x46, 0x3e,
	0xc3, 0x21, 0xff, 0x3b, 0x99, 0xa5, 0xe8, 0x9d, 0x12, 0x4c, 0xaa, 0x34, 0xa3, 0xbf, 0x2a, 0xfd,
	0xda, 0xa0, 0x35, 0xbe, 0x18, 0xbe, 0xa4, 0x68, 0xf8, 0x92, 0x9c, 0x33, 0xe8, 0xf2, 0x24, 0x71,
	0x98, 0x84, 0xd1, 0x0c, 0x0d, 0x9f, 0x69, 0x94, 0xb4, 0xd2, 0xaf, 0x0c, 0x76, 0x6e, 0xeb, 0xde,
	0xe5, 0x48, 0x2e, 0xe0, 0x60, 0x13, 0x49, 0xf9, 0x0f, 0x26, 0x6c, 0x6e, 0x3c, 0xad, 0xf6, 0x2b,
	0x83, 0xfa, 0x2d, 0x4c, 0xc6, 0x37, 0x93, 0x9b, 0xab, 0x77, 0xe3, 0x9b, 0x29, 0x79, 0x0d, 0xbd,
	0xe7, 0x68, 0xa6, 0x7e, 0x22, 0xad, 0x05, 0xb2, 0x31, 0x1d, 0x8d, 0xa3, 0x88, 0x9c, 0x42, 0x57,
	0x68, 0x9b, 0xcb, 0x47, 0xee, 0xc5, 0x1d, 0x43, 0x23, 0xe7, 0x56, 0x19, 0x4f, 0xeb, 0xfd, 0xca,
	0xa0, 0x79, 0xfb, 0x0f, 0xf9, 0x1f, 0x3a, 0x1b, 0xe5, 0xb9, 0x75, 0x9e, 0x36, 0x42, 0x7b, 0x6d,
	0x32, 0xb9, 0x24, 0x67, 0xd0, 0x13, 0x56, 0x6b, 0x14, 0xcf, 0xb4, 0x6c, 0x95, 0xc8, 0x34, 0x8a,
	0xc8, 0x1b, 0xe8, 0x3e, 0x43, 0x82, 0x86, 0xed, 0x40, 0x6c, 0x4f, 0xc7, 0x93, 0xf1, 0xf5, 0x75,
	0xa1, 0xa2, 0x23, 0xac, 0x31, 0x28, 0x3c, 0xf3, 0x2a, 0x45, 0x9b, 0x7b, 0xba, 0x13, 0x88, 0xfa,
	0x55, 0x14, 0x45, 0xa4, 0x0f, 0xfb, 0xa5, 0x17, 0x4c, 0x58, 0x87, 0x4c, 0xe6, 0xe9, 0x3c, 0xa3,
	0xcd, 0xe0, 0x4a, 0x23, 0xe6, 0x3a, 0xc3, 0x82, 0x88, 0xb9, 0xd2, 0x4c, 0xc5, 0xcc, 0xdf, 0x39,
	0xeb, 0x7d, 0xe1, 0x1b, 0x6c, 0x12, 0x27, 0xb0, 0xb7, 0xcc, 0x60, 0xbd, 0x65, 0x6b, 0xb9, 0xe5,
	0x31, 0xec, 0xae, 0x6a, 0x61, 0xc5, 0xdd, 0xf5, 0x8a, 0x47, 0xd0, 0xd4, 0x36, 0x61, 0x1a, 0x1f,
	0x50, 0xd3, 0x76, 0xe0, 0xeb, 0xca, 0xc4, 0x96, 0xfc, 0x07, 0x9d, 0x72, 0xe1, 0x20, 0x5b, 0x59,
	0x93, 0xd1, 0x7f, 0x43, 0x5b, 0x75, 0x3c, 0x21, 0xe7, 0xd0, 0x2d, 0x5f, 0x42, 0xc6, 0x36, 0xdf,
	0x50, 0x27, 0xf4, 0x37, 0xb2, 0x3b, 0xee, 0x24, 0x79, 0x05, 0xed, 0x15, 0x53, 0x4e, 0xdf, 0x0b,
	0xd5, 0x1d, 0x89, 0x9e, 0x2b, 0x8d, 0x92, 0x8c, 0x60, 0x7f, 0x05, 0x18, 0x9e, 0x62, 0x36, 0xe7,
	0x02, 0xe9, 0x7e, 0x80, 0x0e, 0xbf, 0x94, 0x72, 0xbf, 0x39, 0x2b, 0x73, 0x81, 0xee, 0xab, 0x9a,
	0x39, 0xee, 0x16, 0x45, 0xe4, 0xab, 0x96, 0x7c, 0xae, 0x2d, 0x97, 0x4c, 0xa2, 0xe6, 0x0b, 0x4a,
	0xca, 0xc8, 0x0b, 0x33, 0x23, 0x72, 0x02, 0x9d, 0x54, 0x99, 0x27, 0xca, 0xbb, 0xa1, 0x5e, 0x19,
	0x91, 0x23, 0x00, 0xc7, 0x3d, 0x32, 0xad, 0x52, 0xe5, 0x69, 0xaf, 0xf4, 0x61, 0x34, 0x8d, 0xc8,
	0x5b, 0xa0, 0x0e, 0x85, 0x75, 0x32, 0x64, 0x38, 0xcb, 0xe3, 0x18, 0x1d, 0xca, 0x90, 0x16, 0x3d,
	0x58, 0x62, 0x51, 0x44, 0x8e, 0x01, 0x96, 0x98, 0xf7, 0x9a, 0x1e, 0x96, 0xd7, 0x5e, 0x86, 0x6b,
	0xf7, 0x60, 0xcb, 0x61, 0xa2, 0xac, 0xa1, 0x47, 0x4b, 0xdb, 0x4f, 0xa1, 0xe3, 0xf0, 0x7b, 0x8e,
	0xd9, 0x3a, 0x75, 0xfa, 0x24, 0x75, 0xf2, 0x80, 0x4e, 0xc5, 0x0b, 0x26, 0xd0, 0x79, 0x15, 0x2b,
	0xc1, 0x3d, 0xd2, 0xe3, 0xf5, 0xcf, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xef, 0x3c, 0xf1,
	0xfa, 0x03, 0x00, 0x00,
}
