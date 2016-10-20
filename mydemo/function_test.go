package mydemo

import "testing"
import "fmt"

func Test_Abs(t *testing.T) {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}
