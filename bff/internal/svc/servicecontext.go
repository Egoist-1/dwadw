package svc

import (
	"naming/bff/internal/config"
	"naming/bff/internal/rpc_client"
	"naming/naming/pb/naming"
)

type ServiceContext struct {
	Config      config.Config
	CaddyClient naming.NamingClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		CaddyClient: rpc_client.InitCaddyClient(),
	}
}
