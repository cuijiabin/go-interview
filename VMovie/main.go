package main

import (
	"VMovie/models"
	_ "VMovie/routers"

	"github.com/astaxie/beego"
)

func main() {
	//模板函数
	beego.AddFuncMap("getclassname", models.GetMovieClassNameByCid)
	beego.AddFuncMap("GetMovieUpdateEP", models.GetMovieUpdateEP)
	beego.AddFuncMap("GetIPhoto", models.GetIPhoto)
	beego.AddFuncMap("GetMovieUpdateEPString", models.GetMovieUpdateEPString)
	beego.Run()
}
