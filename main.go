package main

import (
	"Dream/models"
	_ "Dream/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"regexp"
	"time"
)

func main() {
	req := httplib.Get("http://acg.178.com/")
	//匹配操作
	reg := regexp.MustCompile(`<div class="news_box">([\w\W]+?)<div class="between_line">`) //匹配div里class为news_box的元素
	rep := regexp.MustCompile(`href="(.*?)" title="(.*?)"`)                                 //匹配a标签里的内容
	rec := regexp.MustCompile(`src="(.*?)"`)                                                //匹配图片url
	red := regexp.MustCompile(`<div class="newstext text_i">([\w\W]+?)<a`)                  //匹配内容简介
	ref := regexp.MustCompile(`<span class="(.*?)">([\w\W]+?)</span>`)                      //匹配文章类型
	rei := regexp.MustCompile(`/&nbsp;</span>([\w\W]+?)&`)                                  //匹配文章发布时间
	//处理操作
	rel := regexp.MustCompile(`title="(.*?)"`)                 //去除title属性
	reh := regexp.MustCompile(`href="(.*?)"`)                  //去除href属性
	re := regexp.MustCompile(`href=`)                          //去除href=
	re1 := regexp.MustCompile(`title=`)                        //去除title=
	re2 := regexp.MustCompile(`src=`)                          //去掉src=
	re3 := regexp.MustCompile(`<div class="newstext text_i">`) //去掉div及里面的属性
	re4 := regexp.MustCompile(`<span class="(.*?)">`)          //去掉span及里面的属性
	re5 := regexp.MustCompile(`/&nbsp;</span>`)                //去掉开头无用元素
	ret := regexp.MustCompile(`"`)                             //去除""
	rew := regexp.MustCompile(` `)                             //去除空格
	rez := regexp.MustCompile(`<a`)                            //去掉<a
	rey := regexp.MustCompile(`</span>`)                       //去掉</span>
	rex := regexp.MustCompile(`&`)                             //去掉&
	str, _ := req.String()
	con := reg.FindAllString(str, -1)
	news := &models.News{}
	for k, v := range con {
		fmt.Println(k)
		//抓取方向
		sub := rep.FindAllString(v, -1) //抓取a标签里的内容
		suc := rec.FindAllString(v, -1) //抓取img的url
		sud := red.FindAllString(v, -1) //抓取内容简介
		sue := ref.FindAllString(v, -1) //抓取文章类型
		suf := rei.FindAllString(v, -1) //抓取文章时间
		//预处理
		sr := rel.ReplaceAllString(sub[0], "") //去掉不需要的title属性
		se := reh.ReplaceAllString(sub[0], "") //去掉不需要的href属性
		sr1 := re.ReplaceAllString(sr, "")     //去掉不需要的href=
		sr2 := re1.ReplaceAllString(se, "")    //去掉不需要的title=
		if suc == nil {
			continue
		}
		sr3 := re2.ReplaceAllString(suc[0], "") //去掉不需要的src=
		sr4 := re3.ReplaceAllString(sud[0], "") //去掉不需要的div及里面的属性
		sr5 := re4.ReplaceAllString(sue[0], "") //去掉不需要的span及里面的属性
		sr6 := re5.ReplaceAllString(suf[0], "") //去掉不需要的</span>
		//获取需要的内容
		urt := ret.ReplaceAllString(sr1, "")   //取到url
		url := rew.ReplaceAllString(urt, "")   //去掉url最后的空格
		title := ret.ReplaceAllString(sr2, "") //取到title
		img := ret.ReplaceAllString(sr3, "")   //取到img的url
		short := rez.ReplaceAllString(sr4, "") //取到contect简介
		style := rey.ReplaceAllString(sr5, "") //取到文章type
		time1 := rex.ReplaceAllString(sr6, "") //取到文章time
		s := url
		res := regexp.MustCompile(`http`)
		status := res.MatchString(s)
		var url1 string
		if status {
			url1 = url
		} else {
			url1 = "http://acg.178.com" + url
		}
		t := news.CheckUrl(url1)
		if t {
			//如果抓取到的url已经在数据空中存在，跳出本次循环
			continue
		} else {
			news.Url = url1
			news.Title = title
			news.Img = img
			news.Short = short
			news.Type = style
			news.Time = time.Now().Format("2006-01-02 15:04:05")
			news.InitialTime = time1
			news.Insert()
		}
	}
	beego.Run()
}
