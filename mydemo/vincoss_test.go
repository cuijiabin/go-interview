package mydemo

import "testing"

func Test_GuIp(t *testing.T) {
	m := GenIp("2502501135")
	for _, v := range m {
		println(v)
	}

}
