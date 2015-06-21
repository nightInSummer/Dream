package models

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Member struct {
	Id          bson.ObjectId `bson:"_id"` //主键
	Email       string        `bson:"e"`   //邮箱
	Phone       string        `bson:"p"`   //手机号
	Password    string        `bson:"ps"`  //密码
	Nickname    string        `bson:"n"`   //昵称
	Image       string        `bson:"i"`   //头像
	MessageN    string        `bson:"mn"`  //用户消息数量
	Message     string        `bson:"m"`   //用户消息
	Relation    string        `bson:"r"`   //用户好友关系
	Error       uint8         `bson:"er"`  //用户剩余机会
	StopTime    time.Time     `bson:"st"`  //账号冻结截止
	LockVersion uint64        `bson:"lv"`  //乐观锁
	FirstTime   time.Time     `bson:"f"`   //登录时刻
	LastTime    time.Time     `bson:"la"`  //最好操作时间
	Theme       string        `bson:"t"`   //用户关注的主题
	Type        string        `bson:"ty"`  //账号类型
	Title       string        `bson:"ti"`  //用户头衔(包括所属社团)
	Status      string        `bson:"s"`   //邮箱激活标记
}

var (
	memberC *mgo.Collection //数据库连接
)

func init() {
	//获取数据库连接
	session, err := mgo.Dial(beego.AppConfig.String("MongoDb"))
	if err != nil {
		panic(err)
	}
	memberC = session.DB("Dream").C("member")
}

func (this *Member) Update(change bson.M) bool {
	//加入乐观锁机制，仅当change更新字段有更新乐观锁时才会启用
	//若change中不含有乐观锁字段+1更新，每次查询必定有结果
	//如果某个更新可能产生并发问题，一定要更新乐观锁，这样下次更新
	//更新可能没有结果，返回err，重新更新
	colQuerier := bson.M{"_id": this.Id, "lv": this.LockVersion}
	err := memberC.Update(colQuerier, change)
	if err != nil { //更新出错
		return false
	} else {
		return true
	}
}

func (this *Member) Insert() string {
	//插入数据
	this.Id = bson.NewObjectId()
	this.FirstTime = bson.Now()
	this.LastTime = bson.Now()
	err := memberC.Insert(this)
	if err != nil {
		panic(err)
	}
	return bson.ObjectId.Hex(this.Id)
}

func (this *Member) Find(email string) uint8 {
	//查找用户账号剩余机会
	err := memberC.Find(bson.M{"e": email}).One(&this)
	if err != nil {
		//处理错误
	}
	return this.Error
}

func (this *Member) CheckPass(email string, password string) int {
	//验证登录信息
	err := memberC.Find(bson.M{"e": email}).One(&this)
	if err != nil {
		//处理错误
	}
	if bson.Now().Before(this.StopTime) { //锁定时间还没过
		return 0 //处于锁定状态
	}
	if this.Password == password {
		this.Error = 6
		this.Update(bson.M{"$set": bson.M{"er": this.Error}})
		return 1 //通过验证
	} else {
		if this.Error <= 1 {
			this.Error = 6
			minute := time.Duration(10) * time.Minute
			this.StopTime = bson.Now().Add(minute)
			this.Update(bson.M{"$set": bson.M{"er": this.Error, "st": this.StopTime}})
			return 2 //用户名或密码不正确
		} else {
			this.Error--
			this.Update(bson.M{"$set": bson.M{"er": this.Error}})
			return 2 //用户名或密码不正确
		}
	}
}
