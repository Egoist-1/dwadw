package logic

import (
	"context"
	"naming/bff/internal/svc"
	"naming/bff/internal/types"
	"naming/naming/pb/naming"

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
	return &types.UpdateHostRes{
		Res: types.Res{
			Code: 200,
		},
	}, err
}
