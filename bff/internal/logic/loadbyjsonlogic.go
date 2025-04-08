package logic

import (
	"context"
	"start/naming/pb/naming"
	"start/pkg/e"

	"start/bff/internal/svc"
	"start/bff/internal/types"

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
	switch err.(type) {
	case error:
		return nil, err
	case e.Err:
		e := err.(e.Err)
		return &types.LoadByjsonRes{
			Res: types.Res{
				Code: 200,
				Msg:  e.Code().String(),
			},
		}, err
	}
	return &types.LoadByjsonRes{
		Res: types.Res{
			Code: 200,
		},
	}, nil
}
