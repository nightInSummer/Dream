package controllers

import (
	"Dream/models"
	"github.com/astaxie/beego"
	"math"
)

type TalkController struct {
	beego.Controller
}

func (this *TalkController) Topic() {
	//专题页面
	talk := models.Talk{}
	allPage := math.Ceil(float64(talk.Count()) / 15)
	this.Data["allpage"] = allPage
	this.TplNames = ""
}

func (this *TalkController) ShowTopic() {
	//专题列表分页操作
	page, _ := this.GetInt("page")
	//查询前page~page+15个我的文章
	talk := models.Talk{}
	//fmt.Println(allPage)
	articles := talk.List(uint32((page-1)*15), uint32(page*15))
	this.Data["json"] = articles
	this.ServeJson()
}

func (this *TalkController) Talk() {
	//专题下二级页面
	talk := models.Talk{}
	allPage := math.Ceil(float64(talk.Count()) / 15)
	this.Data["allpage"] = allPage
	this.TplNames = ""
}

func (this *TalkController) ShowTalk() {
	//专题列表分页操作
	page, _ := this.GetInt("page")
	//查询前page~page+15个我的文章
	talk := models.Talk{}
	//fmt.Println(allPage)
	id, _ := this.GetInt("id")
	articles := talk.Show(int(id), uint32((page-1)*15), uint32(page*15))
	this.Data["json"] = articles
	this.ServeJson()

}

func (this *TalkController) BuildTopic() {
	//创建专题
	talk := models.Talk{}
	talk.Nickname = "" //创建人昵称
	talk.Topic = ""    //用户创建专题
	talk.Short = ""    //话题简介
	talk.Pid = 0       //分类级别专题和单独的话题最高级别pid为0

}

func (this *TalkController) Publish() {
	talk := models.Talk{}
	//发布话题
	status, _ := this.GetInt("status")
	if status == 0 {
		//判断传过来的status，如果为0则是单独的话题
		talk.Nickname = "" //发布人昵称
		talk.Talk = ""     //用户创建专题
		talk.Pid = 0       //分类级别专题和单独的话题最高级别pid为0
	} else {
		//判断传过来的status，如果不为0则是专题下的话题
		id, _ := this.GetInt("id")
		talk.Nickname = "" //发布人昵称
		talk.Talk = ""     //用户创建专题
		talk.Pid = id      //分类级别专题和单独的话题最高级别pid为0
	}
}

func (this *TalkController) Discuss() {
	//对话题的评论操作
	reply := models.Reply{}
	//发布话题
	id, _ := this.GetInt("id")
	status, _ := this.GetInt("status")
	if status == 0 {
		//判断传过来的status，如果为0则是深度评论
		reply.Uid = id
		reply.Nickname = "" //评论人昵称
		reply.Contect = ""  //评论内容
		reply.Pid = 0       //pid为0为对话题的评论
		reply.Status = 0    //深度评论
	} else {
		//判断传过来的status，如果不为0则是水区评论
		reply.Uid = id
		reply.Nickname = "" //评论人昵称
		reply.Contect = ""  //评论内容
		reply.Pid = 0       //pid为0为对话题的评论
		reply.Status = 1    //深度评论
	}
}

func (this *TalkController) Speak() {
	//对评论的评论操作
	reply := models.Reply{}
	//发布话题
	id, _ := this.GetInt("id")
	status, _ := this.GetInt("status")
	if status == 0 {
		//判断传过来的status，如果为0则是深度评论
		reply.Uid = id
		reply.Nickname = "" //评论人昵称
		reply.Contect = ""  //评论内容
		reply.Pid = id      //pid不为0为对评论的评论，数字对应评论的id
		reply.Status = 0    //深度评论
	} else {
		//判断传过来的status，如果不为0则是水区评论
		reply.Uid = id
		reply.Nickname = "" //评论人昵称
		reply.Contect = ""  //评论内容
		reply.Pid = id      //pid不为0为对评论的评论，数字对应评论的id
		reply.Status = 1    //水区评论
	}
}
