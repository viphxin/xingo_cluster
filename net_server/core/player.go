package core

import (
	"github.com/golang/protobuf/proto"
	"github.com/viphxin/xingo/iface"
	"github.com/viphxin/xingo/logger"
	"github.com/viphxin/xingo/utils"
)

type Player struct {
	Pid   int32
	Fconn iface.Iconnection
}

func NewPlayer(pid int32, fconn iface.Iconnection) *Player {
	return &Player{
		Pid:   pid,
		Fconn: fconn,
	}
}

func (this *Player) SendMsg(msgId uint32, data proto.Message) {
	if this.Fconn != nil {
		packdata, err := utils.GlobalObject.Protoc.GetDataPack().Pack(msgId, data)
		if err == nil {
			this.Fconn.Send(packdata)
		} else {
			logger.Error("pack data error")
		}
	}
}
