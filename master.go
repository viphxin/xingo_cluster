package main

import (
	"path/filepath"
	"github.com/viphxin/xingo"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err == nil{
		if true{
			s := xingo.NewXingoMater(filepath.Join(dir, "conf", "clusterconf.json"))
			s.StartMaster()
		}else{
			s := xingo.NewXingoMater(filepath.Join(dir, "conf", "clusterconf_测试网关有root和http.json"))
			s.StartMaster()
		}
	}
}