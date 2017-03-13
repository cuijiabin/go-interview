package rtda

import "jvm-book/ch08/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
