package config

import (
	"app/controller"
	"github.com/limingxinleo/di"
)

func GetTaskMap(di di.Context) map[string]controller.ControllerInterface {
	TaskMap := make(map[string]controller.ControllerInterface)

	TaskMap["Test"] = &controller.TestController{DI:di}
	TaskMap["Log"] = &controller.LogController{DI:di}

	return TaskMap
}

