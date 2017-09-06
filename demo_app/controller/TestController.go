package controller

import (
	"demo/demo_app/config"
	"fmt"
)

type TestController struct {

}

func (this *TestController)Handle() (err error) {
	r := "Hello World " + config.Version
	fmt.Println(r)
	return
}
