package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"naming/naming/internal/svc"
	"naming/naming/pb/naming"
)

type AddHostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHostLogic {
	return &AddHostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddHostLogic) AddHost(in *naming.AddHostReq) (*naming.AddHostRes, error) {
	//遍历host 插入数据
	var err error
	err = l.svcCtx.Caddy.AddHost(l.ctx, l.Logger, in.Host)

	//存储到数据库
	err = l.svcCtx.Repo.StoreHost(l.ctx, in.Host)
	return &naming.AddHostRes{}, err
}
