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
	initLogger()

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/dns", &controllers.DNSController{})
	beego.Router("/dns/del", &controllers.DNSDelController{})
	beego.Run()
}

func initLogger() {
	console, _ := beego.AppConfig.Bool("stdout")
	if !console {
		beego.BeeLogger.DelLogger("console")
	}

	if beego.AppConfig.String("logfile") != "" {
		cfg := fmt.Sprintf(`{"filename":"%s"}`, beego.AppConfig.String("logfile"))
		beego.BeeLogger.SetLogger("file", cfg)
	}

	beego.BeeLogger.SetLevel(beego.LevelDebug)
	beego.BeeLogger.EnableFuncCallDepth(false)
}
