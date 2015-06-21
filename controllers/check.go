package controllers

import (
	"Dream/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type CheckController struct {
	beego.Controller
}

func (this *CheckController) Login() {
	this.TplNames = "index/index.html"
}

func (this *CheckController) PostLogin() {
	type LoginForm struct {
		Account  string `form:"username"`
		Password string `form:"password"`
	}
	form := LoginForm{}
	if err := this.ParseForm(&form); err != nil {
		this.StopRun()
	}
	member := &models.Member{}
	statu := member.CheckPass(form.Account, form.Password)
	if statu == 0 {
		//返回值为0账号处于冻结状态
		this.Ctx.Redirect(302, this.UrlFor("IndexController.Get")+"?get="+"账号已锁定")
	}
	if statu == 1 {
		//返回值为1密码正确
		this.Ctx.SetCookie("WOKUNAME", form.Account, 1577836800)
		this.Ctx.Redirect(302, this.UrlFor("IndexController.Get"))
	} else {
		if statu == 2 {
			//返回值为2用户名或密码不正确
			num := member.Find(form.Account)
			con := int(num)
			str := strconv.Itoa(con)
			this.Ctx.Redirect(302, this.UrlFor("IndexController.Get")+"?get="+"用户名或密码错误，还剩"+str+"次机会")
		}
	}

}

func (this *CheckController) Reg() {
	this.TplNames = "index/reg.html"
}

func (this *CheckController) PostReg() {
	member := &models.Member{}
	member.Nickname = this.GetString("nickname")
	member.Password = this.GetString("pwd")
	member.Email = this.GetString("email")
	member.Status = "0"
	member.Error = 5
	member.Insert()
	//发送激活邮件
	SendEmail([]string{member.Email}, member.Nickname+"：您的账号申请成功，请点击链接激活！"+time.Now().String(), `
		您好：`+member.Nickname+`<br><br>
		（请在一小时内完成）您需要点击以下链接来激活您的账户：<br><br>
		http://`+beego.AppConfig.String("httpWebSite")+`/register.html`)
	member.Type = beego.AppConfig.String("EmailPassword")
	this.Ctx.WriteString(member.Email)
}

func (this *CheckController) Out() {
	this.Ctx.SetCookie("WOKUNAME", "null")
	this.Ctx.Redirect(302, this.UrlFor("IndexController.Get"))
}
