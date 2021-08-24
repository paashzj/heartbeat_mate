package main

import (
	"heartbeat_mate/pkg/hb"
	"heartbeat_mate/pkg/util"
	"os"
	"os/signal"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	interrupt := make(chan os.Signal, 1)
	hb.Start()
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}
