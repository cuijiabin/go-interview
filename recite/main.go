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
	//编辑
	http.HandleFunc("/editRecite", EditRecite)
	//删除
	http.HandleFunc("/delRecite", DelRecite)

	//记录列表
	http.HandleFunc("/rpList", RepeatList)
	//添加记录
	http.HandleFunc("/addRepeat", AddRepeat)
	http.HandleFunc("/detailRepeat", DetailRepeat)
	http.HandleFunc("/editRepeat", EditRepeat)
	http.HandleFunc("/delRepeat", DelRepeat)

	//TODO 标签列表 暂时先不分页来显示
	http.HandleFunc("/lbList", LabelList)
	http.HandleFunc("/addLabel", AddLabel)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
