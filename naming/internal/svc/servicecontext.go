package svc

import (
	"naming/internal/config"
	"naming/internal/ioc"
	"naming/internal/logic/caddy"
	"naming/internal/repo"
	"naming/internal/repo/dao"
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
