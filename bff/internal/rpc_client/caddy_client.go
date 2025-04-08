package rpc_client

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	cl "naming/naming/pb/naming"
)

func InitCaddyClient() cl.NamingClient {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/naming.yaml", &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	client := cl.NewNamingClient(conn.Conn())
	return client
}
