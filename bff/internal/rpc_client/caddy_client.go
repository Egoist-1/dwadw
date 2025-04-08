package rpc_client

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	cl "start/naming/pb/naming"
)

func InitCaddyClient(path string) cl.NamingClient {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad(path, &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	client := cl.NewNamingClient(conn.Conn())
	return client
}
