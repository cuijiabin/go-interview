package mydemo

import "testing"

//测试声明数组
func Test_StateArray(t *testing.T) {
	StateArray()
}

func Test_AssignArray(t *testing.T) {
	AssignArray()
}

func Test_InferArray(t *testing.T) {
	InferArray()
}

func Test_PosiArray(t *testing.T) {
	PosiArray()
}

func Test_StateSlice(t *testing.T) {
	StateSlice()
}

func Test_CapSlice(t *testing.T) {
	CapSlice()
}

func Test_AssignSlice(t *testing.T) {
	AssignSlice()
}

func Test_AppendSlice(t *testing.T) {
	AppendSlice()
}

func Test_RemoveSlice(t *testing.T) {
	RemoveSlice()
}

func Test_StateMap(t *testing.T) {
	StateMap()
}
