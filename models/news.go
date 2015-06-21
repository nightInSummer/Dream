package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type News struct {
	Id          bson.ObjectId `bson:"_id"` //主键
	Title       string        `bson:"ti"`  //标题
	Url         string        `bson:"u"`   //标题url
	Time        string        `bson:"t"`   //写入时间
	InitialTime string        `bson:"in"`  //抓取文章发布时间
	Type        string        `bson:"ty"`  //内容类型
	Img         string        `bson:"i"`   //图片
	Short       string        `bson:"s"`   //内容简介
}

var (
	newsC *mgo.Collection //数据库连接
)

func init() {
	//获取数据库连接
	session, err := mgo.Dial(beego.AppConfig.String("MongoDb"))
	if err != nil {
		panic(err)
	}
	newsC = session.DB("Dream").C("news")
}

func (this *News) Insert() string {
	this.Id = bson.NewObjectId()
	err := newsC.Insert(this)
	if err != nil {
		panic(err)
	}
	return bson.ObjectId.Hex(this.Id)
}

func (this *News) CheckUrl(url string) bool {
	//检查数据空中是否有这个url
	var result []*News
	err := newsC.Find(bson.M{"u": url}).All(&result)
	if err != nil {
		//处理错误
	}
	fmt.Println(err)
	if result == nil {
		return false
	} else {
		return true
	}
}

func (this *News) Show(from uint32, to uint32) []*News {
	var result []*News
	err := newsC.Find(nil).Sort("-_id").Skip(int(from)).Limit(int(to - from)).All(&result)
	if err != nil {
		//处理错误
	}
	return result
}

func (this *News) Count() int {
	count, err := newsC.Find(nil).Count()
	if err != nil {
		//处理错误
	}
	return count
}
