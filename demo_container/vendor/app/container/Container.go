package container

import (
	"app/providers"
	"github.com/limingxinleo/di"
	"fmt"
	"github.com/go-ini/ini"
)

var DI di.Context

func GetInstance() di.Context {
	// Create a Builder with the default scopes.
	builder, _ := di.NewBuilder()

	// Define an object in the App scope.
	builder.AddDefinition(di.Definition{
		Name: "config",
		Scope: di.App,
		Build: func(ctx di.Context) (interface{}, error) {
			fmt.Println("NewConfig")
			cfg, _ := ini.InsensitiveLoad("config.ini")
			return &providers.Config{Cfg:cfg}, nil
		},
	})

	// Define an object in the Request scope.
	builder.AddDefinition(di.Definition{
		Name: "logger",
		Scope: di.Request,
		Build: func(ctx di.Context) (interface{}, error) {
			return &providers.Logger{}, nil
		},
	})

	DI = builder.Build()
	fmt.Println("DI Build")
	return DI
}


