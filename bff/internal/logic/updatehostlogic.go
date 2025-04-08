package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
