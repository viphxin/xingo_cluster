package main

import (
	"path/filepath"
	"github.com/viphxin/xingo/sys_rpc"
	"github.com/viphxin/xingo/clusterserver"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err == nil{
		s := clusterserver.NewMaster(filepath.Join(dir, "conf", "clusterconf.json"))
		s.AddRpcRouter(&sys_rpc.MasterRpc{})
		s.StartMaster()
	}
}