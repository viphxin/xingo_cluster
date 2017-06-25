package admin_server

import (
	"net/http"
	"github.com/viphxin/xingo/logger"
	"strings"
	"time"
	"github.com/viphxin/xingo/fnet"
)

type HelloHttpRouter struct {
	fnet.BaseHttpRouter
}

func (this *HelloHttpRouter)Handle(w http.ResponseWriter, r *http.Request) {
	logger.Info(strings.Repeat("hello", 10))
	time.Sleep(3*time.Second)
	w.Write([]byte("hudfasdkasfas"))
}
