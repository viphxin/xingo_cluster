package net_server

import (
	"github.com/viphxin/xingo/cluster"
	"xingo_cluster/pb"
	"github.com/viphxin/xingo/iface"
)

type PushMsg2ClientRouter struct {
	cluster.BaseRpcRouter
}

func (this *PushMsg2ClientRouter)Handle(request iface.IRpcRequest){
	pid := (request.GetArgs()[0]).(int32)
	message := request.GetArgs()[1].(string)
	p := GlobalPlayerMgr.GetPlayer(pid)
	if p != nil{
		msg := &pb.BroadCast{
			Pid : pid,
			Tp: 1,
			Data: &pb.BroadCast_Content{
				Content: message,
			},
		}
		p.SendMsg(1, msg)
	}
}
