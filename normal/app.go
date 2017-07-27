package main

import (
	"fmt"
	"demo/normal/model"
)

func main() {
	fmt.Printf("Hello World %s\n", "limx");
	var user *model.User = &model.User{Id:1, Name:"limx"};
	var msg string = user.SayHello();
	fmt.Printf(msg + "\n");

}