package admin_server

import (
	"github.com/viphxin/xingo/clusterserver"
	"github.com/viphxin/xingo/cluster"
	"github.com/viphxin/xingo/logger"
	"github.com/viphxin/xingo/iface"
)

type GetGSTimeRouter struct {
	cluster.BaseRpcRouter
}

func (this *GetGSTimeRouter)Handle(request iface.IRpcRequest){
	pid := (request.GetArgs()[0]).(int32)
	//转发到gate
	onegate := clusterserver.GlobalClusterServer.RemoteNodesMgr.GetRandomChild("gate")

	if onegate != nil{
		logger.Debug("chose root: " + onegate.GetName())
		onegate.CallChildNotForResult("Proxy2Game", pid, "2222.3333")
	}
}
