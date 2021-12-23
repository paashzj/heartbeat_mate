package hb

import (
	"github.com/paashzj/gutil"
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
	stdout, stderr, err := gutil.CallScript(path.HbStartScript)
	util.Logger().Info("output is ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	util.Logger().Error("run start coredns scripts failed ", zap.Error(err))
}

func generateHbFile() (err error) {
	// todo
	return nil
}
