package rtda

import "jvm-book/ch10/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
