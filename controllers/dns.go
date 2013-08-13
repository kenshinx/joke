package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
)

func initialRedis() *redis.Client {
	db, _ := beego.AppConfig.Int("redisdb")
	rc := &redis.Client{
		Addr:     beego.AppConfig.String("redisaddr"),
		Db:       db,
		Password: beego.AppConfig.String("redispassword"),
	}
	return rc
}

type DNSController struct {
	beego.Controller
	rc *redis.Client
}

// http basic auth
// redis connect
func (c *DNSController) Prepare() {
	CheckAuth(c.Ctx)
	c.rc = initialRedis()
}

func (c *DNSController) Get() {
	var DNSRecords = make(map[string]string)
	bindkey := beego.AppConfig.String("bindkey")
	err := c.rc.Hgetall(bindkey, DNSRecords)
	if err != nil {
		panic(err)
	}
	fmt.Println(DNSRecords)
	c.Data["DNS"] = DNSRecords
	c.Layout = "layout.html"
	c.TplNames = "dns.html"
}
