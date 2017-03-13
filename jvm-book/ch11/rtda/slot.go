package rtda

import "jvm-book/ch11/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
