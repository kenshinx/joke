package controllers

import (
	"github.com/astaxie/beego"
	auth "github.com/kenshinx/go-http-auth"
	"strings"
)

func Secret(user, realm string) string {
	s := strings.SplitN(beego.AppConfig.String("basic_auth"), ":", 2)
	if len(s) < 2 {
		panic("basic auth: " + beego.AppConfig.String("basic_auth") + " formatting error")
	}
	if user == s[0] {
		return auth.GenSHAPassword(s[1])
	}
	return ""
}

func CheckAuth(ctx *beego.Context) {
	if beego.AppConfig.String("basic_auth") != "" {
		a := auth.NewBasicAuthenticator("joke.sina", Secret)
		if username := a.CheckAuth(ctx.Request); username == "" {
			a.RequireAuth(ctx.ResponseWriter, ctx.Request)
		}
	}
}
