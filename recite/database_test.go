package main

import (
	"fmt"
	"testing"
)

func TestAddMgoRepeat(t *testing.T) {
	re := Repeat{RpContent: "重复的进行复习3", IsCorrect: false}
	id := AddMgoRepeat(re, "58c424463ec97445f492888c")
	fmt.Println(id)
}

func TestGetRepeatById(t *testing.T) {
	repeat := GetRepeatById("58c5268e3ec97411c08e26a5")
	fmt.Println(repeat)
}
func TestUpdateRepeatById(t *testing.T) {
	result := EditMgoRepeat("58c5268e3ec97411c08e26a5", "添加备注呵呵")
	fmt.Println(result)
}

func TestDelRepeatById(t *testing.T) {
	fmt.Println(DelRepeatById("58c51b1f3ec97409407f957a"))
}

func TestListMgoRepeat(t *testing.T) {
	list := ListMgoRepeat("58c424463ec97445f492888c")
	fmt.Println(list)
}

func TestStatConditon(t *testing.T) {
	r := StatConditon("58c424463ec97445f492888c")
	fmt.Println(r.TCount, r.RCount, r.WCount)
}

func TestStatMgoRecite(t *testing.T) {
	r := StatMgoRecite("58c424463ec97445f492888c")
	fmt.Println(r)
}
