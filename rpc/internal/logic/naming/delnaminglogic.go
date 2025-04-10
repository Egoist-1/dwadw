package naminglogic

import (
	"context"

	"naming/rpc/internal/svc"
	"naming/rpc/naming"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelNamingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelNamingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelNamingLogic {
	return &DelNamingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除示例
func (l *DelNamingLogic) DelNaming(in *naming.IdReq) (*naming.StringRes, error) {
	// todo: add your logic here and delete this line

	return &naming.StringRes{}, nil
}
