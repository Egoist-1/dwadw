package logic

import (
	"context"

	"naming/internal/svc"
	"naming/pb/naming"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHostLogic {
	return &UpdateHostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateHostLogic) UpdateHost(in *naming.UpdateHostReq) (*naming.UpdateHostRes, error) {
	//通过org 找到在caddy的下标
	id, err := l.svcCtx.Repo.FindByHost(l.ctx, in.OrgHost)
	if err != nil {
		return nil, err
	}
	//host的下标
	index := id - 1
	err = l.svcCtx.Caddy.UpdateHost(l.ctx, index, in.ExpHost)
	if err != nil {
		l.Error("更新caddy的host 失败", err, in.OrgHost, in.ExpHost)
		return nil, err
	}
	//更新数据库的host
	err = l.svcCtx.Repo.UpdateHost(l.ctx, in.OrgHost, in.ExpHost)
	if err != nil {
		return nil, err
	}
	return &naming.UpdateHostRes{}, nil
}
