package gate_server

import (
	"github.com/viphxin/xingo/cluster"
	"github.com/viphxin/xingo/logger"
	_"github.com/viphxin/xingo/clusterserver"
	"github.com/viphxin/xingo/clusterserver"
	"fmt"
	"strings"
	"github.com/viphxin/xingo/iface"
	"xingo_cluster/pb"
	"github.com/golang/protobuf/proto"
)

type Proxy2GameRouter struct {
	cluster.BaseRpcRouter
}

func (this *Proxy2GameRouter)Handle(request iface.IRpcRequest){
	//Json反序列化数字到interface{}类型的值中，默认解析为float64类型，在使用时要注意
	//pid := int32((request.GetArgs()[0]).(float64))
	//gob序列化没问题
	pid := (request.GetArgs()[0]).(int32)
	message := (request.GetArgs()[1]).(string)
	logger.Debug(pid, message)
	//onenet := clusterserver.GlobalClusterServer.ChildsMgr.GetRandomChild("net")
	//if onenet != nil{
	//	onenet.CallChildNotForResult("PushMsg2Client", pid, message)
	//}
	for _, child := range clusterserver.GlobalClusterServer.ChildsMgr.GetChilds(){
		if strings.Contains(child.GetName(), "net"){
			child.CallChildNotForResult("PushMsg2Client", pid, message)
		}
	}
}

type AddRouter struct {
	cluster.BaseRpcRouter
}

func (this *AddRouter)Handle(request iface.IRpcRequest){
	//Json反序列化数字到interface{}类型的值中，默认解析为float64类型，在使用时要注意
	i := (request.GetArgs()[0]).(int)
	ii := (request.GetArgs()[1]).(int)

	logger.Debug(fmt.Sprintf("%d + %d = %d", i, ii, i + ii))
	request.PushReturn("sum", i + ii)
}

type GetGSTimeRouter struct {
	cluster.BaseRpcRouter
}
func (this *GetGSTimeRouter)Handle(request iface.IRpcRequest){
	//Json反序列化数字到interface{}类型的值中，默认解析为float64类型，在使用时要注意
	pid := (request.GetArgs()[0]).(int32)
	onenet := clusterserver.GlobalClusterServer.ChildsMgr.GetRandomChild("admin")
	logger.Debug(onenet)
	if onenet != nil{
		onenet.CallChildNotForResult("GetGSTime", pid)
	}
}

type BytesCalcRouter struct {
	cluster.BaseRpcRouter
}
func (this *BytesCalcRouter)Handle(request iface.IRpcRequest){
	data := (request.GetArgs()[0]).([]byte)
	addcalc := &pb.AddCalc{}
	err := proto.Unmarshal(data, addcalc)
	if err != nil{
		logger.Error("BytesCalcRouter Handle Err: ", err.Error())
		return
	}
	addcalc.Result = addcalc.A + addcalc.B
	bdata, err := proto.Marshal(addcalc)
	if err != nil{
		logger.Error(err.Error())
		return
	}
	request.PushReturn("ret", bdata)
}