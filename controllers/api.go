package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
	"os"
)

/* ----------此控制器定义全部第三方API操作---------- */

type ApiController struct {
	beego.Controller
}

var (
	ApiLog *logs.BeeLogger //打印日志
	mail   *utils.Email    //邮件
)

func init() {
	//如果没有日志目录则创建日志目录
	_, err := os.Open("log")
	if err != nil && os.IsNotExist(err) {
		os.Mkdir("log", 0777)
	}
	//初始化日志
	ApiLog = logs.NewLogger(10000)
	ApiLog.SetLogger("file", `{"filename":"log/api.log"}`)
	//初始化邮箱
	mail = utils.NewEMail(`{"username":"night@isletsacg.com","password":"` + beego.AppConfig.String("EmailPassword") + `","host":"smtp.isletsacg.com","port":25}`)
}

/* 发送邮件 */
func SendEmail(address []string, subject string, html string) {
	mail.To = address
	mail.From = "night@isletsacg.com"
	mail.Subject = subject
	mail.Text = ""
	mail.HTML = `
			<div style="border-bottom:3px solid #d9d9d9; background:url(http://www.ascode.net/static/img/email_bg.gif) repeat-x 0 1px;">
				<div style="border:1px solid #c8cfda; padding:40px;">
					` + html + `
					<p>&nbsp;</p>
					<div>我酷游戏团队 祝您游戏愉快</div>
					<div>Powered by AsCode</div>
					<img src="http://www.ascode.net/static/img/logo.png">
					</div>
				</div>
			</div>
			`
	err := mail.Send()
	if err != nil { //邮件未发送成功，记录错误日志
		ApiLog.Error("邮件发送失败：", err)
	}
}
