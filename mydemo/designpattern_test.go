package mydemo

import (
	"testing"
	"fmt"
	"bytes"
	//"strings"
)

func TestPxgen(t *testing.T) {
	Pxgen()
}

func TestGetNotice(t *testing.T) {
	jack := GetNotice("jack") //  获取jack的消息
	joe := GetNotice("joe") // 获取joe的消息

	// 获取消息的返回
	fmt.Println(<-jack)
	fmt.Println(<-joe)
}

func TestMergeGo(t *testing.T) {
	MergeGo()
}

func TestSelectDemo(t *testing.T) {
	SelectDemo()
}

func TestPrimeNumber(t *testing.T) {
	PrimeNumber()
}

func TestGenRand(t *testing.T) {
	GenRand()
}

func TestConvert(t *testing.T)  {
	var m int = 1
	var k int32 = 95
	var n int64 = 89
	fmt.Println(int64(m)*n)
	fmt.Println(m*int(n))
	fmt.Println(int32(m)*k)

	//高效字符串拼接
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.WriteString("崔佳彬")
	buffer.WriteString("你好")

	fmt.Println("拼接后的结果为-->", buffer.String())

}
