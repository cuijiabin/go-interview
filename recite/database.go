package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const MGO_URL = "127.0.0.1:27017"

type Recite struct {
	Id         bson.ObjectId `bson:"_id"`
	Title      string        `bson:"title"`
	Content    string        `bson:"content"`
	Tip        string        `bson:"tip"`
	RCondition RepeatCondition `bson:"r_condition"`
	CreateAt   time.Time        `bson:"create_at"`
}

type RepeatCondition struct {
	TCount int `bson:"t_count"`
	RCount int `bson:"r_count"`
	WCount int `bson:"w_count"`
}

type Repeat struct {
	Id        bson.ObjectId `bson:"_id"`
	RId       bson.ObjectId `bson:"r_id"`
	RpContent string        `bson:"rp_content"`
	Remark    string        `bson:"remark"`
	IsCorrect bool        `bson:"is_correct"`
	CreateAt  time.Time        `bson:"create_at"`
}

type PageRecite struct {
	Cp   int `json:"Cp"` //当前页
	Ps   int `json:"Ps"` //每页大小
	tP   int `json:"tP"` //总页数
	tC   int `json:"tC"` //总条数
	list *[]Recite `json:"list"`
}

var (
	mgoSession *mgo.Session
	dataBase   = "myrecite"
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(MGO_URL)
		if err != nil {
			panic(err) //直接终止程序运行
		}
	}
	//最大连接池默认为4096
	return mgoSession.Clone()
}

func witchCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return s(c)
}

//添加背诵选项
func AddMgoRecite(r Recite) string {
	r.Id = bson.NewObjectId()
	r.CreateAt = time.Now()
	query := func(c *mgo.Collection) error {
		return c.Insert(r)
	}
	err := witchCollection("recite", query)
	if err != nil {
		return "false"
	}
	return r.Id.Hex()
}

func EditMgoRecite(id string, r Recite) bool {
	data := bson.M{"$set": bson.M{"title": r.Title, "content": r.Content, "tip": r.Tip}}
	query := func(c *mgo.Collection) error {
		return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, data)
	}
	err := witchCollection("recite", query)
	if err != nil {
		return false
	}
	return true
}

//TODO 更新
func StatMgoRecite(id string) bool {
	rc := StatConditon(id)
	data := bson.M{"$set": bson.M{"r_condition": &rc}}
	query := func(c *mgo.Collection) error {
		return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, data)
	}
	err := witchCollection("recite", query)
	if err != nil {
		fmt.Println("更新统计失败" + err.Error())
		return false
	}
	return true
}

//统计回答数量
func StatConditon(id string) *RepeatCondition {
	rc := &RepeatCondition{0, 0, 0}
	session := getSession()
	defer session.Close()
	c := session.DB(dataBase).C("repeat")
	rc.TCount, _ = c.Find(bson.M{"r_id": bson.ObjectIdHex(id)}).Count()
	rc.RCount, _ = c.Find(bson.M{"r_id": bson.ObjectIdHex(id), "is_correct": true}).Count()
	rc.WCount, _ = c.Find(bson.M{"r_id": bson.ObjectIdHex(id), "is_correct": false}).Count()
	return rc
}

//查询背诵选项列表
func PageMgoRecite(page PageRecite) PageRecite {
	var recites []Recite
	var err error
	//TODO 可以优化处理
	session := getSession()
	defer session.Close()
	c := session.DB(dataBase).C("recite")
	page.tC, err = c.Find(nil).Count()
	if err != nil {
		return page
	}

	query := func(c *mgo.Collection) error {
		return c.Find(nil).Sort("-create_at").Skip((page.Cp - 1) * page.Ps).Limit(page.Ps).All(&recites)
	}
	err = witchCollection("recite", query)
	if err != nil {
		return page
	}
	page.list = &recites
	return page
}

func GetReciteById(id string) *Recite {
	objId := bson.ObjectIdHex(id)
	recite := new(Recite)
	query := func(c *mgo.Collection) error {
		return c.FindId(objId).One(&recite)
	}
	witchCollection("recite", query)
	return recite
}

//添加背诵选项
func AddMgoRepeat(r Repeat, rId string) string {
	r.Id = bson.NewObjectId()
	r.RId = bson.ObjectIdHex(rId)
	r.CreateAt = time.Now()
	query := func(c *mgo.Collection) error {
		return c.Insert(r)
	}
	err := witchCollection("repeat", query)
	if err != nil {
		return "false"
	}
	return r.Id.Hex()
}

//查看背诵列表
func ListMgoRepeat(rId string) []Repeat {
	var repeats []Repeat
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"r_id": bson.ObjectIdHex(rId)}).All(&repeats)
	}
	err := witchCollection("repeat", query)
	if err != nil {
		return repeats
	}
	return repeats
}

//查看背诵详情
func GetRepeatById(id string) *Repeat {
	objId := bson.ObjectIdHex(id)
	repeat := new(Repeat)
	query := func(c *mgo.Collection) error {
		return c.FindId(objId).One(&repeat)
	}
	witchCollection("repeat", query)
	return repeat
}

//修改背诵备注
func UpdateRepeatById(id string, remark string) bool {
	//data := bson.M{"remark": remark}
	data := bson.M{"$set": bson.M{"remark": remark}}
	query := func(c *mgo.Collection) error {
		return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, data)
	}
	err := witchCollection("repeat", query)
	if err != nil {
		fmt.Println("删除失败" + err.Error())
		return false
	}
	return true
}

//删除背诵
func DelRepeatById(id string) bool {
	query := func(c *mgo.Collection) error {
		return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	}
	err := witchCollection("repeat", query)
	if err != nil {
		fmt.Println("删除失败" + err.Error())
		return false
	}
	return true
}
