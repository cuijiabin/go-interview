package rtda

import "jvm-book/ch07/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
