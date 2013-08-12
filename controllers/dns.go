package controllers

import (
	"github.com/astaxie/beego"
)

type DNSController struct {
	beego.Controller
}

// http basic auth
func (c *DNSController) Prepare() {
	CheckAuth(c.Ctx)
}

func (c *DNSController) Get() {
	c.TplNames = "dns.html"
}
