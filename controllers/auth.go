package controllers

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"github.com/astaxie/beego"
	"strings"
)

func GenSHAPassword(plaintext string) string {
	h := sha1.New()
	h.Write([]byte(plaintext))
	bs := h.Sum(nil)
	passwd := base64.StdEncoding.EncodeToString(bs)
	return "{SHA}" + passwd

}

func Secret(user, realm string) string {
	s := strings.SplitN(beego.AppConfig.String("basic_auth"), ":", 2)
	if user == s[0] {
		return GenMD5Password(s[1])
	}
	return ""
}
