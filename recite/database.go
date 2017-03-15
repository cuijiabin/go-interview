package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const MGO_URL = "127.0.0.1:27017"

var (
	mgoSession *mgo.Session
	dataBase   = "myrecite"
	tblRecite  = "recite"
	tblRepeat  = "repeat"
	tblLabel   = "label"
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
	err := witchCollection(tblRecite, query)
	if err != nil {
		return "false"
	}
	return r.Id.Hex()
}

func GetReciteById(id string) *Recite {
	objId := bson.ObjectIdHex(id)
	recite := new(Recite)
	query := func(c *mgo.Collection) error {
		return c.FindId(objId).One(&recite)
	}
	witchCollection(tblRecite, query)
	return recite
}

func EditMgoRecite(id string, r Recite) bool {
	data := bson.M{"$set": bson.M{"title": r.Title, "content": r.Content, "tip": r.Tip}}
	query := func(c *mgo.Collection) error {
		return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, data)
	}
	err := witchCollection(tblRecite, query)
	if err != nil {
		return false
	}
	return true
}
func DelMgoRecite(id string) bool {
	query := func(c *mgo.Collection) error {
		return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	}
	err := witchCollection(tblRecite, query)
	if err != nil {
		fmt.Println("删除失败" + err.Error())
		return false
	}
	return true
}

//更新回答情况
func StatMgoRecite(id string) bool {
	rc := StatConditon(id)
	data := bson.M{"$set": bson.M{"r_condition": &rc}}
	query := func(c *mgo.Collection) error {
		return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, data)
	}
	err := witchCollection(tblRecite, query)
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
	c := session.DB(dataBase).C(tblRepeat)
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
	c := session.DB(dataBase).C(tblRecite)
	page.tC, err = c.Find(nil).Count()
	if err != nil {
		return page
	}

	query := func(c *mgo.Collection) error {
		return c.Find(nil).Sort("-create_at").Skip((page.Cp - 1) * page.Ps).Limit(page.Ps).All(&recites)
	}
	err = witchCollection(tblRecite, query)
	if err != nil {
		return page
	}
	page.list = &recites
	return page
}

//查看背诵列表
func ListMgoRepeat(rId string) []Repeat {
	var repeats []Repeat
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"r_id": bson.ObjectIdHex(rId)}).All(&repeats)
	}
	err := witchCollection(tblRepeat, query)
	if err != nil {
		return repeats
	}
	return repeats
}

//添加背诵选项
func AddMgoRepeat(r Repeat, rId string) string {
	r.Id = bson.NewObjectId()
	r.RId = bson.ObjectIdHex(rId)
	r.CreateAt = time.Now()
	query := func(c *mgo.Collection) error {
		return c.Insert(r)
	}
	err := witchCollection(tblRepeat, query)
	if err != nil {
		return "false"
	}
	return r.Id.Hex()
}

//查看背诵详情
func GetRepeatById(id string) *Repeat {
	objId := bson.ObjectIdHex(id)
	repeat := new(Repeat)
	query := func(c *mgo.Collection) error {
		return c.FindId(objId).One(&repeat)
	}
	witchCollection(tblRepeat, query)
	return repeat
}

//修改背诵备注
func EditMgoRepeat(id string, remark string) bool {
	//data := bson.M{"remark": remark}
	data := bson.M{"$set": bson.M{"remark": remark}}
	query := func(c *mgo.Collection) error {
		return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, data)
	}
	err := witchCollection(tblRepeat, query)
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
	err := witchCollection(tblRepeat, query)
	if err != nil {
		fmt.Println("删除失败" + err.Error())
		return false
	}
	return true
}

func AddMgoLabel(r Label) string {
	r.Id = bson.NewObjectId()
	r.CreateAt = time.Now()
	query := func(c *mgo.Collection) error {
		return c.Insert(r)
	}
	err := witchCollection(tblLabel, query)
	if err != nil {
		return "false"
	}
	return r.Id.Hex()
}

func DetailMgoLabel(id string) *Label {
	objId := bson.ObjectIdHex(id)
	label := new(Label)
	query := func(c *mgo.Collection) error {
		return c.FindId(objId).One(&label)
	}
	witchCollection(tblLabel, query)
	return label
}

func EditMgoLabel(id string, name string) bool {
	data := bson.M{"$set": bson.M{"name": name}}
	query := func(c *mgo.Collection) error {
		return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, data)
	}
	err := witchCollection(tblLabel, query)
	if err != nil {
		return false
	}
	return true
}
func DelMgoLabel(id string) bool {
	query := func(c *mgo.Collection) error {
		return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	}
	err := witchCollection(tblLabel, query)
	if err != nil {
		fmt.Println("删除失败" + err.Error())
		return false
	}
	return true
}
