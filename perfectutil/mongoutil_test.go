package main

import (
	"fmt"
	"testing"
)

//测试添加用户
func TestAddPerson(t *testing.T) {
	p := Person{Name:"cuijiabin",Phone:"18910358924"}
	id := AddPerson(p)
	fmt.Println(id)
}

func TestGetPersonById(t *testing.T) {
	p := GetPersonById("58c397b63ec9741bbcb20576")
	fmt.Println(p)
}
