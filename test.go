package main

import (
	"reflect"
	"github.com/viphxin/xingo/cluster"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	type RpcData struct {
		MsgType cluster.RpcSignal `json:"msgtype"`
		Key string `json:"key,omitempty"`
		Target string `json:"target,omitempty"`
		Args []interface{} `json:"args,omitempty"`
		Result []reflect.Value `json:"result,omitempty"`
	}

	a := &RpcData{
		MsgType: cluster.REQUEST_NORESULT,
		Key: "dasdasdsa",
	}
	bbb, err := json.Marshal(a)
	if err != nil{
		fmt.Println("error")
		fmt.Println(err)
	}else{
		aa := &RpcData{}
		err = json.Unmarshal(bbb, aa)
		if err != nil{
			println(err)
		}else{
			println(aa.Key, aa.MsgType)
		}
	}
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			println("panic recover!!!")
		}
	}()
	for {
		println("sdasdsadas1111111")
		panic("huangxinsss")
		time.Sleep(3*time.Second)
	}
}