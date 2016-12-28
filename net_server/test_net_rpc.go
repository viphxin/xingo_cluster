package net_server

import (
	"github.com/viphxin/xingo/cluster"
	"xingo_cluster/pb"
)

type TestNetRpc struct {

}

func (this *TestNetRpc)PushMsg2Client(request *cluster.RpcRequest){
	pid := int32(request.Rpcdata.Args[0].(float64))
	message := request.Rpcdata.Args[1].(string)
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
