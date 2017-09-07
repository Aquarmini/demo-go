package controller

import (
	log "github.com/sirupsen/logrus"
	"fmt"
	"github.com/limingxinleo/di"
	"app/providers"
)

type LogController struct {
	DI di.Context
}

func (this *LogController)Handle() (err error) {
	//logger := this.DI.Get("logger").(*providers.Logger)
	config := this.DI.Get("config").(*providers.Config)
	version, _ := config.GetKey("application", "version")
	log.Info(fmt.Sprintf("Current:version=%s", version.Value()));
	fmt.Println("查看日志")
	return nil
}
