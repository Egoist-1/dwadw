package logic

import (
	"context"
	"start/naming/pb/naming"
	"start/pkg/e"

	"start/bff/internal/svc"
	"start/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHostLogic {
	return &UpdateHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHostLogic) UpdateHost(req *types.UpdateHostReq) (resp *types.UpdateHostRes, err error) {
	_, err = l.svcCtx.CaddyClient.UpdateHost(l.ctx, &naming.UpdateHostReq{
		OrgHost: req.OrgHost,
		ExpHost: req.ExpHost,
	})
	switch err.(type) {
	case error:
		return nil, err
	case e.Err:
		e := err.(e.Err)
		return &types.UpdateHostRes{
			Res: types.Res{
				Code: 200,
				Msg:  e.Code().String(),
			},
		}, err
	}
	return &types.UpdateHostRes{
		Res: types.Res{
			Code: 200,
		},
	}, nil
}
