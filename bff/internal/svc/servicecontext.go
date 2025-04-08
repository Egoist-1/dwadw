package svc

import (
	"start/bff/internal/config"
	"start/bff/internal/rpc_client"
	"start/naming/pb/naming"
)

type ServiceContext struct {
	Config      config.Config
	CaddyClient naming.NamingClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		CaddyClient: rpc_client.InitCaddyClient(c.Caddy),
	}
}
