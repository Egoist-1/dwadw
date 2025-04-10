package naminglogic

import (
	"context"

	"naming/rpc/internal/svc"
	"naming/rpc/naming"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNamingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNamingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNamingListLogic {
	return &GetNamingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询示例列表
func (l *GetNamingListLogic) GetNamingList(in *naming.NamingReq) (*naming.NamingListRes, error) {
	// todo: add your logic here and delete this line

	return &naming.NamingListRes{}, nil
}
