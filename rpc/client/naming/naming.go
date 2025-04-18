// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: naming.proto

package naming

import (
	"context"

	"naming/rpc/naming"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdListReq     = naming.IdListReq
	IdReq         = naming.IdReq
	NamingInfoRes = naming.NamingInfoRes
	NamingListRes = naming.NamingListRes
	NamingReq     = naming.NamingReq
	PageResult    = naming.PageResult
	StringRes     = naming.StringRes

	Naming interface {
		// 查询示例列表
		GetNamingList(ctx context.Context, in *NamingReq, opts ...grpc.CallOption) (*NamingListRes, error)
		// 查询示例详情
		GetNamingInfo(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*NamingInfoRes, error)
		// 新增示例
		AddNaming(ctx context.Context, in *NamingInfoRes, opts ...grpc.CallOption) (*StringRes, error)
		// 删除示例
		DelNaming(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*StringRes, error)
		// 批量删除示例
		BatchDelNaming(ctx context.Context, in *IdListReq, opts ...grpc.CallOption) (*StringRes, error)
		// 修改示例
		UpdateNaming(ctx context.Context, in *NamingInfoRes, opts ...grpc.CallOption) (*StringRes, error)
	}

	defaultNaming struct {
		cli zrpc.Client
	}
)

func NewNaming(cli zrpc.Client) Naming {
	return &defaultNaming{
		cli: cli,
	}
}

// 查询示例列表
func (m *defaultNaming) GetNamingList(ctx context.Context, in *NamingReq, opts ...grpc.CallOption) (*NamingListRes, error) {
	client := naming.NewNamingClient(m.cli.Conn())
	return client.GetNamingList(ctx, in, opts...)
}

// 查询示例详情
func (m *defaultNaming) GetNamingInfo(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*NamingInfoRes, error) {
	client := naming.NewNamingClient(m.cli.Conn())
	return client.GetNamingInfo(ctx, in, opts...)
}

// 新增示例
func (m *defaultNaming) AddNaming(ctx context.Context, in *NamingInfoRes, opts ...grpc.CallOption) (*StringRes, error) {
	client := naming.NewNamingClient(m.cli.Conn())
	return client.AddNaming(ctx, in, opts...)
}

// 删除示例
func (m *defaultNaming) DelNaming(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*StringRes, error) {
	client := naming.NewNamingClient(m.cli.Conn())
	return client.DelNaming(ctx, in, opts...)
}

// 批量删除示例
func (m *defaultNaming) BatchDelNaming(ctx context.Context, in *IdListReq, opts ...grpc.CallOption) (*StringRes, error) {
	client := naming.NewNamingClient(m.cli.Conn())
	return client.BatchDelNaming(ctx, in, opts...)
}

// 修改示例
func (m *defaultNaming) UpdateNaming(ctx context.Context, in *NamingInfoRes, opts ...grpc.CallOption) (*StringRes, error) {
	client := naming.NewNamingClient(m.cli.Conn())
	return client.UpdateNaming(ctx, in, opts...)
}
