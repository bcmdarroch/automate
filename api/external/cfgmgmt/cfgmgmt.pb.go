// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cfgmgmt/cfgmgmt.proto

package cfgmgmt

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/api/external/cfgmgmt/request"
	response "github.com/chef/automate/api/external/cfgmgmt/response"
	query "github.com/chef/automate/api/external/common/query"
	version "github.com/chef/automate/api/external/common/version"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() { proto.RegisterFile("api/external/cfgmgmt/cfgmgmt.proto", fileDescriptor_ee30e63cf8458da6) }

var fileDescriptor_ee30e63cf8458da6 = []byte{
	// 922 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0xdc, 0x44,
	0x14, 0xd7, 0x56, 0x28, 0x54, 0x83, 0xd2, 0x86, 0x47, 0x9b, 0x7a, 0xbd, 0x29, 0x08, 0xf3, 0x27,
	0xa1, 0xb0, 0x1e, 0x4a, 0x39, 0xed, 0x89, 0x12, 0x89, 0x08, 0x89, 0x7f, 0x5a, 0xa4, 0x1e, 0x7a,
	0x59, 0x66, 0xbd, 0xb3, 0xce, 0xa8, 0xeb, 0x19, 0x67, 0x66, 0x1c, 0x08, 0x55, 0x2f, 0x3e, 0x06,
	0xc1, 0x81, 0x7e, 0x02, 0xee, 0x5c, 0x7d, 0xe1, 0xc2, 0x85, 0x6f, 0xd0, 0x2f, 0xc0, 0x01, 0xbe,
	0x02, 0x88, 0x1b, 0xf2, 0x78, 0xec, 0xb5, 0xb7, 0x59, 0xe3, 0x92, 0x9c, 0xac, 0xdd, 0xf9, 0xbd,
	0xdf, 0xbc, 0xdf, 0xfc, 0xde, 0x7b, 0x7a, 0xc8, 0x23, 0x31, 0xc3, 0xf4, 0x1b, 0x4d, 0x25, 0x27,
	0x0b, 0x1c, 0xcc, 0xc3, 0x28, 0x8c, 0x74, 0xf9, 0xf5, 0x63, 0x29, 0xb4, 0x80, 0x7e, 0x70, 0x48,
	0xe7, 0x3e, 0x49, 0xb4, 0x88, 0x88, 0xa6, 0x3e, 0x89, 0x99, 0x6f, 0x01, 0xee, 0xad, 0x33, 0xc3,
	0x25, 0x3d, 0x4a, 0xa8, 0xd2, 0x98, 0x04, 0x9a, 0x09, 0xae, 0x0a, 0x1a, 0x77, 0xaf, 0x15, 0xcb,
	0xc5, 0x8c, 0x76, 0x43, 0x2a, 0x4d, 0x74, 0x89, 0xdc, 0x6d, 0x45, 0x4a, 0x21, 0xac, 0x06, 0xf7,
	0xed, 0x35, 0x40, 0x15, 0x0b, 0xae, 0xe8, 0x4a, 0xa6, 0x6f, 0xb5, 0x83, 0xeb, 0xa9, 0xfe, 0x07,
	0xb4, 0x9e, 0xeb, 0xca, 0x5b, 0x89, 0x28, 0x12, 0x1c, 0x1f, 0x25, 0x54, 0x9e, 0xe0, 0x98, 0x48,
	0x12, 0x51, 0x4d, 0xe5, 0x1a, 0xda, 0x02, 0x7b, 0x4c, 0xa5, 0x62, 0xcb, 0xaf, 0x85, 0xee, 0x84,
	0x42, 0x84, 0x0b, 0x8a, 0xcd, 0xaf, 0x69, 0x32, 0xc7, 0x4a, 0xcb, 0x24, 0xd0, 0x2b, 0xa7, 0x39,
	0x1f, 0xe1, 0x5c, 0x68, 0x52, 0x17, 0xfa, 0x8e, 0xf9, 0x04, 0xc3, 0x90, 0xf2, 0xa1, 0xfa, 0x9a,
	0x84, 0x21, 0x95, 0x58, 0xc4, 0x06, 0x71, 0x06, 0xfa, 0x83, 0x40, 0x44, 0xb1, 0xe0, 0x94, 0x6b,
	0x85, 0xcb, 0x6a, 0x18, 0x86, 0x32, 0x0e, 0x70, 0x8d, 0x26, 0x16, 0x0b, 0x16, 0x9c, 0xac, 0xb9,
	0xef, 0x59, 0x18, 0x18, 0x89, 0x9e, 0x66, 0x78, 0xef, 0xa7, 0x97, 0x10, 0xda, 0x17, 0x7c, 0xce,
	0xc2, 0x4f, 0xc3, 0x48, 0xc3, 0xcf, 0x3d, 0x74, 0xf9, 0x80, 0xea, 0xcf, 0x72, 0x47, 0x60, 0xcf,
	0x5f, 0x5b, 0xa8, 0xbe, 0x2d, 0x09, 0xdf, 0x20, 0x5d, 0xd7, 0x2f, 0x9e, 0xc5, 0x2f, 0x1f, 0xcd,
	0xff, 0x84, 0x29, 0x7d, 0x8f, 0x2c, 0x12, 0xea, 0xdd, 0x4f, 0x33, 0xe7, 0x2a, 0xda, 0xb4, 0x81,
	0x23, 0x63, 0x76, 0x9a, 0x39, 0x1b, 0xf0, 0x9c, 0xa4, 0x64, 0x76, 0x9a, 0x39, 0x9b, 0xe8, 0x05,
	0xc6, 0xe7, 0x92, 0x14, 0x47, 0xa7, 0x99, 0x03, 0xb0, 0x55, 0xfb, 0x63, 0xb4, 0x60, 0x4a, 0xa7,
	0x4f, 0xfe, 0x78, 0x7c, 0x69, 0x0b, 0xae, 0x54, 0x15, 0x61, 0x8e, 0xe0, 0x49, 0x0f, 0x3d, 0x7f,
	0x40, 0xf5, 0x38, 0xe1, 0x0a, 0x76, 0x3b, 0x64, 0x9b, 0x03, 0x5b, 0x93, 0x4d, 0x7b, 0x69, 0xe6,
	0xbc, 0x8c, 0x76, 0x1a, 0xd9, 0x8e, 0x1e, 0xe6, 0x9f, 0x09, 0x9b, 0x3d, 0x1a, 0xc9, 0x84, 0x37,
	0x93, 0xbf, 0x81, 0xae, 0xd7, 0x73, 0xad, 0x90, 0x2d, 0x32, 0x5e, 0x81, 0x9b, 0x4d, 0x19, 0xb8,
	0x8a, 0xc2, 0x39, 0x3f, 0xfc, 0xd9, 0x43, 0x57, 0x4a, 0x13, 0xf6, 0x45, 0xc2, 0xb5, 0x02, 0xbf,
	0xab, 0x15, 0x05, 0xde, 0xc5, 0xad, 0xf8, 0xa2, 0x99, 0xea, 0x01, 0x9e, 0x4c, 0x33, 0x67, 0x80,
	0xfa, 0xa5, 0x6e, 0xd3, 0x67, 0x26, 0xf1, 0x49, 0x60, 0x00, 0xff, 0xcf, 0xb1, 0x1d, 0x70, 0x2b,
	0xa9, 0x86, 0x12, 0xd7, 0x28, 0xe1, 0xf7, 0x1e, 0xda, 0xb4, 0xee, 0x59, 0x99, 0xc3, 0x8e, 0x1e,
	0x5a, 0x95, 0x7e, 0x17, 0x95, 0x4b, 0xbc, 0x17, 0xa7, 0x99, 0xe3, 0x22, 0xa7, 0x29, 0x52, 0x26,
	0xfc, 0x5c, 0x1a, 0x07, 0xd0, 0x5f, 0xd1, 0xb8, 0x64, 0x84, 0xbf, 0x7b, 0x08, 0x59, 0x2b, 0xc7,
	0x09, 0x87, 0x5b, 0x1d, 0x6d, 0x1c, 0x27, 0xdc, 0xdd, 0xed, 0x28, 0xce, 0x7b, 0x9c, 0xd7, 0xec,
	0x2e, 0x7a, 0xa3, 0xad, 0x66, 0x47, 0x0f, 0xf3, 0xa4, 0xd8, 0xec, 0x51, 0xd7, 0xe2, 0x7d, 0x11,
	0xae, 0xd6, 0x8f, 0x42, 0x5a, 0x88, 0xdd, 0x83, 0x37, 0x5b, 0x6b, 0x17, 0x97, 0xf7, 0xc0, 0x2f,
	0x45, 0x11, 0x7f, 0x99, 0x84, 0x21, 0x55, 0x66, 0xe2, 0x9c, 0xad, 0xde, 0x8c, 0x62, 0xdf, 0x8c,
	0x6d, 0x7f, 0x09, 0x6e, 0x6d, 0xd2, 0xaf, 0x2e, 0x6c, 0xa2, 0x6c, 0xc3, 0xb5, 0xa5, 0x77, 0xb5,
	0x4c, 0x7f, 0xeb, 0xa1, 0xad, 0x03, 0xaa, 0x3f, 0x97, 0x21, 0xe1, 0xec, 0xdb, 0x62, 0x60, 0xc2,
	0xbb, 0x1d, 0xcc, 0x6b, 0x44, 0xb4, 0x8a, 0x98, 0x5e, 0x98, 0x08, 0x07, 0xb6, 0x2b, 0x11, 0xa2,
	0x91, 0xf1, 0xaf, 0xd6, 0x03, 0x91, 0xc8, 0x80, 0x7e, 0x74, 0x34, 0xe3, 0xdd, 0x06, 0x49, 0x0d,
	0xdf, 0x2a, 0x81, 0x5c, 0x98, 0x84, 0x1b, 0x70, 0x7d, 0xe9, 0x83, 0xb9, 0x7d, 0x32, 0x37, 0xe9,
	0xfe, 0x53, 0x8c, 0x88, 0xbb, 0x5a, 0x4b, 0x36, 0x4d, 0x34, 0xed, 0x36, 0xe6, 0xf3, 0x16, 0x72,
	0x6f, 0x77, 0x1d, 0x81, 0x15, 0xb9, 0xf7, 0x5d, 0xde, 0x49, 0xaf, 0xa1, 0x57, 0xd7, 0x75, 0x12,
	0x29, 0x91, 0xe7, 0xec, 0xa2, 0xd7, 0xc1, 0x5b, 0xdb, 0x45, 0xd5, 0x1d, 0xf0, 0x57, 0x31, 0x3b,
	0xee, 0x15, 0xdb, 0x09, 0xbc, 0xbf, 0xbe, 0x7b, 0xca, 0x05, 0xc6, 0x42, 0x3f, 0xe6, 0x73, 0x31,
	0x2e, 0x9e, 0xc2, 0x1d, 0x3e, 0x53, 0x94, 0x97, 0xf6, 0x7e, 0xbc, 0x7b, 0x19, 0x6d, 0x1c, 0xb2,
	0xd9, 0x8c, 0xf2, 0x34, 0x73, 0xb6, 0xd1, 0x35, 0x45, 0xe5, 0x31, 0x0b, 0xe8, 0x84, 0xf1, 0xb9,
	0x18, 0xd9, 0x98, 0x86, 0x7e, 0x07, 0x6d, 0xab, 0x13, 0xa5, 0x69, 0x34, 0xb2, 0xd0, 0x12, 0x75,
	0x9a, 0x39, 0x03, 0xe8, 0x37, 0xcf, 0xec, 0x65, 0xd5, 0x53, 0x00, 0x6c, 0x55, 0x4f, 0x61, 0xc3,
	0xe0, 0xfb, 0x4b, 0x08, 0x0e, 0xa8, 0xfe, 0xc2, 0xec, 0x2d, 0xfb, 0x42, 0x3c, 0x98, 0x0a, 0xf1,
	0x40, 0xc1, 0xed, 0x0e, 0xce, 0x17, 0x31, 0x63, 0x7a, 0xcc, 0x72, 0x26, 0xf7, 0x4e, 0x97, 0x1a,
	0x58, 0xb9, 0xc7, 0xfb, 0x21, 0xaf, 0x82, 0x9b, 0x68, 0xb0, 0x52, 0x05, 0xd2, 0x92, 0x3e, 0x35,
	0x45, 0x07, 0xa8, 0xdf, 0xf0, 0xbf, 0x0e, 0x6c, 0xa9, 0xf9, 0xfa, 0x28, 0x2d, 0x56, 0xb4, 0x49,
	0x19, 0x89, 0x1b, 0x1c, 0x1f, 0xe2, 0xfb, 0xc3, 0x90, 0xe9, 0xc3, 0x64, 0x9a, 0x3b, 0x87, 0x73,
	0x45, 0xd5, 0xb2, 0x87, 0xcf, 0x5a, 0x97, 0xa7, 0x1b, 0xa6, 0x59, 0xef, 0xfc, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x23, 0xd5, 0xd6, 0xa4, 0x8d, 0x0c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConfigMgmtClient is the client API for ConfigMgmt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigMgmtClient interface {
	//
	//GetNodes
	//
	//Returns a list of infra nodes that have checked in to Automate.
	//Adding a filter makes a list of all nodes that meet the filter criteria.
	//Filters for the same field are ORd together, while filters across different fields are ANDed together.
	//Supports pagination, filtering (with wildcard support), and sorting.
	//Limited to 10k results.
	//
	//Example:
	//```
	//cfgmgmt/nodes?pagination.page=1&pagination.size=100&sorting.field=name&sorting.order=ASC&filter=name:mySO*&filter=platform:ubun*
	//```
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetNodes(ctx context.Context, in *request.Nodes, opts ...grpc.CallOption) (*_struct.ListValue, error)
	//
	//GetRuns
	//
	//Returns a list of run metadata (id, start and end time, and status) for the provided node ID.
	//Supports pagination.
	//Accepts a `start` parameter to denote start date for the list and a filter of type `status`.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetRuns(ctx context.Context, in *request.Runs, opts ...grpc.CallOption) (*_struct.ListValue, error)
	//
	//GetNodesCounts
	//
	//Returns totals for failed, success, missing, and overall total infra nodes that have reported into Automate.
	//Supports filtering.
	//
	//Example:
	//```
	//cfgmgmt/stats/node_counts?filter=name:mySO*&filter=platform:ubun*
	//```
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetNodesCounts(ctx context.Context, in *request.NodesCounts, opts ...grpc.CallOption) (*response.NodesCounts, error)
	//
	//GetRunsCounts
	//
	//Returns totals for failed and successful runs given a `node_id`.
	//
	//Example:
	//```
	//cfgmgmt/stats/run_counts?node_id=821fff07-abc9-4160-96b1-83d68ae5cfdd&start=2019-11-02
	//```
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetRunsCounts(ctx context.Context, in *request.RunsCounts, opts ...grpc.CallOption) (*response.RunsCounts, error)
	//
	//GetNodeRun
	//
	//Returns the infra run report for the provided node ID and run ID.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:get
	//```
	GetNodeRun(ctx context.Context, in *request.NodeRun, opts ...grpc.CallOption) (*response.Run, error)
	//
	//GetSuggestions
	//
	//Returns possible filter values given a valid `type` parameter. All values returned until two or more
	//characters are provided for the `text` parameter.
	//Supports wildcard (* and ?).
	//
	//Example:
	//```
	//cfgmgmt/suggestions?type=environment&text=_d
	//```
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetSuggestions(ctx context.Context, in *query.Suggestion, opts ...grpc.CallOption) (*_struct.ListValue, error)
	//
	//GetOrganizations
	//
	//Returns a list of all organizations associated with nodes that have checked in to Automate.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetOrganizations(ctx context.Context, in *request.Organizations, opts ...grpc.CallOption) (*_struct.ListValue, error)
	//
	//GetSourceFqdns
	//
	//Returns a list of all chef servers associated with nodes that have checked in to Automate.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetSourceFqdns(ctx context.Context, in *request.SourceFqdns, opts ...grpc.CallOption) (*_struct.ListValue, error)
	//
	//GetAttributes
	//
	//Returns the latest reported attributes for the provided node ID.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:get
	//```
	GetAttributes(ctx context.Context, in *request.Node, opts ...grpc.CallOption) (*response.NodeAttribute, error)
	GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error)
	//
	//GetPolicyCookbooks
	//
	//Returns Policy Names with a list of cookbook names and associated policy identifiers based on a policy revision ID.
	//Policy revision ids are sent with an infra run report and identifies which instance of a policy the node used for this run.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetPolicyCookbooks(ctx context.Context, in *request.PolicyRevision, opts ...grpc.CallOption) (*response.PolicyCookbooks, error)
}

type configMgmtClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigMgmtClient(cc grpc.ClientConnInterface) ConfigMgmtClient {
	return &configMgmtClient{cc}
}

func (c *configMgmtClient) GetNodes(ctx context.Context, in *request.Nodes, opts ...grpc.CallOption) (*_struct.ListValue, error) {
	out := new(_struct.ListValue)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMgmtClient) GetRuns(ctx context.Context, in *request.Runs, opts ...grpc.CallOption) (*_struct.ListValue, error) {
	out := new(_struct.ListValue)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetRuns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMgmtClient) GetNodesCounts(ctx context.Context, in *request.NodesCounts, opts ...grpc.CallOption) (*response.NodesCounts, error) {
	out := new(response.NodesCounts)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodesCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMgmtClient) GetRunsCounts(ctx context.Context, in *request.RunsCounts, opts ...grpc.CallOption) (*response.RunsCounts, error) {
	out := new(response.RunsCounts)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetRunsCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMgmtClient) GetNodeRun(ctx context.Context, in *request.NodeRun, opts ...grpc.CallOption) (*response.Run, error) {
	out := new(response.Run)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodeRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMgmtClient) GetSuggestions(ctx context.Context, in *query.Suggestion, opts ...grpc.CallOption) (*_struct.ListValue, error) {
	out := new(_struct.ListValue)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetSuggestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMgmtClient) GetOrganizations(ctx context.Context, in *request.Organizations, opts ...grpc.CallOption) (*_struct.ListValue, error) {
	out := new(_struct.ListValue)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetOrganizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMgmtClient) GetSourceFqdns(ctx context.Context, in *request.SourceFqdns, opts ...grpc.CallOption) (*_struct.ListValue, error) {
	out := new(_struct.ListValue)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetSourceFqdns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMgmtClient) GetAttributes(ctx context.Context, in *request.Node, opts ...grpc.CallOption) (*response.NodeAttribute, error) {
	out := new(response.NodeAttribute)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMgmtClient) GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error) {
	out := new(version.VersionInfo)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMgmtClient) GetPolicyCookbooks(ctx context.Context, in *request.PolicyRevision, opts ...grpc.CallOption) (*response.PolicyCookbooks, error) {
	out := new(response.PolicyCookbooks)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cfgmgmt.ConfigMgmt/GetPolicyCookbooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigMgmtServer is the server API for ConfigMgmt service.
type ConfigMgmtServer interface {
	//
	//GetNodes
	//
	//Returns a list of infra nodes that have checked in to Automate.
	//Adding a filter makes a list of all nodes that meet the filter criteria.
	//Filters for the same field are ORd together, while filters across different fields are ANDed together.
	//Supports pagination, filtering (with wildcard support), and sorting.
	//Limited to 10k results.
	//
	//Example:
	//```
	//cfgmgmt/nodes?pagination.page=1&pagination.size=100&sorting.field=name&sorting.order=ASC&filter=name:mySO*&filter=platform:ubun*
	//```
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetNodes(context.Context, *request.Nodes) (*_struct.ListValue, error)
	//
	//GetRuns
	//
	//Returns a list of run metadata (id, start and end time, and status) for the provided node ID.
	//Supports pagination.
	//Accepts a `start` parameter to denote start date for the list and a filter of type `status`.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetRuns(context.Context, *request.Runs) (*_struct.ListValue, error)
	//
	//GetNodesCounts
	//
	//Returns totals for failed, success, missing, and overall total infra nodes that have reported into Automate.
	//Supports filtering.
	//
	//Example:
	//```
	//cfgmgmt/stats/node_counts?filter=name:mySO*&filter=platform:ubun*
	//```
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetNodesCounts(context.Context, *request.NodesCounts) (*response.NodesCounts, error)
	//
	//GetRunsCounts
	//
	//Returns totals for failed and successful runs given a `node_id`.
	//
	//Example:
	//```
	//cfgmgmt/stats/run_counts?node_id=821fff07-abc9-4160-96b1-83d68ae5cfdd&start=2019-11-02
	//```
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetRunsCounts(context.Context, *request.RunsCounts) (*response.RunsCounts, error)
	//
	//GetNodeRun
	//
	//Returns the infra run report for the provided node ID and run ID.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:get
	//```
	GetNodeRun(context.Context, *request.NodeRun) (*response.Run, error)
	//
	//GetSuggestions
	//
	//Returns possible filter values given a valid `type` parameter. All values returned until two or more
	//characters are provided for the `text` parameter.
	//Supports wildcard (* and ?).
	//
	//Example:
	//```
	//cfgmgmt/suggestions?type=environment&text=_d
	//```
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetSuggestions(context.Context, *query.Suggestion) (*_struct.ListValue, error)
	//
	//GetOrganizations
	//
	//Returns a list of all organizations associated with nodes that have checked in to Automate.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetOrganizations(context.Context, *request.Organizations) (*_struct.ListValue, error)
	//
	//GetSourceFqdns
	//
	//Returns a list of all chef servers associated with nodes that have checked in to Automate.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetSourceFqdns(context.Context, *request.SourceFqdns) (*_struct.ListValue, error)
	//
	//GetAttributes
	//
	//Returns the latest reported attributes for the provided node ID.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:get
	//```
	GetAttributes(context.Context, *request.Node) (*response.NodeAttribute, error)
	GetVersion(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
	//
	//GetPolicyCookbooks
	//
	//Returns Policy Names with a list of cookbook names and associated policy identifiers based on a policy revision ID.
	//Policy revision ids are sent with an infra run report and identifies which instance of a policy the node used for this run.
	//
	//Authorization Action:
	//
	//```
	//infra:nodes:list
	//```
	GetPolicyCookbooks(context.Context, *request.PolicyRevision) (*response.PolicyCookbooks, error)
}

// UnimplementedConfigMgmtServer can be embedded to have forward compatible implementations.
type UnimplementedConfigMgmtServer struct {
}

func (*UnimplementedConfigMgmtServer) GetNodes(ctx context.Context, req *request.Nodes) (*_struct.ListValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodes not implemented")
}
func (*UnimplementedConfigMgmtServer) GetRuns(ctx context.Context, req *request.Runs) (*_struct.ListValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRuns not implemented")
}
func (*UnimplementedConfigMgmtServer) GetNodesCounts(ctx context.Context, req *request.NodesCounts) (*response.NodesCounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodesCounts not implemented")
}
func (*UnimplementedConfigMgmtServer) GetRunsCounts(ctx context.Context, req *request.RunsCounts) (*response.RunsCounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRunsCounts not implemented")
}
func (*UnimplementedConfigMgmtServer) GetNodeRun(ctx context.Context, req *request.NodeRun) (*response.Run, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeRun not implemented")
}
func (*UnimplementedConfigMgmtServer) GetSuggestions(ctx context.Context, req *query.Suggestion) (*_struct.ListValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuggestions not implemented")
}
func (*UnimplementedConfigMgmtServer) GetOrganizations(ctx context.Context, req *request.Organizations) (*_struct.ListValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizations not implemented")
}
func (*UnimplementedConfigMgmtServer) GetSourceFqdns(ctx context.Context, req *request.SourceFqdns) (*_struct.ListValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSourceFqdns not implemented")
}
func (*UnimplementedConfigMgmtServer) GetAttributes(ctx context.Context, req *request.Node) (*response.NodeAttribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttributes not implemented")
}
func (*UnimplementedConfigMgmtServer) GetVersion(ctx context.Context, req *version.VersionInfoRequest) (*version.VersionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (*UnimplementedConfigMgmtServer) GetPolicyCookbooks(ctx context.Context, req *request.PolicyRevision) (*response.PolicyCookbooks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyCookbooks not implemented")
}

func RegisterConfigMgmtServer(s *grpc.Server, srv ConfigMgmtServer) {
	s.RegisterService(&_ConfigMgmt_serviceDesc, srv)
}

func _ConfigMgmt_GetNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Nodes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetNodes(ctx, req.(*request.Nodes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMgmt_GetRuns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Runs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetRuns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetRuns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetRuns(ctx, req.(*request.Runs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMgmt_GetNodesCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.NodesCounts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetNodesCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodesCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetNodesCounts(ctx, req.(*request.NodesCounts))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMgmt_GetRunsCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.RunsCounts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetRunsCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetRunsCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetRunsCounts(ctx, req.(*request.RunsCounts))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMgmt_GetNodeRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.NodeRun)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetNodeRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodeRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetNodeRun(ctx, req.(*request.NodeRun))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMgmt_GetSuggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.Suggestion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetSuggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetSuggestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetSuggestions(ctx, req.(*query.Suggestion))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMgmt_GetOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Organizations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetOrganizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetOrganizations(ctx, req.(*request.Organizations))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMgmt_GetSourceFqdns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SourceFqdns)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetSourceFqdns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetSourceFqdns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetSourceFqdns(ctx, req.(*request.SourceFqdns))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMgmt_GetAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetAttributes(ctx, req.(*request.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMgmt_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.VersionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetVersion(ctx, req.(*version.VersionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMgmt_GetPolicyCookbooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.PolicyRevision)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMgmtServer).GetPolicyCookbooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cfgmgmt.ConfigMgmt/GetPolicyCookbooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMgmtServer).GetPolicyCookbooks(ctx, req.(*request.PolicyRevision))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigMgmt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.cfgmgmt.ConfigMgmt",
	HandlerType: (*ConfigMgmtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodes",
			Handler:    _ConfigMgmt_GetNodes_Handler,
		},
		{
			MethodName: "GetRuns",
			Handler:    _ConfigMgmt_GetRuns_Handler,
		},
		{
			MethodName: "GetNodesCounts",
			Handler:    _ConfigMgmt_GetNodesCounts_Handler,
		},
		{
			MethodName: "GetRunsCounts",
			Handler:    _ConfigMgmt_GetRunsCounts_Handler,
		},
		{
			MethodName: "GetNodeRun",
			Handler:    _ConfigMgmt_GetNodeRun_Handler,
		},
		{
			MethodName: "GetSuggestions",
			Handler:    _ConfigMgmt_GetSuggestions_Handler,
		},
		{
			MethodName: "GetOrganizations",
			Handler:    _ConfigMgmt_GetOrganizations_Handler,
		},
		{
			MethodName: "GetSourceFqdns",
			Handler:    _ConfigMgmt_GetSourceFqdns_Handler,
		},
		{
			MethodName: "GetAttributes",
			Handler:    _ConfigMgmt_GetAttributes_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ConfigMgmt_GetVersion_Handler,
		},
		{
			MethodName: "GetPolicyCookbooks",
			Handler:    _ConfigMgmt_GetPolicyCookbooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/external/cfgmgmt/cfgmgmt.proto",
}
