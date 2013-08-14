package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
	"log"
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

type Host struct {
	Domain string `form:"domain"`
	IP     string `form:"ip"`
}

type DNSController struct {
	beego.Controller
	rc *redis.Client
}

// http basic auth
// init redis connect
func (c *DNSController) Prepare() {
	CheckAuth(c.Ctx)
	c.rc = initialRedis()
}

func (c *DNSController) Get() {
	var HostsRecord = make(map[string]string)
	bindkey := beego.AppConfig.String("bindkey")
	err := c.rc.Hgetall(bindkey, HostsRecord)
	if err != nil {
		panic(err)
	}
	c.Data["Hosts"] = HostsRecord
	c.Layout = "layout.html"
	c.TplNames = "dns.html"
	c.Render()
}

func (c *DNSController) Post() {
	h := new(Host)
	if err := c.ParseForm(h); err != nil {
		c.Ctx.Abort(400, "Invalid post data")
	}
	if h.Domain == "" || h.IP == "" {
		c.Ctx.Abort(400, "Both domain and iP can't be empty")
	}
	bindkey := beego.AppConfig.String("bindkey")
	if ok, err := c.rc.Hset(bindkey, h.Domain, []byte(h.IP)); !ok {
		c.Ctx.Abort(500, "Save hosts record failed")
		log.Println(err)
	}
	log.Printf("[%s:%s] insert into redis", h.Domain, h.IP)

}
