package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

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

type Label struct{
	Id        bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	CreateAt  time.Time        `bson:"create_at"`
}
