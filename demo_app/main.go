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

	TaskMap := make(map[string]func() (string, error))

	TaskMap["Say"] = controller.Handle

	for i := 1; i < len(os.Args); i++ {
		params = append(params, os.Args[i])
	}

	fmt.Println(TaskMap)

	result, _ := TaskMap["Say"]()
	fmt.Println(result)
}
