package controllers

import (
	"Dream/models"
	"github.com/astaxie/beego"
	"math"
)

type InformationController struct {
	beego.Controller
}

func (this *InformationController) Enter() {
	news := models.News{}
	allPage := math.Ceil(float64(news.Count()) / 15)
	this.Data["allpage"] = allPage
	this.TplNames = "information/jinru.html"
}

func (this *InformationController) Contect() {
	page, _ := this.GetInt("page")
	//查询前page~page+15个我的文章
	news := models.News{}
	//fmt.Println(allPage)
	articles := news.Show(uint32((page-1)*15), uint32(page*15))
	this.Data["json"] = articles
	this.ServeJson()
}
