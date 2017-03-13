package mydemo

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	test := []int{49, 38, 65, 97, 76, 13, 27, 49}
	InsertionSort(InputArray(test), 0, len(test))
	fmt.Println(test)
}

func TestBubbleSort(t *testing.T) {
	test := []int{49, 38, 65, 97, 76, 13, 27, 49}
	BubbleSort(InputArray(test))
	fmt.Println(test)

}

func TestQuickSort(t *testing.T) {
	test := []int{49, 38, 65, 97, 76, 13, 27, 49}
	QuickSort(test, 0, len(test)-1)
	fmt.Println(test)
}

func TestHeapSort(t *testing.T) {
	test := []int{49, 38, 65, 97, 76, 13, 27, 49}
	HeapSort(InputArray(test), 0, len(test))
	fmt.Println(test)
}

func TestConvert2(t *testing.T) {
	/**
	  二进制转换成十进制
	  核心:
	  入栈、出栈、类型转换
	  //sum += int(v-48) * int(math.Pow(2, N))
	**/

	stack := list.New()

	var input string
	var sum int
	var stnum, conum float64 = 0, 2

	fmt.Printf("请输入一段二进制数字:")
	fmt.Scanf("%s", &input)
	for _, c := range input {
		// 入栈 type rune
		stack.PushBack(c)
	}

	length := stack.Len()
	fmt.Printf("栈的当前容量是 %d\n", length)

	// 出栈
	for e := stack.Back(); e != nil; e = e.Prev() {
		// rune是int32的别名
		v := e.Value.(int32)
		sum += int(v-48) * int(math.Pow(conum, stnum))
		stnum++
	}
	fmt.Printf("二进制转化为十进制结果是 %d\n", sum)

}
