package mydemo

import "testing"

func Test_Tree(t *testing.T) {
	nt := NewTree(10, 1)

	//	if Search(t, 6) {
	//		fmt.Println("Search(6) true")
	//	} else {
	//		fmt.Println("Search(6) false")
	//	}
	Print(nt)

	//	if Delete(t, 6) {
	//		fmt.Println("Delete(6) true")
	//	} else {
	//		fmt.Println("Delete(6) false")
	//	}
	//	Print(t)

	//	if Delete(t, 9) {
	//		fmt.Println("Delete(9) true")
	//	} else {
	//		fmt.Println("Delete(9) false")
	//	}
	//	Print(t)

	//	min, foundMin := GetMin(t)
	//	if foundMin {
	//		fmt.Println("GetMin() =", min)
	//	}

	//	max, foundMax := GetMax(t)
	//	if foundMax {
	//		fmt.Println("GetMax() =", max)
	//	}

	//	t2 := New(100, 1)
	//	fmt.Println(Compare(t2, New(100, 1)), " Compare() Same Contents")
	//	fmt.Println(Compare(t2, New(99, 1)), " Compare() Differing Sizes")
}
