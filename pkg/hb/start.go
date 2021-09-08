package hb

import (
	"go.uber.org/zap"
	"heartbeat_mate/pkg/path"
	"heartbeat_mate/pkg/util"
)

func Start() {
	startHb()
}

func startHb() {
	err := generateHbFile()
	if err != nil {
		util.Logger().Error("generate heartbeat config file failed ", zap.Error(err))
		return
	}
	err = util.CallScript(path.HbStartScript)
	util.Logger().Error("run start heartbeat scripts failed ", zap.Error(err))
}

func generateHbFile() (err error) {
	// todo
	return nil
}
