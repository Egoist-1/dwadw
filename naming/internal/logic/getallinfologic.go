package logic

import (
	"context"
	"naming/internal/svc"
	"naming/pb/naming"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllInfoLogic {
	return &GetAllInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllInfoLogic) GetAllInfo(in *naming.GetAllInfoReq) (*naming.GetAllInfoReqRes, error) {
	res, err := l.svcCtx.Caddy.GetAllInfo(l.ctx, l.Logger)
	return &naming.GetAllInfoReqRes{
		Info: res,
	}, err
}
