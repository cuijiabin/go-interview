package rtda

import "jvm-book/ch09/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
