package svc

import (
	"naming/naming/internal/config"
	"naming/naming/internal/ioc"
	"naming/naming/internal/logic/caddy"
	"naming/naming/internal/repo"
	"naming/naming/internal/repo/dao"
)

type ServiceContext struct {
	Config config.Config
	Repo   repo.NamingRepo
	Caddy  caddy.Caddy
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ioc.InitGorm(c.DSN)
	dao := dao.NewNamingGORM(db)
	r := repo.NewNamingRepo(dao)
	ca := caddy.NewCaddy(c.CaddyHost)
	return &ServiceContext{
		Config: c,
		Repo:   r,
		Caddy:  ca,
	}
}
