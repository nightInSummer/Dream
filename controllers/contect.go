package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type ContectController struct {
	beego.Controller
}

func (this *ContectController) Get() {
	req := httplib.Get("http://acg.178.com/")
	str, err := req.String()
	if err != nil {
	}
	this.Ctx.WriteString(str)
}
