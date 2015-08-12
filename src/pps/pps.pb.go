// Code generated by protoc-gen-go.
// source: pps/pps.proto
// DO NOT EDIT!

/*
Package pps is a generated protocol buffer package.

It is generated from these files:
	pps/pps.proto

It has these top-level messages:
	PipelineRunStatus
	Input
	Output
	Node
	DockerService
	Element
	Pipeline
	GithubPipelineSource
	PipelineSource
	PipelineRun
	PipelineRunContainer
	PipelineRunLog
	GetPipelineRequest
	GetPipelineResponse
	StartPipelineRunRequest
	StartPipelineRunResponse
	GetPipelineRunStatusRequest
	GetPipelineRunStatusResponse
	GetPipelineRunLogsRequest
	GetPipelineRunLogsResponse
*/
package pps

import proto "github.com/golang/protobuf/proto"
import google_protobuf "github.com/peter-edge/go-google-protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type PipelineRunStatusType int32

const (
	PipelineRunStatusType_PIPELINE_RUN_STATUS_TYPE_NONE    PipelineRunStatusType = 0
	PipelineRunStatusType_PIPELINE_RUN_STATUS_TYPE_ADDED   PipelineRunStatusType = 1
	PipelineRunStatusType_PIPELINE_RUN_STATUS_TYPE_STARTED PipelineRunStatusType = 2
	PipelineRunStatusType_PIPELINE_RUN_STATUS_TYPE_ERROR   PipelineRunStatusType = 3
	PipelineRunStatusType_PIPELINE_RUN_STATUS_TYPE_SUCCESS PipelineRunStatusType = 4
)

var PipelineRunStatusType_name = map[int32]string{
	0: "PIPELINE_RUN_STATUS_TYPE_NONE",
	1: "PIPELINE_RUN_STATUS_TYPE_ADDED",
	2: "PIPELINE_RUN_STATUS_TYPE_STARTED",
	3: "PIPELINE_RUN_STATUS_TYPE_ERROR",
	4: "PIPELINE_RUN_STATUS_TYPE_SUCCESS",
}
var PipelineRunStatusType_value = map[string]int32{
	"PIPELINE_RUN_STATUS_TYPE_NONE":    0,
	"PIPELINE_RUN_STATUS_TYPE_ADDED":   1,
	"PIPELINE_RUN_STATUS_TYPE_STARTED": 2,
	"PIPELINE_RUN_STATUS_TYPE_ERROR":   3,
	"PIPELINE_RUN_STATUS_TYPE_SUCCESS": 4,
}

func (x PipelineRunStatusType) String() string {
	return proto.EnumName(PipelineRunStatusType_name, int32(x))
}

type OutputStream int32

const (
	OutputStream_OUTPUT_STREAM_NONE   OutputStream = 0
	OutputStream_OUTPUT_STREAM_STDOUT OutputStream = 1
	OutputStream_OUTPUT_STREAM_STDERR OutputStream = 2
)

var OutputStream_name = map[int32]string{
	0: "OUTPUT_STREAM_NONE",
	1: "OUTPUT_STREAM_STDOUT",
	2: "OUTPUT_STREAM_STDERR",
}
var OutputStream_value = map[string]int32{
	"OUTPUT_STREAM_NONE":   0,
	"OUTPUT_STREAM_STDOUT": 1,
	"OUTPUT_STREAM_STDERR": 2,
}

func (x OutputStream) String() string {
	return proto.EnumName(OutputStream_name, int32(x))
}

type PipelineRunStatus struct {
	PipelineRunId         string                     `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
	PipelineRunStatusType PipelineRunStatusType      `protobuf:"varint,2,opt,name=pipeline_run_status_type,enum=pps.PipelineRunStatusType" json:"pipeline_run_status_type,omitempty"`
	Timestamp             *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *PipelineRunStatus) Reset()         { *m = PipelineRunStatus{} }
func (m *PipelineRunStatus) String() string { return proto.CompactTextString(m) }
func (*PipelineRunStatus) ProtoMessage()    {}

func (m *PipelineRunStatus) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Input struct {
	Node []string          `protobuf:"bytes,1,rep,name=node" json:"node,omitempty"`
	Host map[string]string `protobuf:"bytes,2,rep,name=host" json:"host,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Pfs  map[string]string `protobuf:"bytes,3,rep,name=pfs" json:"pfs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}

func (m *Input) GetHost() map[string]string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Input) GetPfs() map[string]string {
	if m != nil {
		return m.Pfs
	}
	return nil
}

type Output struct {
	Host map[string]string `protobuf:"bytes,1,rep,name=host" json:"host,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Pfs  map[string]string `protobuf:"bytes,2,rep,name=pfs" json:"pfs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}

func (m *Output) GetHost() map[string]string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Output) GetPfs() map[string]string {
	if m != nil {
		return m.Pfs
	}
	return nil
}

type Node struct {
	Service string   `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Input   *Input   `protobuf:"bytes,2,opt,name=input" json:"input,omitempty"`
	Output  *Output  `protobuf:"bytes,3,opt,name=output" json:"output,omitempty"`
	Run     []string `protobuf:"bytes,4,rep,name=run" json:"run,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}

func (m *Node) GetInput() *Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Node) GetOutput() *Output {
	if m != nil {
		return m.Output
	}
	return nil
}

type DockerService struct {
	Image      string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Build      string `protobuf:"bytes,2,opt,name=build" json:"build,omitempty"`
	Dockerfile string `protobuf:"bytes,3,opt,name=dockerfile" json:"dockerfile,omitempty"`
}

func (m *DockerService) Reset()         { *m = DockerService{} }
func (m *DockerService) String() string { return proto.CompactTextString(m) }
func (*DockerService) ProtoMessage()    {}

type Element struct {
	Name          string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Node          *Node          `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	DockerService *DockerService `protobuf:"bytes,3,opt,name=docker_service" json:"docker_service,omitempty"`
}

func (m *Element) Reset()         { *m = Element{} }
func (m *Element) String() string { return proto.CompactTextString(m) }
func (*Element) ProtoMessage()    {}

func (m *Element) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Element) GetDockerService() *DockerService {
	if m != nil {
		return m.DockerService
	}
	return nil
}

type Pipeline struct {
	NameToElement map[string]*Element `protobuf:"bytes,1,rep,name=name_to_element" json:"name_to_element,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Pipeline) Reset()         { *m = Pipeline{} }
func (m *Pipeline) String() string { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()    {}

func (m *Pipeline) GetNameToElement() map[string]*Element {
	if m != nil {
		return m.NameToElement
	}
	return nil
}

type GithubPipelineSource struct {
	ContextDir  string `protobuf:"bytes,1,opt,name=context_dir" json:"context_dir,omitempty"`
	User        string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Repository  string `protobuf:"bytes,3,opt,name=repository" json:"repository,omitempty"`
	Branch      string `protobuf:"bytes,4,opt,name=branch" json:"branch,omitempty"`
	AccessToken string `protobuf:"bytes,5,opt,name=access_token" json:"access_token,omitempty"`
}

func (m *GithubPipelineSource) Reset()         { *m = GithubPipelineSource{} }
func (m *GithubPipelineSource) String() string { return proto.CompactTextString(m) }
func (*GithubPipelineSource) ProtoMessage()    {}

type PipelineSource struct {
	GithubPipelineSource *GithubPipelineSource `protobuf:"bytes,1,opt,name=github_pipeline_source" json:"github_pipeline_source,omitempty"`
}

func (m *PipelineSource) Reset()         { *m = PipelineSource{} }
func (m *PipelineSource) String() string { return proto.CompactTextString(m) }
func (*PipelineSource) ProtoMessage()    {}

func (m *PipelineSource) GetGithubPipelineSource() *GithubPipelineSource {
	if m != nil {
		return m.GithubPipelineSource
	}
	return nil
}

type PipelineRun struct {
	Id             string          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	PipelineSource *PipelineSource `protobuf:"bytes,2,opt,name=pipeline_source" json:"pipeline_source,omitempty"`
	Pipeline       *Pipeline       `protobuf:"bytes,3,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *PipelineRun) Reset()         { *m = PipelineRun{} }
func (m *PipelineRun) String() string { return proto.CompactTextString(m) }
func (*PipelineRun) ProtoMessage()    {}

func (m *PipelineRun) GetPipelineSource() *PipelineSource {
	if m != nil {
		return m.PipelineSource
	}
	return nil
}

func (m *PipelineRun) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type PipelineRunContainer struct {
	PipelineRunId string `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
	ContainerId   string `protobuf:"bytes,2,opt,name=container_id" json:"container_id,omitempty"`
	Node          string `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
}

func (m *PipelineRunContainer) Reset()         { *m = PipelineRunContainer{} }
func (m *PipelineRunContainer) String() string { return proto.CompactTextString(m) }
func (*PipelineRunContainer) ProtoMessage()    {}

type PipelineRunLog struct {
	PipelineRunId string                     `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
	ContainerId   string                     `protobuf:"bytes,2,opt,name=container_id" json:"container_id,omitempty"`
	Node          string                     `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
	Timestamp     *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	OutputStream  OutputStream               `protobuf:"varint,5,opt,name=output_stream,enum=pps.OutputStream" json:"output_stream,omitempty"`
	Data          []byte                     `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PipelineRunLog) Reset()         { *m = PipelineRunLog{} }
func (m *PipelineRunLog) String() string { return proto.CompactTextString(m) }
func (*PipelineRunLog) ProtoMessage()    {}

func (m *PipelineRunLog) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type GetPipelineRequest struct {
	PipelineSource *PipelineSource `protobuf:"bytes,1,opt,name=pipeline_source" json:"pipeline_source,omitempty"`
}

func (m *GetPipelineRequest) Reset()         { *m = GetPipelineRequest{} }
func (m *GetPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRequest) ProtoMessage()    {}

func (m *GetPipelineRequest) GetPipelineSource() *PipelineSource {
	if m != nil {
		return m.PipelineSource
	}
	return nil
}

type GetPipelineResponse struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *GetPipelineResponse) Reset()         { *m = GetPipelineResponse{} }
func (m *GetPipelineResponse) String() string { return proto.CompactTextString(m) }
func (*GetPipelineResponse) ProtoMessage()    {}

func (m *GetPipelineResponse) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type StartPipelineRunRequest struct {
	PipelineSource *PipelineSource `protobuf:"bytes,1,opt,name=pipeline_source" json:"pipeline_source,omitempty"`
}

func (m *StartPipelineRunRequest) Reset()         { *m = StartPipelineRunRequest{} }
func (m *StartPipelineRunRequest) String() string { return proto.CompactTextString(m) }
func (*StartPipelineRunRequest) ProtoMessage()    {}

func (m *StartPipelineRunRequest) GetPipelineSource() *PipelineSource {
	if m != nil {
		return m.PipelineSource
	}
	return nil
}

type StartPipelineRunResponse struct {
	PipelineRunId string `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
}

func (m *StartPipelineRunResponse) Reset()         { *m = StartPipelineRunResponse{} }
func (m *StartPipelineRunResponse) String() string { return proto.CompactTextString(m) }
func (*StartPipelineRunResponse) ProtoMessage()    {}

type GetPipelineRunStatusRequest struct {
	PipelineRunId string `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
}

func (m *GetPipelineRunStatusRequest) Reset()         { *m = GetPipelineRunStatusRequest{} }
func (m *GetPipelineRunStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRunStatusRequest) ProtoMessage()    {}

type GetPipelineRunStatusResponse struct {
	PipelineRunStatus *PipelineRunStatus `protobuf:"bytes,1,opt,name=pipeline_run_status" json:"pipeline_run_status,omitempty"`
}

func (m *GetPipelineRunStatusResponse) Reset()         { *m = GetPipelineRunStatusResponse{} }
func (m *GetPipelineRunStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRunStatusResponse) ProtoMessage()    {}

func (m *GetPipelineRunStatusResponse) GetPipelineRunStatus() *PipelineRunStatus {
	if m != nil {
		return m.PipelineRunStatus
	}
	return nil
}

type GetPipelineRunLogsRequest struct {
	PipelineRunId string `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
	Node          string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
}

func (m *GetPipelineRunLogsRequest) Reset()         { *m = GetPipelineRunLogsRequest{} }
func (m *GetPipelineRunLogsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRunLogsRequest) ProtoMessage()    {}

type GetPipelineRunLogsResponse struct {
	PipelineRunLog []*PipelineRunLog `protobuf:"bytes,1,rep,name=pipeline_run_log" json:"pipeline_run_log,omitempty"`
}

func (m *GetPipelineRunLogsResponse) Reset()         { *m = GetPipelineRunLogsResponse{} }
func (m *GetPipelineRunLogsResponse) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRunLogsResponse) ProtoMessage()    {}

func (m *GetPipelineRunLogsResponse) GetPipelineRunLog() []*PipelineRunLog {
	if m != nil {
		return m.PipelineRunLog
	}
	return nil
}

func init() {
	proto.RegisterEnum("pps.PipelineRunStatusType", PipelineRunStatusType_name, PipelineRunStatusType_value)
	proto.RegisterEnum("pps.OutputStream", OutputStream_name, OutputStream_value)
}

// Client API for Api service

type ApiClient interface {
	GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*GetPipelineResponse, error)
	StartPipelineRun(ctx context.Context, in *StartPipelineRunRequest, opts ...grpc.CallOption) (*StartPipelineRunResponse, error)
	GetPipelineRunStatus(ctx context.Context, in *GetPipelineRunStatusRequest, opts ...grpc.CallOption) (*GetPipelineRunStatusResponse, error)
	GetPipelineRunLogs(ctx context.Context, in *GetPipelineRunLogsRequest, opts ...grpc.CallOption) (*GetPipelineRunLogsResponse, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*GetPipelineResponse, error) {
	out := new(GetPipelineResponse)
	err := grpc.Invoke(ctx, "/pps.Api/GetPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) StartPipelineRun(ctx context.Context, in *StartPipelineRunRequest, opts ...grpc.CallOption) (*StartPipelineRunResponse, error) {
	out := new(StartPipelineRunResponse)
	err := grpc.Invoke(ctx, "/pps.Api/StartPipelineRun", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetPipelineRunStatus(ctx context.Context, in *GetPipelineRunStatusRequest, opts ...grpc.CallOption) (*GetPipelineRunStatusResponse, error) {
	out := new(GetPipelineRunStatusResponse)
	err := grpc.Invoke(ctx, "/pps.Api/GetPipelineRunStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetPipelineRunLogs(ctx context.Context, in *GetPipelineRunLogsRequest, opts ...grpc.CallOption) (*GetPipelineRunLogsResponse, error) {
	out := new(GetPipelineRunLogsResponse)
	err := grpc.Invoke(ctx, "/pps.Api/GetPipelineRunLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiServer interface {
	GetPipeline(context.Context, *GetPipelineRequest) (*GetPipelineResponse, error)
	StartPipelineRun(context.Context, *StartPipelineRunRequest) (*StartPipelineRunResponse, error)
	GetPipelineRunStatus(context.Context, *GetPipelineRunStatusRequest) (*GetPipelineRunStatusResponse, error)
	GetPipelineRunLogs(context.Context, *GetPipelineRunLogsRequest) (*GetPipelineRunLogsResponse, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_GetPipeline_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(GetPipelineRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetPipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_StartPipelineRun_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(StartPipelineRunRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).StartPipelineRun(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_GetPipelineRunStatus_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(GetPipelineRunStatusRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetPipelineRunStatus(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_GetPipelineRunLogs_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(GetPipelineRunLogsRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetPipelineRunLogs(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pps.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPipeline",
			Handler:    _Api_GetPipeline_Handler,
		},
		{
			MethodName: "StartPipelineRun",
			Handler:    _Api_StartPipelineRun_Handler,
		},
		{
			MethodName: "GetPipelineRunStatus",
			Handler:    _Api_GetPipelineRunStatus_Handler,
		},
		{
			MethodName: "GetPipelineRunLogs",
			Handler:    _Api_GetPipelineRunLogs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
