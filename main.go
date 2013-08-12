package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/kenshinx/joke/controllers"
)

/*
Joke: The web console of godns
*/

const (
	Author  = "kenshin"
	Version = "0.1.1"
)

func main() {
	fmt.Printf("%s begining running on %s:%d \n", beego.AppName, beego.HttpAddr, beego.HttpPort)

	beego.Router("/", &controllers.DNSController{})
	beego.Run()
}
