package main

import (
	"path/filepath"
	"github.com/viphxin/xingo"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err == nil{
		s := xingo.NewXingoMater(filepath.Join(dir, "conf", "clusterconf.json"))
		s.StartMaster()
	}
}