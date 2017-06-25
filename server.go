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
	"strings"
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
			if strings.Contains(args[1], "net") {
				s.AddRouter("0", &net_server.Api0Router{})
				s.AddRouter("10", &net_server.Api10Router{})
				s.AddRouter("11", &net_server.Api11Router{})
				s.AddRouter("12", &net_server.Api12Router{})
				s.AddRpcRouter("PushMsg2Client", &net_server.PushMsg2ClientRouter{})
				s.AddHttpRouter("/hello", &net_server.HelloHttpRouter{})
			}
			if strings.Contains(args[1], "gate") {
				s.AddRpcRouter("Proxy2Game", &gate_server.Proxy2GameRouter{})
				s.AddRpcRouter("Add", &gate_server.AddRouter{})
				s.AddRpcRouter("GetGSTime", &gate_server.GetGSTimeRouter{})
				s.AddRpcRouter("BytesCalc", &gate_server.BytesCalcRouter{})
			}
			if strings.Contains(args[1], "admin") {
				s.AddHttpRouter("/hello", &admin_server.HelloHttpRouter{})
				s.AddRpcRouter("GetGSTime", &admin_server.GetGSTimeRouter{})
			}
			s.StartClusterServer()
		}else{
			s := xingo.NewXingoCluterServer(args[1], filepath.Join(dir, "conf", "clusterconf_测试网关有root和http.json"))
			/*
			注册分布式服务器
			*/
			if strings.Contains(args[1], "net") {
				s.AddRouter("0", &net_server.Api0Router测试网关有root和http{})
				s.AddRouter("10", &net_server.Api10Router测试网关有root和http{})
				s.AddRouter("11", &net_server.Api11Router测试网关有root和http{})
				s.AddRpcRouter("PushMsg2Client", &net_server.PushMsg2ClientRouter{})
			}
			if strings.Contains(args[1], "game") {
				s.AddRpcRouter("Add", &game_server.AddRouter{})
				s.AddRpcRouter("Proxy2Game", &game_server.Proxy2GameRouter{})
			}

			s.StartClusterServer()
		}

	}
}
