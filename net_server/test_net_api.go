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

type Api0Router struct {
	fnet.BaseRouter
}

func (this *Api0Router)Handle(request iface.IRequest){
	msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err == nil {
		logger.Debug(fmt.Sprintf("user talk: content: %s.", msg.Content))
		pid, err1 := request.GetConnection().GetProperty("pid")
		if err1 == nil{
			//转发到gate
			onegate := clusterserver.GlobalClusterServer.RemoteNodesMgr.GetRandomChild("gate")

			if onegate != nil{
				logger.Debug("chose root: " + onegate.GetName())
				onegate.CallChildNotForResult("Proxy2Game", pid.(int32), msg.Content)
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

type Api10Router struct {
	fnet.BaseRouter
}

func (this *Api10Router)Handle(request iface.IRequest){
	//test rpc for result
	//转发到gate
	onegate := clusterserver.GlobalClusterServer.RemoteNodesMgr.GetRandomChild("gate")

	if onegate != nil{
		logger.Debug("chose root: " + onegate.GetName())
		i := rand.Intn(10)
		ii := rand.Intn(10)
		response, err := onegate.CallChildForResult("Add", i, ii)
		if err == nil{
			pid, _ := request.GetConnection().GetProperty("pid")
			p := GlobalPlayerMgr.GetPlayer(pid.(int32))
			if p != nil{
				msg := &pb.BroadCast{
					Pid : pid.(int32),
					Tp: 1,
					Data: &pb.BroadCast_Content{
						Content: fmt.Sprintf("%d + %d = %d", i, ii, response.Result["sum"].(int)),
					},
				}
				p.SendMsg(10, msg)
			}
		}else{
			logger.Error(err)
		}
	}
}

type Api11Router struct {
	fnet.BaseRouter
}

func (this *Api11Router)Handle(request iface.IRequest){
	//test rpc for result
	//转发到gate
	onegate := clusterserver.GlobalClusterServer.RemoteNodesMgr.GetRandomChild("gate")

	if onegate != nil{
		pid, _ := request.GetConnection().GetProperty("pid")
		onegate.CallChildNotForResult("GetGSTime", pid)

	}
}

/*
测试[]byte数据
*/
type Api12Router struct {
	fnet.BaseRouter
}

func (this *Api12Router)Handle(request iface.IRequest){
	//test rpc for result
	//转发到gate
	onegate := clusterserver.GlobalClusterServer.RemoteNodesMgr.GetRandomChild("gate")

	if onegate != nil{
		data, err := onegate.CallChildForResult("BytesCalc", request.GetData())
		if err != nil{
			logger.Error("Api12Router Handle Err:", err.Error())
			return
		}
		pid, _ := request.GetConnection().GetProperty("pid")
		p := GlobalPlayerMgr.GetPlayer(pid.(int32))
		if p != nil{
			pdata := &pb.AddCalc{}
			err = proto.Unmarshal(data.Result["ret"].([]byte), pdata)
			if err != nil{
				logger.Error(err.Error())
				return
			}
			p.SendMsg(12, pdata)
		}
	}
}
