package logic

import (
	"context"
	"naming/internal/svc"
	"naming/pb/naming"

	"github.com/zeromicro/go-zero/core/logx"
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
	for _, val := range in.Host {
		err = l.svcCtx.Caddy.AddHost(l.ctx, l.Logger, val)
		if err != nil {
			return nil, err
		}
	}
	//存储到数据库
	err = l.svcCtx.Repo.StoreHost(l.ctx, in.Host)
	return &naming.AddHostRes{}, err
}
