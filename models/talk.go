package models

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Talk struct {
	Id       bson.ObjectId `bson:"_id"` //主键
	Nickname string        `bson:"n"`   //专题发起人昵称
	Topic    string        `bson:"t"`   //专题
	Talk     string        `bson:"ta"`  //话题
	Time     string        `bson:"ti"`  //发布时间
	Short    string        `bson:"s"`   //专题简介
	Pid      int64         `bson:"pi"`  //父类id
}

var (
	talkC *mgo.Collection //数据库连接
)

func init() {
	//获取数据库连接
	session, err := mgo.Dial(beego.AppConfig.String("MongoDb"))
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	talkC = session.DB("Dream").C("talk")
}

func (this *Talk) Insert() string {
	//数据表中插入数据
	this.Id = bson.NewObjectId()
	this.Time = time.Now().Format("2006-01-02 15:04:05")
	err := talkC.Insert(this)
	if err != nil {
		panic(err)
	}
	return bson.ObjectId.Hex(this.Id)
}

func (this *Talk) Count() int {
	count, err := talkC.Find(nil).Count()
	if err != nil {
		//处理错误
	}
	return count
}

func (this *Talk) List(from uint32, to uint32) []*Talk {
	//查询讨论列表
	//使用无限分类，读出列表的是时候查询pid为0的数据，读出所有专题和单独的话题
	var result []*Talk
	err := talkC.Find(bson.M{"pi": 0}).Skip(int(from)).Limit(int(to - from)).All(&result)
	if err != nil {
		//处理错误
	}
	return result
}

func (this *Talk) Show(id int, from uint32, to uint32) []*Talk {
	//查询专题下话题列表
	var result []*Talk
	err := talkC.Find(bson.M{"pi": id}).Skip(int(from)).Limit(int(to - from)).All(&result)
	if err != nil {
		//处理错误
	}
	return result
}
