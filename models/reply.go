package models

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Reply struct {
	Id       bson.ObjectId `bson:"_id"` //主键
	Uid      int64         `bson:"u"`   //评论对应的话题id
	Nickname string        `bson:"n"`   //专题发起人昵称
	Contect  string        `bson:"c"`   //评论内容
	Time     string        `bson:"t"`   //评论发出时间
	Pid      int64         `bson:"p"`   //评论级别为0 是对话题的评论，若不为0则为评论别人的评论，数字对应评论的id
	Status   int64         `bson:"st"`  //标记字段若为0则为深度评论，若为1则为水区评论
	Type     int64         `bson:"ty"`  //标记字段，1为讨论区评论，2为作品区评论
}

var (
	replyC *mgo.Collection //数据库连接
)

func init() {
	//获取数据库连接
	session, err := mgo.Dial(beego.AppConfig.String("MongoDb"))
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	replyC = session.DB("Dream").C("reply")
}

func (this *Reply) Insert() string {
	//数据表中插入数据
	this.Id = bson.NewObjectId()
	this.Time = time.Now().Format("2006-01-02 15:04:05")
	err := talkC.Insert(this)
	if err != nil {
		panic(err)
	}
	return bson.ObjectId.Hex(this.Id)
}
