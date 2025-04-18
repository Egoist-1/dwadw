package logic

import (
	"context"
	"naming/naming/pb/naming"

	"github.com/zeromicro/go-zero/core/logx"
	"naming/bff/internal/svc"
	"naming/bff/internal/types"
)

type AddHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHostLogic {
	return &AddHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddHostLogic) AddHost(req *types.AddHostReq) (resp *types.AddHostRes, err error) {
	_, err = l.svcCtx.CaddyClient.AddHost(l.ctx, &naming.AddHostReq{
		Host: req.Host,
	})
	return &types.AddHostRes{}, err
}
