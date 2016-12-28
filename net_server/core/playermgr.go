package core

import (
	"sync"
	"github.com/viphxin/xingo/iface"
)

type PlayerMgr struct {
	sync.RWMutex
	players map[int32]*Player
}

func NewPlayerMgr() *PlayerMgr{
	return &PlayerMgr{
		players: make(map[int32]*Player, 0),
	}
}

func (this *PlayerMgr)Add(fconn iface.Iconnection) *Player{
	this.Lock()
	defer this.Unlock()

	p := NewPlayer(int32(fconn.GetSessionId()), fconn)
	this.players[p.Pid] = p
	return p
}

func (this *PlayerMgr)Remove(pid int32){
	this.Lock()
	defer this.Unlock()

	delete(this.players, pid)
}

func (this *PlayerMgr)GetPlayer(pid int32) *Player{
	this.RLock()
	defer this.RUnlock()

	p, ok := this.players[pid]
	if ok{
		return p
	}else{
		return nil
	}
}
