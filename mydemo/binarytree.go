package mydemo

/*
一些经典算法的实现 二叉树的遍历
*/
import (
	"fmt"
	"math/rand"
)

//结构体
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

//创建一颗树
func NewTree(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = Insert(t, (1+v)*k)
	}
	return t
}

func Insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = Insert(t.Left, v)
		return t
	}
	t.Right = Insert(t.Right, v)
	return t
}

//中序遍历
func Print(t *Tree) { //Recursive
	if t == nil {
		return
	}
	Print(t.Left)
	fmt.Println("node:", t.Value)
	Print(t.Right)
}

func Search(t *Tree, v int) bool {

	if t == nil {
		return false
	}
	switch {
	case v == t.Value:
		return true
	case v < t.Value:
		return Search(t.Left, v)
	case v > t.Value:
		return Search(t.Right, v)
	}
	return false
}

func GetMin(t *Tree) (int, bool) {
	if t == nil {
		return -1, false
	}

	for {
		if t.Left != nil {
			t = t.Left
		} else {
			return t.Value, true
		}
	}
}

func GetMax(t *Tree) (int, bool) {
	if t == nil {
		return -1, false
	}
	for {
		if t.Right != nil {
			t = t.Right
		} else {
			return t.Value, true
		}
	}
}

func Delete(t *Tree, v int) bool {
	if t == nil {
		return false
	}

	parent := t
	found := false
	for {
		if t == nil {
			break
		}
		if v == t.Value {
			found = true
			break
		}

		parent = t
		if v < t.Value { //left
			t = t.Left
		} else {
			t = t.Right
		}
	}

	if found == false {
		return false
	}
	return deleteNode(parent, t)
}

func deleteNode(parent, t *Tree) bool {
	if t.Left == nil && t.Right == nil {
		fmt.Println("delete() 左右树都为空 ")
		if parent.Left == t {
			parent.Left = nil
		} else if parent.Right == t {
			parent.Right = nil
		}
		t = nil
		return true
	}

	if t.Right == nil { //右树为空
		fmt.Println("delete() 右树为空 ")
		parent.Left = t.Left.Left
		parent.Value = t.Left.Value
		parent.Right = t.Left.Right
		t.Left = nil
		t = nil
		return true
	}

	if t.Left == nil { //左树为空
		fmt.Println("delete() 左树为空 ")
		parent.Left = t.Right.Left
		parent.Value = t.Right.Value
		parent.Right = t.Right.Right
		t.Right = nil
		t = nil
		return true
	}

	fmt.Println("delete() 左右树都不为空 ")
	previous := t
	//找到左子节点的最右叶节点，将其值替换至被删除节点
	//然后将这个最右叶节点清除，所以说，为了维持树，
	//这种情况下，这个最右叶节点才是真正被删除的节点
	next := t.Left
	for {
		if next.Right == nil {
			break
		}
		previous = next
		next = next.Right
	}

	t.Value = next.Value
	if previous.Left == next {
		previous.Left = next.Left
	} else {
		previous.Right = next.Right
	}
	next.Left = nil
	next.Right = nil
	next = nil
	return true
}

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *Tree) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

// Compare reads values from two Walkers
// that run simultaneously, and returns true
// if t1 and t2 have the same contents.
func Compare(t1, t2 *Tree) bool {
	c1, c2 := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}
