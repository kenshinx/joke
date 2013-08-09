package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/kenshinx/joke/controllers"
)

/*
Joke: The web console of godns
*/

const (
	Appname = "Joke"
	Author  = "kenshin"
	Version = "0.1.1"
)

var (
	Addr          string
	Port          int
	Runmode       string
	Redis         string
	RedisDB       int
	RedisPassword string
	BasicAuth     string
	LogFile       string
	LogRotate     bool
)

func main() {
	fmt.Printf("%s begining running on %s:%d \n", Appname, Addr, Port)
	initServer()

	beego.Router("/", &controllers.ZoneController{})
	beego.Run()
}

func initServer() {
	beego.AppName = Appname
	beego.HttpAddr = Addr
	beego.HttpPort = Port
	beego.RunMode = Runmode
	beego.AutoRender = true
	beego.ViewsPath = "views"
	beego.SetStaticPath("/static", "static")

	beego.AppConfig.SetValue("basic_auth", BasicAuth)

}

func init() {
	flag.StringVar(&Addr, "addr", "127.0.0.1", "run on the given address")
	flag.IntVar(&Port, "port", 8008, "run on the given port")
	flag.StringVar(&Runmode, "runmode", "dev", "running mode: dev|prof")
	flag.StringVar(&Redis, "redis_addr", "127.0.0.1:6379", "redis server that the zone records storage.")
	flag.IntVar(&RedisDB, "redis_db", 0, "redis server db")
	flag.StringVar(&RedisPassword, "redis_password", "", "redis server password")
	flag.StringVar(&BasicAuth, "basic_auth", "", "colon separated user-password to enable")
	flag.StringVar(&LogFile, "log", "", "log file")
	flag.BoolVar(&LogRotate, "log_rotate", true, "rotate log")

	flag.Parse()
}
