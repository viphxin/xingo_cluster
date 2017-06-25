package game_server

import (
	"github.com/viphxin/xingo/cluster"
	"github.com/viphxin/xingo/logger"
	_"github.com/viphxin/xingo/clusterserver"
	"github.com/viphxin/xingo/clusterserver"
	"fmt"
	"strings"
	"github.com/viphxin/xingo/iface"
)

type Proxy2GameRouter struct {
	cluster.BaseRpcRouter
}

func (this *Proxy2GameRouter)Handle(request iface.IRpcRequest){
	//Json反序列化数字到interface{}类型的值中，默认解析为float64类型，在使用时要注意
	pid := int32((request.GetArgs()[0]).(float64))
	message := (request.GetArgs()[1]).(string)
	logger.Debug(pid, message)
	//onenet := clusterserver.GlobalClusterServer.ChildsMgr.GetRandomChild("net")
	//if onenet != nil{
	//	onenet.CallChildNotForResult("PushMsg2Client", pid, message)
	//}
	for _, child := range clusterserver.GlobalClusterServer.RemoteNodesMgr.GetChilds(){
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
	i := int32((request.GetArgs()[0]).(float64))
	ii := int32((request.GetArgs()[1]).(float64))

	logger.Debug(fmt.Sprintf("%d + %d = %d", i, ii, i + ii))
	request.PushReturn("sum",i + ii)
}