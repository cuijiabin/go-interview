package mydemo

import (
	"testing"
)

var path = "E:/File/图书/golang/go语言面试题.txt"

func TestOpenFile(t *testing.T) {
	OpenFile(path)
}
func TestReadAll(t *testing.T) {
	ReadAll(path)
}
func TestBufRead(t *testing.T) {
	BufRead(path)
}
func TestBFIOReader(t *testing.T) {
	BFIOReader(path)
}
func TestCreateFile(t *testing.T) {
	CreateFile("E:/File/图书/golang/go语言面试题2.txt")
}
func TestCheckExist(t *testing.T) {
	CheckExist(path)
}

func TestIoWrite(t *testing.T) {
	IoWrite("E:/File/图书/golang/go语言面试题2.txt","崔佳彬",4)
}
