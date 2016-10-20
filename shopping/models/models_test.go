package models

import "testing"
import (
	"fmt"
	"time"
)

//测试声明数组
func Test_TableName(t *testing.T) {
	u := User{Name: "Tom", Email: "cuijia_bin@126.com"}
	t_name := u.TableName()

	fmt.Println(t_name)

}

func Test_Insert(t *testing.T) {
	u := &User{Name: "cuijiabin3", Email: "cuijia_bin@126.com", Password: Md5("123456"), Mobile: "18910358924", CreateAt: time.Now().Unix(), UpdateAt: time.Now().Unix()}
	fmt.Println(u.Insert())
}

func Test_Update(t *testing.T) {
	u := &User{Id: 2, Name: "修改后", Email: "cuijia_bin@164.com"}
	fmt.Println(u.Update("Email"))
}

func Test_DeleteById(t *testing.T) {
	u := &User{Id: 2}
	fmt.Println(u.DeleteById())
}

func Test_Delete(t *testing.T) {
	u := &User{Id: 1}
	fmt.Println(u.Delete())
}

func Test_Md5(t *testing.T) {
	fmt.Println(Md5("cuijiabin"))
}

func TestUser_Read(t *testing.T) {
	u := &User{Id: 2}
	u.Read("Password")
	fmt.Println(u.Password)
}

func TestUser_InsertMulti(t *testing.T) {
	u1 := User{Name: "cuijiabin4", Email: "cuijia_bin@126.com", Password: Md5("123456"), Mobile: "18910358924", CreateAt: time.Now().Unix(), UpdateAt: time.Now().Unix()}
	u2 := User{Name: "cuijiabin5", Email: "cuijia_bin@126.com", Password: Md5("123456"), Mobile: "18910358924", CreateAt: time.Now().Unix(), UpdateAt: time.Now().Unix()}
	var users [2]User
	users[0] = u1
	users[1] = u2
}

func TestUser_QueryUser(t *testing.T) {
	//u := &User{Email: "cuijia_bin@126.com"}
	u := &User{Id: 1} // 不管用
	l := u.QueryUser()
	for _, v := range l {
		fmt.Println(*v)
	}
}

func TestUser_SqlQuery(t *testing.T) {
	u := &User{Id: 1}
	u.SqlQuery()
}
