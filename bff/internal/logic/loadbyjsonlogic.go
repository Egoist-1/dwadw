package logic

import (
	"context"
	"naming/bff/internal/svc"
	"naming/bff/internal/types"
	"naming/naming/pb/naming"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoadByjsonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoadByjsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoadByjsonLogic {
	return &LoadByjsonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoadByjsonLogic) LoadByjson(req *types.LoadByjsonReq) (resp *types.LoadByjsonRes, err error) {
	_, err = l.svcCtx.CaddyClient.Load(l.ctx, &naming.LoadReq{
		Cfg: req.LoadJson,
	})
	return &types.LoadByjsonRes{}, nil
}
