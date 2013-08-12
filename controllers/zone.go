package controllers

import (
	"github.com/astaxie/beego"
)

type ZoneController struct {
	beego.Controller
}

// http basic auth
func (c *ZoneController) Prepare() {
	CheckAuth(c.Ctx)
}

func (c *ZoneController) Get() {
	c.Data["Website"] = "kenshinx.me"
	c.Data["Email"] = "kenshin"
	c.TplNames = "index.html"
}
