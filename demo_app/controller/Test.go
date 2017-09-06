package controller

import "demo/demo_app/config"

func Say() (r string, err error) {
	r = "Hello World " + config.Version
	return
}
