package net_server

import (
	"net/http"
	"net/url"
	"fmt"
	"github.com/viphxin/xingo/fnet"
)

type HelloHttpRouter struct {
	fnet.BaseHttpRouter
}

func (this *HelloHttpRouter)Handle(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err == nil && len(queryForm["name"]) > 0 {
		w.Write([]byte(fmt.Sprintf("hello %s", queryForm["name"][0])))
	}else{
		w.Write([]byte("hello 陌生人"))
	}

}
