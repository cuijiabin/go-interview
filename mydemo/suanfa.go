package mydemo

type SortAble interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type InputArray []int

func (a InputArray) Len() int      { return len(a) }
func (a InputArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a InputArray) Less(i, j int) bool {
	return a[i] < a[j]
}

// 建立树函数
// 父节点root的左孩子2*root + 1
func siftDown(data SortAble, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi { // child 超出了数组长度，也就是该结点无孩子结点，返回
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) { // 右孩子结点存在
			child++
		}
		// 以上是为了在孩子结点当中找到较大的结点，用child表示
		if !data.Less(first+root, first+child) {
			return
		}
		// 如果根结点小于较大的孩子结点，则进行交换；该孩子结点的堆结构可能会被影响，继续去处理孩子结点
		data.Swap(first+root, first+child)
		root = child
	}
}

//插入排序
func InsertionSort(data SortAble, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

//TODO 还是有一些问题的
func BubbleSort(data SortAble) {
	flag := true
	for i := 1; i < data.Len(); i++ {
		flag = true
		for j := 0; j < data.Len()-i && data.Less(j+1, j); j++ {
			data.Swap(j+1, j)
			flag = false
		}
		if flag == true {
			break
		}
	}
}

func QuickSort(values []int, left, right int) {

	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}

	}

	values[p] = temp

	if p-left > 1 {
		QuickSort(values, left, p-1)
	}
	if right-p > 1 {
		QuickSort(values, p+1, right)
	}

}

//堆排序
func HeapSort(data SortAble, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	// 二叉树结构当中最后一个有子结点的结点
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}

//TODO 归并排序

//TODO 二分查找算法
