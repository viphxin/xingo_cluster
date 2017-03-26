package net_server

import (
	"github.com/viphxin/xingo/iface"
	"github.com/viphxin/xingo/logger"
	"github.com/viphxin/xingo/utils"
	"xingo_cluster/net_server/core"
)

func DoConnectionMade(fconn iface.Iconnection) {
	logger.Debug("111111111111111111111111")
	p := GlobalPlayerMgr.Add(fconn)
	fconn.SetProperty("pid", p.Pid)
}

func DoConnectionLost(fconn iface.Iconnection) {
	logger.Debug("222222222222222222222222")
	pid, err := fconn.GetProperty("pid")
	if err == nil {
		GlobalPlayerMgr.Remove(pid.(int32))
	}

}

func testGSTimer(args ...interface{}) {
	logger.Info("timer call: ", args[0], "========", args[1])
	utils.GlobalObject.GsTimeScheduel.CreateTimer(3000, testGSTimer, []interface{}{111, 222})
}

var GlobalPlayerMgr *core.PlayerMgr = core.NewPlayerMgr()

func init() {
	utils.GlobalObject.OnConnectioned = DoConnectionMade
	utils.GlobalObject.OnClosed = DoConnectionLost
	utils.GlobalObject.GsTimeScheduel.CreateTimer(3000, testGSTimer, []interface{}{111, 222})
}
