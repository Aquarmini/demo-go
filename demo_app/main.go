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

	TaskMap["Say"] = controller.Say

	for i := 1; i < len(os.Args); i++ {
		params = append(params, os.Args[i])
	}

	fmt.Println(TaskMap)

	result, _ := TaskMap["Say"]()
	fmt.Println(result)
	//return
	//task()
	//for i := 1; i < len(os.Args); i++ {
	//	fn = append(fn, os.Args[i])
	//	fmt.Println(os.Args[i])
	//}

	//symbol := []func() string{test1, test2, test3}
	//for _, v := range symbol {
	//	fmt.Println(v())
	//}
	//fmt.Println(fn)
	//fmt.Println(config.Version)
}
