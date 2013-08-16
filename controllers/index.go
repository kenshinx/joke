package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Redirect("/dns", 302)
}
