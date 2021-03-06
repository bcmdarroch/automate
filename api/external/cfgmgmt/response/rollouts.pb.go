// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cfgmgmt/response/rollouts.proto

package response

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SCMType int32

const (
	SCMType_UNKNOWN_SCM SCMType = 0
	SCMType_GIT         SCMType = 1
)

var SCMType_name = map[int32]string{
	0: "UNKNOWN_SCM",
	1: "GIT",
}

var SCMType_value = map[string]int32{
	"UNKNOWN_SCM": 0,
	"GIT":         1,
}

func (x SCMType) String() string {
	return proto.EnumName(SCMType_name, int32(x))
}

func (SCMType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3bb6ceeab2493f2, []int{0}
}

type SCMWebType int32

const (
	SCMWebType_UNKNOWN_SCM_WEB SCMWebType = 0
	SCMWebType_GITHUB          SCMWebType = 1
)

var SCMWebType_name = map[int32]string{
	0: "UNKNOWN_SCM_WEB",
	1: "GITHUB",
}

var SCMWebType_value = map[string]int32{
	"UNKNOWN_SCM_WEB": 0,
	"GITHUB":          1,
}

func (x SCMWebType) String() string {
	return proto.EnumName(SCMWebType_name, int32(x))
}

func (SCMWebType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3bb6ceeab2493f2, []int{1}
}

//
//A "Rollout" represents the process of distributing Chef Infra code (with
//Policyfiles) to a set of nodes. It's used to track which nodes have run the
//latest version of the Chef Infra code assigned to them and also provide the
//user insights about the code by aggregating Chef Client run results
//according to the version of Chef Infra code applied. Metadata about the code
//is stored in order to provide the user with convenient references back to
//systems they already use (such as SCM and Ci/CD systems) to manage their code.
//
//Nodes are segmented by a triple of policy name, policy group, and policy domain URL:
// policy name generally describes what kind of system it is, e.g., a database server
// policy group generally describes where the system fits in the user's code
//lifecycle, e.g., "QA" or "production"
// policy domain URL identifies the system that distributes the Chef Infra code
//and is the owner of the namespaces for policy name and group. E.g., a Chef
//Server URL with the `/organizations/:orgname` part.
//
//There is one (or zero) revision(s) of the Chef Infra code applied to any
//segment at a time. Rollouts track the changes to which revision of the code is
//applied to the node segments over time.
//
type Rollout struct {
	// The name of the policy, i.e., the `name` attribute in the Policyfile
	PolicyName string `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// The group of nodes which are targeted by the rollout. In the Chef Server
	// case, this is the policy_group to which the user is pushing the policy.
	PolicyNodeGroup string `protobuf:"bytes,2,opt,name=policy_node_group,json=policyNodeGroup,proto3" json:"policy_node_group,omitempty"`
	// The revision_id of the compiled policy being rolled out
	PolicyRevisionId string `protobuf:"bytes,3,opt,name=policy_revision_id,json=policyRevisionId,proto3" json:"policy_revision_id,omitempty"`
	// In the Chef Server case, the policy domain URL is the Chef Server URL
	// with the `/organizations/:orgname` portion of the URL path included. In
	// general, this can be a URL for any content storage/distribution service,
	// as long as the combination of policy_name and policy_node_group is unique
	// on that system.
	//
	// The set of nodes configured to fetch policy content from the
	// policy_domain_url and configured with the same policy_name and
	// policy_node_group form the target set of nodes for a rollout and are
	// expected to apply the policy revision described by the rollout.
	PolicyDomainUrl string `protobuf:"bytes,4,opt,name=policy_domain_url,json=policyDomainUrl,proto3" json:"policy_domain_url,omitempty"`
	// The source control system used with the policyfile
	ScmType SCMType `protobuf:"varint,5,opt,name=scm_type,json=scmType,proto3,enum=chef.automate.api.cfgmgmt.response.SCMType" json:"scm_type,omitempty"`
	// The software/service used to host the source code repository
	ScmWebType SCMWebType `protobuf:"varint,6,opt,name=scm_web_type,json=scmWebType,proto3,enum=chef.automate.api.cfgmgmt.response.SCMWebType" json:"scm_web_type,omitempty"`
	// The URL used to obtain a copy of the source code repository
	PolicyScmUrl string `protobuf:"bytes,7,opt,name=policy_scm_url,json=policyScmUrl,proto3" json:"policy_scm_url,omitempty"`
	// The URL used to view the source code repository via the web
	PolicyScmWebUrl string `protobuf:"bytes,8,opt,name=policy_scm_web_url,json=policyScmWebUrl,proto3" json:"policy_scm_web_url,omitempty"`
	// The source control system's identifier for the repository version. This
	// should be the version where the policy's lockfile was committed.
	PolicyScmCommit string `protobuf:"bytes,9,opt,name=policy_scm_commit,json=policyScmCommit,proto3" json:"policy_scm_commit,omitempty"`
	// A free-form description of the rollout, as given by the user.
	Description string `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	// If the rollout was initiated via Ci/CD or similar system, the web URL
	// for the job that initiated the rollout.
	CiJobUrl string `protobuf:"bytes,11,opt,name=ci_job_url,json=ciJobUrl,proto3" json:"ci_job_url,omitempty"`
	// If the rollout was initiated by Ci/CD or similar system, the id of the job
	// that initiated the rollout. Should include the Ci system's nickname or
	// other identifying information users would need to associate the job ID to
	// the Ci/CD system.
	CiJobId              string   `protobuf:"bytes,12,opt,name=ci_job_id,json=ciJobId,proto3" json:"ci_job_id,omitempty"`
	Id                   string   `protobuf:"bytes,13,opt,name=id,proto3" json:"id,omitempty"`
	StartTime            string   `protobuf:"bytes,14,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string   `protobuf:"bytes,15,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rollout) Reset()         { *m = Rollout{} }
func (m *Rollout) String() string { return proto.CompactTextString(m) }
func (*Rollout) ProtoMessage()    {}
func (*Rollout) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3bb6ceeab2493f2, []int{0}
}

func (m *Rollout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rollout.Unmarshal(m, b)
}
func (m *Rollout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rollout.Marshal(b, m, deterministic)
}
func (m *Rollout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rollout.Merge(m, src)
}
func (m *Rollout) XXX_Size() int {
	return xxx_messageInfo_Rollout.Size(m)
}
func (m *Rollout) XXX_DiscardUnknown() {
	xxx_messageInfo_Rollout.DiscardUnknown(m)
}

var xxx_messageInfo_Rollout proto.InternalMessageInfo

func (m *Rollout) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *Rollout) GetPolicyNodeGroup() string {
	if m != nil {
		return m.PolicyNodeGroup
	}
	return ""
}

func (m *Rollout) GetPolicyRevisionId() string {
	if m != nil {
		return m.PolicyRevisionId
	}
	return ""
}

func (m *Rollout) GetPolicyDomainUrl() string {
	if m != nil {
		return m.PolicyDomainUrl
	}
	return ""
}

func (m *Rollout) GetScmType() SCMType {
	if m != nil {
		return m.ScmType
	}
	return SCMType_UNKNOWN_SCM
}

func (m *Rollout) GetScmWebType() SCMWebType {
	if m != nil {
		return m.ScmWebType
	}
	return SCMWebType_UNKNOWN_SCM_WEB
}

func (m *Rollout) GetPolicyScmUrl() string {
	if m != nil {
		return m.PolicyScmUrl
	}
	return ""
}

func (m *Rollout) GetPolicyScmWebUrl() string {
	if m != nil {
		return m.PolicyScmWebUrl
	}
	return ""
}

func (m *Rollout) GetPolicyScmCommit() string {
	if m != nil {
		return m.PolicyScmCommit
	}
	return ""
}

func (m *Rollout) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Rollout) GetCiJobUrl() string {
	if m != nil {
		return m.CiJobUrl
	}
	return ""
}

func (m *Rollout) GetCiJobId() string {
	if m != nil {
		return m.CiJobId
	}
	return ""
}

func (m *Rollout) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Rollout) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Rollout) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

type Rollouts struct {
	Rollouts             []*Rollout `protobuf:"bytes,1,rep,name=rollouts,proto3" json:"rollouts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Rollouts) Reset()         { *m = Rollouts{} }
func (m *Rollouts) String() string { return proto.CompactTextString(m) }
func (*Rollouts) ProtoMessage()    {}
func (*Rollouts) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3bb6ceeab2493f2, []int{1}
}

func (m *Rollouts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rollouts.Unmarshal(m, b)
}
func (m *Rollouts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rollouts.Marshal(b, m, deterministic)
}
func (m *Rollouts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rollouts.Merge(m, src)
}
func (m *Rollouts) XXX_Size() int {
	return xxx_messageInfo_Rollouts.Size(m)
}
func (m *Rollouts) XXX_DiscardUnknown() {
	xxx_messageInfo_Rollouts.DiscardUnknown(m)
}

var xxx_messageInfo_Rollouts proto.InternalMessageInfo

func (m *Rollouts) GetRollouts() []*Rollout {
	if m != nil {
		return m.Rollouts
	}
	return nil
}

//
//NodeSegmentsWithRolloutProgress
//
//A Node Segment is the set of Chef Infra nodes with a shared policy_name,
//policy_node_group, and policy_domain_url.
//
//NodeSegmentsWithRolloutProgress lists all of the node segments matching the
//request with information about the progress and status of the code rollouts for each segment.
type NodeSegmentsWithRolloutProgress struct {
	// The NodeSegmentRolloutProgress are sorted by policy group, policy
	// name, then domain URL.
	NodeSegmentRolloutProgress []*NodeSegmentRolloutProgress `protobuf:"bytes,1,rep,name=node_segment_rollout_progress,json=nodeSegmentRolloutProgress,proto3" json:"node_segment_rollout_progress,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                      `json:"-"`
	XXX_unrecognized           []byte                        `json:"-"`
	XXX_sizecache              int32                         `json:"-"`
}

func (m *NodeSegmentsWithRolloutProgress) Reset()         { *m = NodeSegmentsWithRolloutProgress{} }
func (m *NodeSegmentsWithRolloutProgress) String() string { return proto.CompactTextString(m) }
func (*NodeSegmentsWithRolloutProgress) ProtoMessage()    {}
func (*NodeSegmentsWithRolloutProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3bb6ceeab2493f2, []int{2}
}

func (m *NodeSegmentsWithRolloutProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeSegmentsWithRolloutProgress.Unmarshal(m, b)
}
func (m *NodeSegmentsWithRolloutProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeSegmentsWithRolloutProgress.Marshal(b, m, deterministic)
}
func (m *NodeSegmentsWithRolloutProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeSegmentsWithRolloutProgress.Merge(m, src)
}
func (m *NodeSegmentsWithRolloutProgress) XXX_Size() int {
	return xxx_messageInfo_NodeSegmentsWithRolloutProgress.Size(m)
}
func (m *NodeSegmentsWithRolloutProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeSegmentsWithRolloutProgress.DiscardUnknown(m)
}

var xxx_messageInfo_NodeSegmentsWithRolloutProgress proto.InternalMessageInfo

func (m *NodeSegmentsWithRolloutProgress) GetNodeSegmentRolloutProgress() []*NodeSegmentRolloutProgress {
	if m != nil {
		return m.NodeSegmentRolloutProgress
	}
	return nil
}

type NodeSegmentRolloutProgress struct {
	// policy_name, policy_node_group, policy_domain_url make up a "compound
	// id" for the node segment
	PolicyName      string `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyNodeGroup string `protobuf:"bytes,2,opt,name=policy_node_group,json=policyNodeGroup,proto3" json:"policy_node_group,omitempty"`
	PolicyDomainUrl string `protobuf:"bytes,3,opt,name=policy_domain_url,json=policyDomainUrl,proto3" json:"policy_domain_url,omitempty"`
	// total nodes in elasticsearch in the node segment
	TotalNodes             int32                   `protobuf:"varint,4,opt,name=total_nodes,json=totalNodes,proto3" json:"total_nodes,omitempty"`
	CurrentRolloutProgress *CurrentRolloutProgress `protobuf:"bytes,5,opt,name=current_rollout_progress,json=currentRolloutProgress,proto3" json:"current_rollout_progress,omitempty"`
	// This is the last, say 2 or 4 rollouts before the current one (to give a
	// total of 3 or 5)
	PreviousRollouts     []*PastRolloutProgress `protobuf:"bytes,6,rep,name=previous_rollouts,json=previousRollouts,proto3" json:"previous_rollouts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *NodeSegmentRolloutProgress) Reset()         { *m = NodeSegmentRolloutProgress{} }
func (m *NodeSegmentRolloutProgress) String() string { return proto.CompactTextString(m) }
func (*NodeSegmentRolloutProgress) ProtoMessage()    {}
func (*NodeSegmentRolloutProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3bb6ceeab2493f2, []int{3}
}

func (m *NodeSegmentRolloutProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeSegmentRolloutProgress.Unmarshal(m, b)
}
func (m *NodeSegmentRolloutProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeSegmentRolloutProgress.Marshal(b, m, deterministic)
}
func (m *NodeSegmentRolloutProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeSegmentRolloutProgress.Merge(m, src)
}
func (m *NodeSegmentRolloutProgress) XXX_Size() int {
	return xxx_messageInfo_NodeSegmentRolloutProgress.Size(m)
}
func (m *NodeSegmentRolloutProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeSegmentRolloutProgress.DiscardUnknown(m)
}

var xxx_messageInfo_NodeSegmentRolloutProgress proto.InternalMessageInfo

func (m *NodeSegmentRolloutProgress) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *NodeSegmentRolloutProgress) GetPolicyNodeGroup() string {
	if m != nil {
		return m.PolicyNodeGroup
	}
	return ""
}

func (m *NodeSegmentRolloutProgress) GetPolicyDomainUrl() string {
	if m != nil {
		return m.PolicyDomainUrl
	}
	return ""
}

func (m *NodeSegmentRolloutProgress) GetTotalNodes() int32 {
	if m != nil {
		return m.TotalNodes
	}
	return 0
}

func (m *NodeSegmentRolloutProgress) GetCurrentRolloutProgress() *CurrentRolloutProgress {
	if m != nil {
		return m.CurrentRolloutProgress
	}
	return nil
}

func (m *NodeSegmentRolloutProgress) GetPreviousRollouts() []*PastRolloutProgress {
	if m != nil {
		return m.PreviousRollouts
	}
	return nil
}

type CurrentRolloutProgress struct {
	// Rollout is the full rollout object, but we can change this to be a subset only.
	Rollout *Rollout `protobuf:"bytes,1,opt,name=rollout,proto3" json:"rollout,omitempty"`
	// Nodes that have run the code being rolled out thus far
	NodeCount int32 `protobuf:"varint,2,opt,name=node_count,json=nodeCount,proto3" json:"node_count,omitempty"`
	// I'm assuming it's easy to get the status when we get the counts.
	LatestRunSuccessfulCount int32    `protobuf:"varint,3,opt,name=latest_run_successful_count,json=latestRunSuccessfulCount,proto3" json:"latest_run_successful_count,omitempty"`
	LatestRunErroredCount    int32    `protobuf:"varint,4,opt,name=latest_run_errored_count,json=latestRunErroredCount,proto3" json:"latest_run_errored_count,omitempty"`
	SourceLink               string   `protobuf:"bytes,5,opt,name=source_link,json=sourceLink,proto3" json:"source_link,omitempty"`
	BuildLink                string   `protobuf:"bytes,6,opt,name=build_link,json=buildLink,proto3" json:"build_link,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *CurrentRolloutProgress) Reset()         { *m = CurrentRolloutProgress{} }
func (m *CurrentRolloutProgress) String() string { return proto.CompactTextString(m) }
func (*CurrentRolloutProgress) ProtoMessage()    {}
func (*CurrentRolloutProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3bb6ceeab2493f2, []int{4}
}

func (m *CurrentRolloutProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentRolloutProgress.Unmarshal(m, b)
}
func (m *CurrentRolloutProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentRolloutProgress.Marshal(b, m, deterministic)
}
func (m *CurrentRolloutProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentRolloutProgress.Merge(m, src)
}
func (m *CurrentRolloutProgress) XXX_Size() int {
	return xxx_messageInfo_CurrentRolloutProgress.Size(m)
}
func (m *CurrentRolloutProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentRolloutProgress.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentRolloutProgress proto.InternalMessageInfo

func (m *CurrentRolloutProgress) GetRollout() *Rollout {
	if m != nil {
		return m.Rollout
	}
	return nil
}

func (m *CurrentRolloutProgress) GetNodeCount() int32 {
	if m != nil {
		return m.NodeCount
	}
	return 0
}

func (m *CurrentRolloutProgress) GetLatestRunSuccessfulCount() int32 {
	if m != nil {
		return m.LatestRunSuccessfulCount
	}
	return 0
}

func (m *CurrentRolloutProgress) GetLatestRunErroredCount() int32 {
	if m != nil {
		return m.LatestRunErroredCount
	}
	return 0
}

func (m *CurrentRolloutProgress) GetSourceLink() string {
	if m != nil {
		return m.SourceLink
	}
	return ""
}

func (m *CurrentRolloutProgress) GetBuildLink() string {
	if m != nil {
		return m.BuildLink
	}
	return ""
}

type PastRolloutProgress struct {
	// Rollout is the full rollout object, but we can change this to be a subset only.
	Rollout *Rollout `protobuf:"bytes,1,opt,name=rollout,proto3" json:"rollout,omitempty"`
	// The number of nodes in the node segment for which the last recorded CCR
	// was part of this rollout. Note that no breakdown of success/errored is
	// provided, since some nodes may have moved on to the current rollout and
	// are not included in the count.
	LatestRunNodeCount   int32    `protobuf:"varint,2,opt,name=latest_run_node_count,json=latestRunNodeCount,proto3" json:"latest_run_node_count,omitempty"`
	SourceLink           string   `protobuf:"bytes,5,opt,name=source_link,json=sourceLink,proto3" json:"source_link,omitempty"`
	BuildLink            string   `protobuf:"bytes,6,opt,name=build_link,json=buildLink,proto3" json:"build_link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PastRolloutProgress) Reset()         { *m = PastRolloutProgress{} }
func (m *PastRolloutProgress) String() string { return proto.CompactTextString(m) }
func (*PastRolloutProgress) ProtoMessage()    {}
func (*PastRolloutProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3bb6ceeab2493f2, []int{5}
}

func (m *PastRolloutProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PastRolloutProgress.Unmarshal(m, b)
}
func (m *PastRolloutProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PastRolloutProgress.Marshal(b, m, deterministic)
}
func (m *PastRolloutProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PastRolloutProgress.Merge(m, src)
}
func (m *PastRolloutProgress) XXX_Size() int {
	return xxx_messageInfo_PastRolloutProgress.Size(m)
}
func (m *PastRolloutProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_PastRolloutProgress.DiscardUnknown(m)
}

var xxx_messageInfo_PastRolloutProgress proto.InternalMessageInfo

func (m *PastRolloutProgress) GetRollout() *Rollout {
	if m != nil {
		return m.Rollout
	}
	return nil
}

func (m *PastRolloutProgress) GetLatestRunNodeCount() int32 {
	if m != nil {
		return m.LatestRunNodeCount
	}
	return 0
}

func (m *PastRolloutProgress) GetSourceLink() string {
	if m != nil {
		return m.SourceLink
	}
	return ""
}

func (m *PastRolloutProgress) GetBuildLink() string {
	if m != nil {
		return m.BuildLink
	}
	return ""
}

func init() {
	proto.RegisterEnum("chef.automate.api.cfgmgmt.response.SCMType", SCMType_name, SCMType_value)
	proto.RegisterEnum("chef.automate.api.cfgmgmt.response.SCMWebType", SCMWebType_name, SCMWebType_value)
	proto.RegisterType((*Rollout)(nil), "chef.automate.api.cfgmgmt.response.Rollout")
	proto.RegisterType((*Rollouts)(nil), "chef.automate.api.cfgmgmt.response.Rollouts")
	proto.RegisterType((*NodeSegmentsWithRolloutProgress)(nil), "chef.automate.api.cfgmgmt.response.NodeSegmentsWithRolloutProgress")
	proto.RegisterType((*NodeSegmentRolloutProgress)(nil), "chef.automate.api.cfgmgmt.response.NodeSegmentRolloutProgress")
	proto.RegisterType((*CurrentRolloutProgress)(nil), "chef.automate.api.cfgmgmt.response.CurrentRolloutProgress")
	proto.RegisterType((*PastRolloutProgress)(nil), "chef.automate.api.cfgmgmt.response.PastRolloutProgress")
}

func init() {
	proto.RegisterFile("api/external/cfgmgmt/response/rollouts.proto", fileDescriptor_d3bb6ceeab2493f2)
}

var fileDescriptor_d3bb6ceeab2493f2 = []byte{
	// 811 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xae, 0x37, 0xdd, 0xfc, 0x1c, 0x2f, 0x49, 0x3a, 0x55, 0x2b, 0xb3, 0x50, 0x25, 0x0a, 0x5c,
	0xac, 0xb6, 0xc5, 0x11, 0x45, 0xa2, 0x08, 0x09, 0x2e, 0x36, 0x2c, 0x21, 0x40, 0xc3, 0xca, 0xc9,
	0x2a, 0x12, 0x37, 0x96, 0x33, 0x9e, 0xcd, 0x0e, 0xf5, 0x78, 0xac, 0x99, 0x31, 0xb0, 0x97, 0xbc,
	0x02, 0xef, 0xc0, 0x2b, 0xf0, 0x1a, 0x3c, 0x0a, 0xaf, 0x80, 0xe6, 0xc7, 0xa9, 0xb5, 0x4d, 0x21,
	0xa8, 0xbd, 0x8b, 0xbe, 0xf3, 0x7d, 0xdf, 0x99, 0xf3, 0x17, 0xc3, 0x93, 0xa4, 0xa0, 0x63, 0xf2,
	0xab, 0x22, 0x22, 0x4f, 0xb2, 0x31, 0xbe, 0xda, 0xb0, 0x0d, 0x53, 0x63, 0x41, 0x64, 0xc1, 0x73,
	0x49, 0xc6, 0x82, 0x67, 0x19, 0x2f, 0x95, 0x0c, 0x0b, 0xc1, 0x15, 0x47, 0x23, 0x7c, 0x4d, 0xae,
	0xc2, 0xa4, 0x54, 0x9c, 0x25, 0x8a, 0x84, 0x49, 0x41, 0x43, 0x27, 0x09, 0x2b, 0xc9, 0xe8, 0xef,
	0xbb, 0xd0, 0x8a, 0xac, 0x0c, 0x0d, 0xc0, 0x2f, 0x78, 0x46, 0xf1, 0x4d, 0x9c, 0x27, 0x8c, 0x04,
	0xde, 0xd0, 0x3b, 0xe9, 0x44, 0x60, 0xa1, 0x79, 0xc2, 0x08, 0x3a, 0x85, 0x7b, 0x15, 0x81, 0xa7,
	0x24, 0xde, 0x08, 0x5e, 0x16, 0xc1, 0x81, 0xa1, 0xf5, 0x1c, 0x8d, 0xa7, 0x64, 0xaa, 0x61, 0xf4,
	0x04, 0x90, 0xe3, 0x0a, 0xf2, 0x33, 0x95, 0x94, 0xe7, 0x31, 0x4d, 0x83, 0x86, 0x21, 0xf7, 0x6d,
	0x24, 0x72, 0x81, 0x59, 0x5a, 0x73, 0x4e, 0x39, 0x4b, 0x68, 0x1e, 0x97, 0x22, 0x0b, 0xee, 0xd6,
	0x9d, 0xbf, 0x32, 0xf8, 0xa5, 0xc8, 0xd0, 0xd7, 0xd0, 0x96, 0x98, 0xc5, 0xea, 0xa6, 0x20, 0xc1,
	0xe1, 0xd0, 0x3b, 0xe9, 0x3e, 0x7d, 0x1c, 0xfe, 0x77, 0xa5, 0xe1, 0x62, 0xf2, 0x7c, 0x79, 0x53,
	0x90, 0xa8, 0x25, 0x31, 0xd3, 0x3f, 0xd0, 0x05, 0x1c, 0x69, 0x9f, 0x5f, 0xc8, 0xda, 0x7a, 0x35,
	0x8d, 0x57, 0xb8, 0xa7, 0xd7, 0x8a, 0xac, 0x8d, 0x1d, 0x48, 0xcc, 0xdc, 0x6f, 0xf4, 0x21, 0x74,
	0x5d, 0x15, 0xda, 0x58, 0x97, 0xd0, 0x32, 0x25, 0x1c, 0x59, 0x74, 0x81, 0x99, 0x7e, 0xff, 0xe3,
	0x6d, 0x67, 0xaa, 0xf4, 0x9a, 0xd9, 0xae, 0x17, 0xbb, 0x30, 0x9e, 0x9a, 0xfc, 0xb2, 0x31, 0x9a,
	0x8c, 0x39, 0x63, 0x54, 0x05, 0x9d, 0x5b, 0xdc, 0x89, 0x81, 0xd1, 0x10, 0xfc, 0x94, 0x48, 0x2c,
	0x68, 0xa1, 0x28, 0xcf, 0x03, 0x30, 0xac, 0x3a, 0x84, 0xde, 0x07, 0xc0, 0x34, 0xfe, 0x89, 0xdb,
	0x94, 0xbe, 0x21, 0xb4, 0x31, 0xfd, 0x96, 0x9b, 0x5c, 0xc7, 0xd0, 0x71, 0x51, 0x9a, 0x06, 0x47,
	0x26, 0xd8, 0x32, 0xc1, 0x59, 0x8a, 0xba, 0x70, 0x40, 0xd3, 0xe0, 0x1d, 0x03, 0x1e, 0xd0, 0x14,
	0x3d, 0x02, 0x90, 0x2a, 0x11, 0x2a, 0x56, 0x94, 0x91, 0xa0, 0x6b, 0xf0, 0x8e, 0x41, 0x96, 0x94,
	0x11, 0xf4, 0x2e, 0xb4, 0x49, 0x9e, 0xda, 0x60, 0xcf, 0x3a, 0x91, 0x3c, 0xd5, 0xa1, 0xd1, 0x02,
	0xda, 0x6e, 0xe1, 0x24, 0x9a, 0x42, 0xbb, 0xda, 0xd9, 0xc0, 0x1b, 0x36, 0x4e, 0xfc, 0xfd, 0x46,
	0xe9, 0xf4, 0xd1, 0x56, 0x3c, 0xfa, 0xc3, 0x83, 0x81, 0xde, 0xbd, 0x05, 0xd9, 0x30, 0x92, 0x2b,
	0xb9, 0xa2, 0xea, 0xda, 0xb1, 0x2e, 0x04, 0xdf, 0x08, 0x22, 0x25, 0xfa, 0xcd, 0x83, 0x47, 0x66,
	0x6f, 0xa5, 0x25, 0xc5, 0x4e, 0x1d, 0x17, 0x8e, 0xe1, 0x9e, 0xf0, 0xe5, 0x3e, 0x4f, 0xa8, 0x25,
	0xbb, 0x95, 0x27, 0x3a, 0xce, 0x5f, 0x1b, 0x1b, 0xfd, 0xde, 0x80, 0xe3, 0xd7, 0x4b, 0xdf, 0xee,
	0x05, 0xee, 0xbc, 0xa9, 0xc6, 0xee, 0x9b, 0x1a, 0x80, 0xaf, 0xb8, 0x4a, 0x32, 0x63, 0x2b, 0xcd,
	0xe5, 0x1d, 0x46, 0x60, 0x20, 0x6d, 0x28, 0x91, 0x82, 0x00, 0x97, 0x42, 0xec, 0x6c, 0x9b, 0x3e,
	0x42, 0xff, 0xe9, 0xe7, 0xfb, 0xb4, 0x6d, 0x62, 0x3d, 0x6e, 0xb7, 0xec, 0x21, 0xde, 0x89, 0xa3,
	0x14, 0xee, 0x15, 0xfa, 0xef, 0x83, 0x97, 0x32, 0xde, 0x2e, 0x4a, 0xd3, 0x4c, 0xe9, 0xd9, 0x3e,
	0xe9, 0x2e, 0x12, 0xf9, 0x4a, 0xae, 0x7e, 0xe5, 0x58, 0x6d, 0xe1, 0xe8, 0xcf, 0x03, 0x78, 0xb8,
	0xfb, 0x61, 0xe8, 0x1c, 0x5a, 0x2e, 0xaf, 0x19, 0xc6, 0xff, 0xdc, 0xcf, 0x4a, 0xab, 0xaf, 0xc5,
	0xcc, 0x0b, 0xf3, 0x32, 0x57, 0x66, 0x5e, 0x87, 0x51, 0x47, 0x23, 0x13, 0x0d, 0xa0, 0x2f, 0xe0,
	0xbd, 0x2c, 0x51, 0x44, 0xaa, 0x58, 0x94, 0x79, 0x2c, 0x4b, 0x8c, 0x89, 0x94, 0x57, 0x65, 0xe6,
	0xf8, 0x0d, 0xc3, 0x0f, 0x2c, 0x25, 0x2a, 0xf3, 0xc5, 0x96, 0x60, 0xe5, 0xcf, 0x20, 0xa8, 0xc9,
	0x89, 0x10, 0x5c, 0x90, 0xd4, 0x69, 0xed, 0x24, 0x1f, 0x6c, 0xb5, 0xe7, 0x36, 0x6a, 0x85, 0x03,
	0xf0, 0x25, 0x2f, 0x05, 0x26, 0x71, 0x46, 0xf3, 0x17, 0x66, 0x8e, 0x9d, 0x08, 0x2c, 0xf4, 0x3d,
	0xcd, 0x5f, 0xe8, 0x77, 0xaf, 0x4b, 0x9a, 0xa5, 0x36, 0xde, 0xb4, 0x57, 0x6e, 0x10, 0x1d, 0x1e,
	0xfd, 0xe5, 0xc1, 0xfd, 0x1d, 0x2d, 0x7e, 0x5b, 0x5d, 0xfb, 0x18, 0x1e, 0xd4, 0xea, 0x7a, 0xa5,
	0x81, 0x68, 0x5b, 0xd4, 0x7c, 0xdb, 0xc9, 0x37, 0xac, 0xe8, 0xf4, 0x03, 0x68, 0xb9, 0xef, 0x04,
	0xea, 0x81, 0x7f, 0x39, 0xff, 0x6e, 0xfe, 0xc3, 0x6a, 0x1e, 0x2f, 0x26, 0xcf, 0xfb, 0x77, 0x50,
	0x0b, 0x1a, 0xd3, 0xd9, 0xb2, 0xef, 0x9d, 0x7e, 0x04, 0xf0, 0xf2, 0x03, 0x80, 0xee, 0x43, 0xaf,
	0xc6, 0x8b, 0x57, 0xe7, 0x67, 0xfd, 0x3b, 0x08, 0xa0, 0x39, 0x9d, 0x2d, 0xbf, 0xb9, 0x3c, 0xeb,
	0x7b, 0x67, 0x9f, 0xfd, 0xf8, 0xe9, 0x86, 0xaa, 0xeb, 0x72, 0x1d, 0x62, 0xce, 0xc6, 0xba, 0x11,
	0xe3, 0xaa, 0x11, 0xe3, 0x7f, 0xfd, 0x9e, 0xaf, 0x9b, 0xe6, 0x3b, 0xfe, 0xc9, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0x8a, 0x33, 0xdd, 0xf7, 0x07, 0x00, 0x00,
}
