package logic

import (
	"context"
	"naming/naming/internal/svc"
	"naming/naming/pb/naming"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoadLogic {
	return &LoadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoadLogic) Load(in *naming.LoadReq) (*naming.LoadRes, error) {
	//调用 caddy的load host/config
	err := l.svcCtx.Caddy.Load(l.ctx, l.Logger, in.Cfg)
	if err != nil {
		return nil, err
	}
	//获取caddy的host /apps/..../host
	hosts, err := l.svcCtx.Caddy.GetHostInfo(l.ctx, l.Logger)
	//拿到host 存入 数据库
	err = l.svcCtx.Repo.StoreHost(l.ctx, hosts)
	return &naming.LoadRes{}, err
}
