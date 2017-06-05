package net_server

import (
	"github.com/viphxin/xingo/fnet"
	"github.com/golang/protobuf/proto"
	"fmt"
	"xingo_cluster/pb"
	"github.com/viphxin/xingo/logger"
	"github.com/viphxin/xingo/clusterserver"
	"math/rand"
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

func (this *TestNetApi)Api_10(request *fnet.PkgAll){
	//test rpc for result
	//转发到gate
	onegate := clusterserver.GlobalClusterServer.RemoteNodesMgr.GetRandomChild("gate")

	if onegate != nil{
		logger.Info("chose root: " + onegate.GetName())
		i := rand.Intn(10)
		ii := rand.Intn(10)
		response, err := onegate.CallChildForResult("Add", i, ii)
		if err == nil{
			pid, _ := request.Fconn.GetProperty("pid")
			p := GlobalPlayerMgr.GetPlayer(pid.(int32))
			if p != nil{
				msg := &pb.BroadCast{
					Pid : pid.(int32),
					Tp: 1,
					Data: &pb.BroadCast_Content{
						Content: fmt.Sprintf("%d + %d = %f", i, ii, response.Result["sum"].(float64)),
					},
				}
				p.SendMsg(10, msg)
			}
		}else{
			logger.Error(err)
		}
	}
}

func (this *TestNetApi)Api_11(request *fnet.PkgAll){
	//test rpc for result
	//转发到gate
	onegate := clusterserver.GlobalClusterServer.RemoteNodesMgr.GetRandomChild("gate")

	if onegate != nil{
		pid, _ := request.Fconn.GetProperty("pid")
		onegate.CallChildNotForResult("GetGSTime", pid)

	}
}
