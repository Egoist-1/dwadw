package logic

import (
	"context"
	"start/naming/pb/naming"

	"github.com/zeromicro/go-zero/core/logx"
	"start/bff/internal/svc"
	"start/bff/internal/types"
	"start/pkg/e"
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
	switch err.(type) {
	case error:
		return nil, err
	case e.Err:
		e := err.(e.Err)
		return &types.AddHostRes{
			Res: types.Res{
				Code: 200,
				Msg:  e.Code().String(),
			},
		}, err
	}
	return &types.AddHostRes{
		Res: types.Res{
			Code: 200,
		},
	}, nil
}
