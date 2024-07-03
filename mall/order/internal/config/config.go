package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//去etcd拿注册地址
	UserRpc zrpc.RpcClientConf
}
