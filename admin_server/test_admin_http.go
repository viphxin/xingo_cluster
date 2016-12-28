package admin_server

import (
	"net/http"
	"github.com/viphxin/xingo/logger"
	"strings"
	"time"
)

type TestAdminHttp struct {

}

func (this *TestAdminHttp)Hello(w http.ResponseWriter, r *http.Request) {
	logger.Info(strings.Repeat("hello", 10))
	time.Sleep(3*time.Second)
	w.Write([]byte("hudfasdkasfas"))
}
