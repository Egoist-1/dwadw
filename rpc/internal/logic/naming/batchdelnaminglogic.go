package naminglogic

import (
	"context"

	"naming/rpc/internal/svc"
	"naming/rpc/naming"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDelNamingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDelNamingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDelNamingLogic {
	return &BatchDelNamingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除示例
func (l *BatchDelNamingLogic) BatchDelNaming(in *naming.IdListReq) (*naming.StringRes, error) {
	// todo: add your logic here and delete this line

	return &naming.StringRes{}, nil
}
