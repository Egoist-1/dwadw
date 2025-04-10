package naminglogic

import (
	"context"

	"naming/rpc/internal/svc"
	"naming/rpc/naming"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNamingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNamingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNamingLogic {
	return &UpdateNamingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改示例
func (l *UpdateNamingLogic) UpdateNaming(in *naming.NamingInfoRes) (*naming.StringRes, error) {
	// todo: add your logic here and delete this line

	return &naming.StringRes{}, nil
}
