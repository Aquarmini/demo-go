package main

import "fmt"
import (
	"os"
	"demo/demo_app/controller"
)

func main() {
	var params []string
	if len(os.Args) <= 1 {
		fmt.Println("请输入方法名")
		return
	}

	TaskMap := make(map[string]controller.ControllerInterface)

	TaskMap["Test"] = &controller.TestController{}

	for i := 1; i < len(os.Args); i++ {
		params = append(params, os.Args[i])
	}

	fmt.Println(TaskMap)

	err := TaskMap["Test"].Handle()
	if err != nil {
		fmt.Println(err.Error())
	}
}