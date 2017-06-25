package net_server

import (
	"github.com/viphxin/xingo/fnet"
	"github.com/golang/protobuf/proto"
	"fmt"
	"xingo_cluster/pb"
	"github.com/viphxin/xingo/logger"
	"github.com/viphxin/xingo/clusterserver"
	"math/rand"
	"github.com/viphxin/xingo/iface"
)

type Api0Router测试网关有root和http struct {
	fnet.BaseRouter
}

func (this *Api0Router测试网关有root和http)Handle(request iface.IRequest){
	msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err == nil {
		logger.Debug(fmt.Sprintf("user talk: content: %s.", msg.Content))
		pid, err1 := request.GetConnection().GetProperty("pid")
		if err1 == nil{
			//转发到game
			onegame := clusterserver.GlobalClusterServer.ChildsMgr.GetRandomChild("game")

			if onegame != nil{
				logger.Debug("chose root: " + onegame.GetName())
				onegame.CallChildNotForResult("Proxy2Game", pid.(int32), msg.Content)
			}
		}else{
			logger.Error(err1)
			request.GetConnection().LostConnection()
		}

	} else {
		logger.Error(err)
		request.GetConnection().LostConnection()
	}
}

type Api10Router测试网关有root和http struct {
	fnet.BaseRouter
}

func (this *Api10Router测试网关有root和http)Handle(request iface.IRequest){
	//test rpc for result
	//转发到gate
	onegame := clusterserver.GlobalClusterServer.ChildsMgr.GetRandomChild("game")

	if onegame != nil{
		logger.Debug("chose child: " + onegame.GetName())
		i := rand.Intn(10)
		ii := rand.Intn(10)
		response, err := onegame.CallChildForResult("Add", i, ii)
		if err == nil{
			pid, _ := request.GetConnection().GetProperty("pid")
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

type Api11Router测试网关有root和http struct {
	fnet.BaseRouter
}

func (this *Api11Router测试网关有root和http)Handle(request iface.IRequest){
	//test rpc for result
	//转发到gate
	onegame := clusterserver.GlobalClusterServer.ChildsMgr.GetRandomChild("game")

	if onegame != nil{
		pid, _ := request.GetConnection().GetProperty("pid")
		onegame.CallChildNotForResult("GetGSTime", pid)

	}
}
