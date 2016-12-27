package main

import (
	"github.com/viphxin/xingo/clusterserver"
	"github.com/viphxin/xingo/sys_rpc"
	"os"
	"path/filepath"
	"xingo_cluster/net_server"
	"xingo_cluster/gate_server"
)

func main() {
	args := os.Args
	dir, err := filepath.Abs(filepath.Dir("."))
	if err == nil{
		s := clusterserver.NewClusterServer(args[1], filepath.Join(dir, "conf", "clusterconf.json"))
		s.AddRpcRouter(&sys_rpc.ChildRpc{})
		s.AddRpcRouter(&sys_rpc.RootRpc{})
		/*
		注册分布式服务器
		*/
		//net server
		s.AddModule("net", &net_server.TestNetApi{}, &net_server.TestNetRpc{})
		s.AddModule("gate", nil, &gate_server.TestGateRpc{})

		s.StartClusterServer()
	}
}