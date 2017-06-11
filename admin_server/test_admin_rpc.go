package admin_server

import (
	"github.com/viphxin/xingo/clusterserver"
	"github.com/viphxin/xingo/cluster"
	"github.com/viphxin/xingo/logger"
)

type TestAdminRpc struct {

}

func (this *TestAdminRpc)GetGSTime(request *cluster.RpcRequest){
	pid := (request.Rpcdata.Args[0]).(float64)
	//转发到gate
	onegate := clusterserver.GlobalClusterServer.RemoteNodesMgr.GetRandomChild("gate")

	if onegate != nil{
		logger.Debug("chose root: " + onegate.GetName())
		onegate.CallChildNotForResult("Proxy2Game", pid, "2222.3333")
	}
}
