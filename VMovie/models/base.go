package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqluser := "root"           //beego.AppConfig.String("mysqluser")
	mysqlpass := "123456"         //beego.AppConfig.String("mysqlpass")
	mysqlurls := "127.0.0.1:3306" //beego.AppConfig.String("mysqlurls")
	mysqldb := "beego"            //beego.AppConfig.String("mysqldb")
	orm.RegisterModel(new(RelationInfo), new(PageInfo), new(MovieInfo), new(MovieClassInfo), new(TvStationInfo), new(DownAddrClassInfo), new(DownAddrInfo), new(TagInfo), new(UserInfo), new(RecommendInfo))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+")/"+mysqldb+"?charset=utf8&loc=Asia%2FShanghai")
	orm.RunSyncdb("default", false, true)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}
