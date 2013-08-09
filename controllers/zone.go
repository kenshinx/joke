package controllers

import (
	auth "github.com/abbot/go-http-auth"
	"github.com/astaxie/beego"
)

type ZoneController struct {
	beego.Controller
}

func (c *ZoneController) Prepare() {
	if beego.AppConfig.String("basic_auth") != "" {
		a := auth.NewBasicAuthenticator("joke.sina", Secret)
		if username := a.CheckAuth(c.Ctx.Request); username == "" {
			a.RequireAuth(c.Ctx.ResponseWriter, c.Ctx.Request)
		}
	}

}

func (c *ZoneController) Get() {
	c.Data["Website"] = "kenshinx.me"
	c.Data["Email"] = "kenshin"
	c.TplNames = "index.html"
}
