// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: compute/v1/compute_instance.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationComputeInstanceCreate = "/github.com.mohaijiang.api.compute.v1.ComputeInstance/Create"
const OperationComputeInstanceDelete = "/github.com.mohaijiang.api.compute.v1.ComputeInstance/Delete"
const OperationComputeInstanceGet = "/github.com.mohaijiang.api.compute.v1.ComputeInstance/Get"
const OperationComputeInstanceList = "/github.com.mohaijiang.api.compute.v1.ComputeInstance/List"
const OperationComputeInstanceListComputeImage = "/github.com.mohaijiang.api.compute.v1.ComputeInstance/ListComputeImage"
const OperationComputeInstanceListComputeInstanceDuration = "/github.com.mohaijiang.api.compute.v1.ComputeInstance/ListComputeInstanceDuration"
const OperationComputeInstanceListComputeSpec = "/github.com.mohaijiang.api.compute.v1.ComputeInstance/ListComputeSpec"
const OperationComputeInstanceStartInstance = "/github.com.mohaijiang.api.compute.v1.ComputeInstance/StartInstance"
const OperationComputeInstanceStopInstance = "/github.com.mohaijiang.api.compute.v1.ComputeInstance/StopInstance"

type ComputeInstanceHTTPServer interface {
	// Create 创建实例
	Create(context.Context, *CreateInstanceRequest) (*CreateInstanceReply, error)
	// Delete删除实例
	Delete(context.Context, *DeleteInstanceRequest) (*DeleteInstanceReply, error)
	// Get获取实例详情
	Get(context.Context, *GetInstanceRequest) (*GetInstanceReply, error)
	// List实例列表
	List(context.Context, *ListInstanceRequest) (*ListInstanceReply, error)
	// ListComputeImage 查询镜像
	ListComputeImage(context.Context, *ListComputeImageRequest) (*ListComputeImageReply, error)
	// ListComputeInstanceDuration 查询到期时间
	ListComputeInstanceDuration(context.Context, *ListComputeDurationRequest) (*ListComputeDurationReply, error)
	// ListComputeSpec 查询规格
	ListComputeSpec(context.Context, *ListComputeSpecRequest) (*ListComputeSpecReply, error)
	// StartInstance启动实例
	StartInstance(context.Context, *GetInstanceRequest) (*StartInstanceReply, error)
	// StopInstance停止实例
	StopInstance(context.Context, *GetInstanceRequest) (*StopInstanceReply, error)
}

func RegisterComputeInstanceHTTPServer(s *http.Server, srv ComputeInstanceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/compute/spec", _ComputeInstance_ListComputeSpec0_HTTP_Handler(srv))
	r.GET("/v1/compute/image", _ComputeInstance_ListComputeImage0_HTTP_Handler(srv))
	r.GET("/v1/compute/duration", _ComputeInstance_ListComputeInstanceDuration0_HTTP_Handler(srv))
	r.POST("/v1/instance", _ComputeInstance_Create0_HTTP_Handler(srv))
	r.DELETE("/v1/instance/{id}", _ComputeInstance_Delete0_HTTP_Handler(srv))
	r.GET("/v1/instance/{id}", _ComputeInstance_Get0_HTTP_Handler(srv))
	r.GET("/v1/instance", _ComputeInstance_List0_HTTP_Handler(srv))
	r.PUT("/v1/instance/{id}/stop", _ComputeInstance_StopInstance0_HTTP_Handler(srv))
	r.PUT("/v1/instance/{id}/start", _ComputeInstance_StartInstance0_HTTP_Handler(srv))
}

func _ComputeInstance_ListComputeSpec0_HTTP_Handler(srv ComputeInstanceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListComputeSpecRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputeInstanceListComputeSpec)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListComputeSpec(ctx, req.(*ListComputeSpecRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListComputeSpecReply)
		return ctx.Result(200, reply)
	}
}

func _ComputeInstance_ListComputeImage0_HTTP_Handler(srv ComputeInstanceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListComputeImageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputeInstanceListComputeImage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListComputeImage(ctx, req.(*ListComputeImageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListComputeImageReply)
		return ctx.Result(200, reply)
	}
}

func _ComputeInstance_ListComputeInstanceDuration0_HTTP_Handler(srv ComputeInstanceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListComputeDurationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputeInstanceListComputeInstanceDuration)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListComputeInstanceDuration(ctx, req.(*ListComputeDurationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListComputeDurationReply)
		return ctx.Result(200, reply)
	}
}

func _ComputeInstance_Create0_HTTP_Handler(srv ComputeInstanceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateInstanceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputeInstanceCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*CreateInstanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateInstanceReply)
		return ctx.Result(200, reply)
	}
}

func _ComputeInstance_Delete0_HTTP_Handler(srv ComputeInstanceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteInstanceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputeInstanceDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*DeleteInstanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteInstanceReply)
		return ctx.Result(200, reply)
	}
}

func _ComputeInstance_Get0_HTTP_Handler(srv ComputeInstanceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetInstanceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputeInstanceGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*GetInstanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetInstanceReply)
		return ctx.Result(200, reply)
	}
}

func _ComputeInstance_List0_HTTP_Handler(srv ComputeInstanceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListInstanceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputeInstanceList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*ListInstanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListInstanceReply)
		return ctx.Result(200, reply)
	}
}

func _ComputeInstance_StopInstance0_HTTP_Handler(srv ComputeInstanceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetInstanceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputeInstanceStopInstance)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StopInstance(ctx, req.(*GetInstanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StopInstanceReply)
		return ctx.Result(200, reply)
	}
}

func _ComputeInstance_StartInstance0_HTTP_Handler(srv ComputeInstanceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetInstanceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputeInstanceStartInstance)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StartInstance(ctx, req.(*GetInstanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StartInstanceReply)
		return ctx.Result(200, reply)
	}
}

type ComputeInstanceHTTPClient interface {
	Create(ctx context.Context, req *CreateInstanceRequest, opts ...http.CallOption) (rsp *CreateInstanceReply, err error)
	Delete(ctx context.Context, req *DeleteInstanceRequest, opts ...http.CallOption) (rsp *DeleteInstanceReply, err error)
	Get(ctx context.Context, req *GetInstanceRequest, opts ...http.CallOption) (rsp *GetInstanceReply, err error)
	List(ctx context.Context, req *ListInstanceRequest, opts ...http.CallOption) (rsp *ListInstanceReply, err error)
	ListComputeImage(ctx context.Context, req *ListComputeImageRequest, opts ...http.CallOption) (rsp *ListComputeImageReply, err error)
	ListComputeInstanceDuration(ctx context.Context, req *ListComputeDurationRequest, opts ...http.CallOption) (rsp *ListComputeDurationReply, err error)
	ListComputeSpec(ctx context.Context, req *ListComputeSpecRequest, opts ...http.CallOption) (rsp *ListComputeSpecReply, err error)
	StartInstance(ctx context.Context, req *GetInstanceRequest, opts ...http.CallOption) (rsp *StartInstanceReply, err error)
	StopInstance(ctx context.Context, req *GetInstanceRequest, opts ...http.CallOption) (rsp *StopInstanceReply, err error)
}

type ComputeInstanceHTTPClientImpl struct {
	cc *http.Client
}

func NewComputeInstanceHTTPClient(client *http.Client) ComputeInstanceHTTPClient {
	return &ComputeInstanceHTTPClientImpl{client}
}

func (c *ComputeInstanceHTTPClientImpl) Create(ctx context.Context, in *CreateInstanceRequest, opts ...http.CallOption) (*CreateInstanceReply, error) {
	var out CreateInstanceReply
	pattern := "/v1/instance"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationComputeInstanceCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputeInstanceHTTPClientImpl) Delete(ctx context.Context, in *DeleteInstanceRequest, opts ...http.CallOption) (*DeleteInstanceReply, error) {
	var out DeleteInstanceReply
	pattern := "/v1/instance/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationComputeInstanceDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputeInstanceHTTPClientImpl) Get(ctx context.Context, in *GetInstanceRequest, opts ...http.CallOption) (*GetInstanceReply, error) {
	var out GetInstanceReply
	pattern := "/v1/instance/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationComputeInstanceGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputeInstanceHTTPClientImpl) List(ctx context.Context, in *ListInstanceRequest, opts ...http.CallOption) (*ListInstanceReply, error) {
	var out ListInstanceReply
	pattern := "/v1/instance"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationComputeInstanceList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputeInstanceHTTPClientImpl) ListComputeImage(ctx context.Context, in *ListComputeImageRequest, opts ...http.CallOption) (*ListComputeImageReply, error) {
	var out ListComputeImageReply
	pattern := "/v1/compute/image"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationComputeInstanceListComputeImage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputeInstanceHTTPClientImpl) ListComputeInstanceDuration(ctx context.Context, in *ListComputeDurationRequest, opts ...http.CallOption) (*ListComputeDurationReply, error) {
	var out ListComputeDurationReply
	pattern := "/v1/compute/duration"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationComputeInstanceListComputeInstanceDuration))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputeInstanceHTTPClientImpl) ListComputeSpec(ctx context.Context, in *ListComputeSpecRequest, opts ...http.CallOption) (*ListComputeSpecReply, error) {
	var out ListComputeSpecReply
	pattern := "/v1/compute/spec"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationComputeInstanceListComputeSpec))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputeInstanceHTTPClientImpl) StartInstance(ctx context.Context, in *GetInstanceRequest, opts ...http.CallOption) (*StartInstanceReply, error) {
	var out StartInstanceReply
	pattern := "/v1/instance/{id}/start"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationComputeInstanceStartInstance))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ComputeInstanceHTTPClientImpl) StopInstance(ctx context.Context, in *GetInstanceRequest, opts ...http.CallOption) (*StopInstanceReply, error) {
	var out StopInstanceReply
	pattern := "/v1/instance/{id}/stop"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationComputeInstanceStopInstance))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
