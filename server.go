package main

import (
	"os"
	"path/filepath"
	"xingo_cluster/net_server"
	"xingo_cluster/gate_server"
	"xingo_cluster/admin_server"
        _ "net/http"
	_ "net/http/pprof"
	"github.com/viphxin/xingo"
	"xingo_cluster/game_server"
)

func main() {
	//pprof
	//go func() {
	//	println(http.ListenAndServe("localhost:6060", nil))
	//}()

	//server code
	args := os.Args
	dir, err := filepath.Abs(filepath.Dir("."))
	if err == nil{
		if true{
			s := xingo.NewXingoCluterServer(args[1], filepath.Join(dir, "conf", "clusterconf.json"))
			/*
			注册分布式服务器
			*/
			//net server
			s.AddModule("net", &net_server.TestNetApi{}, nil, &net_server.TestNetRpc{})
			//gate server
			s.AddModule("gate", nil, nil, &gate_server.TestGateRpc{})
			//admin server
			s.AddModule("admin", nil, &admin_server.TestAdminHttp{}, &admin_server.TestAdminRpc{})

			s.StartClusterServer()
		}else{
			s := xingo.NewXingoCluterServer(args[1], filepath.Join(dir, "conf", "clusterconf_测试网关有root和http.json"))
			/*
			注册分布式服务器
			*/
			//net server
			s.AddModule("net", &net_server.TestNetApi2{}, &net_server.TestNetHttp{}, &net_server.TestNetRpc{})
			//game server
			s.AddModule("game", nil, nil, &game_server.TestGameRpc{})

			s.StartClusterServer()
		}

	}
}
