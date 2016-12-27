package net_server

import (
	"github.com/viphxin/xingo/fnet"
	"github.com/golang/protobuf/proto"
	"fmt"
	"xingo_cluster/pb"
	"github.com/viphxin/xingo/logger"
	"github.com/viphxin/xingo/clusterserver"
)

type TestNetApi struct {

}

func (this *TestNetApi)Api_0(request *fnet.PkgAll){
	msg := &pb.Talk{}
	err := proto.Unmarshal(request.Pdata.Data, msg)
	if err == nil {
		logger.Debug(fmt.Sprintf("user talk: content: %s.", msg.Content))
		pid, err1 := request.Fconn.GetProperty("pid")
		if err1 == nil{
			//转发到gate
			onegate := clusterserver.GlobalClusterServer.RemoteNodesMgr.GetRandomChild("gate")

			if onegate != nil{
				logger.Info("chose root: " + onegate.GetName())
				onegate.CallChildNotForResult("Proxy2Game", pid.(int32), msg.Content)
			}
		}else{
			logger.Error(err1)
			request.Fconn.LostConnection()
		}

	} else {
		logger.Error(err)
		request.Fconn.LostConnection()
	}
}
