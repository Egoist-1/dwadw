package naminglogic

import (
	"context"

	"naming/rpc/internal/svc"
	"naming/rpc/naming"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNamingInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNamingInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNamingInfoLogic {
	return &GetNamingInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询示例详情
func (l *GetNamingInfoLogic) GetNamingInfo(in *naming.IdReq) (*naming.NamingInfoRes, error) {
	// todo: add your logic here and delete this line

	return &naming.NamingInfoRes{}, nil
}
