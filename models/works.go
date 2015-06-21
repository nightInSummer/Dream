package models

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Works struct {
	Id   bson.ObjectId `bson:"_id"` //主键
	Uid  int64         `bson:"u"`   //作品发起人id
	Time string        `bson:"t"`   //作品发布时间
	Type int64         `bson:"ty"`  //作品类型
	Pid  int64         `bson:"p"`   //父类id

}

var (
	worksC *mgo.Collection //数据库连接
)

func init() {
	//获取数据库连接
	session, err := mgo.Dial(beego.AppConfig.String("MongoDb"))
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	worksC = session.DB("Dream").C("works")
}

func (this *Works) Insert() string {
	//数据表中插入数据
	this.Id = bson.NewObjectId()
	this.Time = time.Now().Format("2006-01-02 15:04:05")
	err := talkC.Insert(this)
	if err != nil {
		panic(err)
	}
	return bson.ObjectId.Hex(this.Id)
}
