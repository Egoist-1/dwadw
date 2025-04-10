package naminglogic

import (
	"context"

	"naming/rpc/internal/svc"
	"naming/rpc/naming"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNamingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddNamingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNamingLogic {
	return &AddNamingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增示例
func (l *AddNamingLogic) AddNaming(in *naming.NamingInfoRes) (*naming.StringRes, error) {
	// todo: add your logic here and delete this line

	return &naming.StringRes{}, nil
}
