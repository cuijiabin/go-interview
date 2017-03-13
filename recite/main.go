package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	wd := "E:/another/gospace/src/recite/public"
	fmt.Println(wd)
	http.Handle("/", http.FileServer(http.Dir(wd)))

	//列表
	http.HandleFunc("/list", ReciteList)
	//详情
	http.HandleFunc("/detailRecite", DetailRecite)
	// 添加
	http.HandleFunc("/addRecite", AddRecite)
	http.HandleFunc("/editRecite", EditRecite)

	//添加背诵记录
	http.HandleFunc("/addRepeat", AddRepeat)
	http.HandleFunc("/rpList", RepeatList)

	//TODO 修改
	//TODO 删除
	//TODO 重复

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
