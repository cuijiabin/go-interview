package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64 `orm:"pk;auto"` //主键，自动增长
	Name     string
	Email    string
	Password string
	Mobile   string
	CreateAt int64
	UpdateAt int64
}

func init() {
	fmt.Print("初始化数据库的配置")
	// 需要在init中注册定义的model
	mysqluser := "root"
	mysqlpass := "123456"
	mysqlurls := "127.0.0.1:3306"
	mysqldb := "beego"
	orm.RegisterModel(new(User))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+")/"+mysqldb+"?charset=utf8&loc=Asia%2FShanghai")
	orm.RunSyncdb("default", false, true)
}

func (m *User) TableName() string {
	return "user"
}

//新增用户信息
func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

//批量插入 TODO 貌似不对呀
//func (ua []*User) InsertMul() error {
//	if _, err := orm.NewOrm().InsertMulti(100, ua); err != nil {
//		return err
//	}
//	return nil
//}

//删除用户
func (u *User) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

//删除用户
func (u *User) DeleteById() error {
	if _, err := orm.NewOrm().Delete(&User{Id: u.Id}); err != nil {
		return err
	}
	return nil
}

//修改用户信息
func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

//TODO Read 有什么用呢？
func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func Md5(str string) string {

	m := md5.New()
	m.Write([]byte(str))
	cipherText := m.Sum(nil)
	hexText := make([]byte, 32)
	hex.Encode(hexText, cipherText)

	return string(hexText)
}

//高级查询了
func (u *User) QueryUser() []*User {
	list := make([]*User, 0)
	qs := orm.NewOrm().QueryTable(u)
	qs.All(&list, "Id", "Name", "Email", "Password", "Mobile", "CreateAt", "UpdateAt")
	return list
}

//原生sql查询
func (u *User) SqlQuery() {
	var users []User
	r := orm.NewOrm().Raw("SELECT id, name FROM user WHERE email = ?")
	num, err := r.SetArgs("cuijia_bin@126.com").QueryRows(&users)
	if err == nil {
		fmt.Println("user nums: ", num)
		fmt.Println("users: ", users)
	}
}
