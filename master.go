package main

import (
	"path/filepath"
	"github.com/viphxin/xingo"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err == nil{
		if true{
			s := xingo.NewXingoMaster(filepath.Join(dir, "conf", "clusterconf.json"))
			s.StartMaster()
		}else{
			s := xingo.NewXingoMaster(filepath.Join(dir, "conf", "clusterconf_测试网关有root和http.json"))
			s.StartMaster()
		}
	}
}
