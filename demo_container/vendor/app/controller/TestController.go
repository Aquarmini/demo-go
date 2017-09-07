package controller

import (
	"fmt"
	"github.com/limingxinleo/di"
	"app/providers"
)

type TestController struct {
	DI di.Context
}

func (this *TestController)Handle() (err error) {
	config := this.DI.Get("config").(*providers.Config)
	version, _ := config.GetKey("application", "version")
	r := "Hello World " + version.Value()
	fmt.Println(r)

	return nil
}
