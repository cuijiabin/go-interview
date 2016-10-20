package mydemo

import "fmt"

//1.1声明数组
func StateArray() {
	//1.声明数组
	var a [2]int
	a[0] = 1
	a[1] = 2

	for m, v := range a {
		fmt.Println(m, v)
	}

	fmt.Println("数组长度：", len(a))

}

//func CopyArray(){
//	var a [89] int
//	a[1] = 1
//	for m,k := range a{
//		fmt.Print(m,k)
//	}
//}

//1.2数组直接赋值
func AssignArray() {
	a := [3]int{1, 2, 3} //声明并且赋值
	for m, v := range a {
		fmt.Println(m, v)
	}
}

//1.3推断数组长度
func InferArray() {
	a := [...]int{1, 2, 3} //声明并且赋值
	for m, v := range a {
		fmt.Printf("%p \n", &v)
		fmt.Println(m, v)
	}

	fmt.Println("数组长度：", len(a))
}

//1.4指定数组位置
func PosiArray() {
	a := [8]int{2: 6, 3: 4} //声明并且赋值且指定位置
	for m, v := range a {
		fmt.Println(m, v)
	}
	fmt.Println("数组长度：", len(a))
	fmt.Println("数组容量：", cap(a))
}

//注意对大内存对象时 考虑使用指针

//2.1切片操作 如果只指定了长度，那么容量默认等于长度
func StateSlice() {
	var ss []string
	fmt.Printf("length:%v \taddr:%p \tisnil:%v", len(ss), ss, ss == nil)

	slice := make([]string, 5)
	fmt.Println(slice)

	fmt.Println("切片长度", len(slice))
	fmt.Println("切片容量：", cap(slice))
	fmt.Println("容量是否等于长度：", len(slice) == cap(slice))
}

//2.2声明切片的长度还有容量
func CapSlice() {
	slice := make([]string, 3, 5)
	fmt.Println(slice)

	fmt.Println("切片长度", len(slice))
	fmt.Println("切片容量：", cap(slice))
	fmt.Println("容量是否等于长度：", len(slice) == cap(slice))
}

//2.3 直接切片赋值
func AssignSlice() {
	slice := []string{"cui", "jia", "bin", "qie"}

	fmt.Println("切片长度", len(slice))
	fmt.Println("切片容量：", cap(slice))

	newSlice := slice[1:3] //如果有第三个参数 则表明需要限定容量
	newSlice[0] = "新值"

	for idx, v := range newSlice {

		fmt.Println(idx, v)
	}
	for idx, v := range slice {
		fmt.Printf("%p \n", &v)
		fmt.Println(idx, v)
	}
}

//2.4 切片增长
func AppendSlice() {
	slice := []string{"cui", "jia", "bin", "qie"}
	fmt.Println("切片容量：", cap(slice))
	slice = append(slice, "新", "添", "加")
	fmt.Println("切片容量：", cap(slice))
	for idx, v := range slice {
		fmt.Println(idx, v)
	}
}

//2.5 切片删除元素
func RemoveSlice() {

	slice := []string{"cui", "jia", "bin", "qie", "新", "添", "加"}
	index := 4
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("after delete", slice)

	//插入元素
	rear := append([]string{}, slice[index:]...)
	slice = append(slice[0:index], "inserted")
	slice = append(slice, rear...)
	fmt.Println("after insert", slice)
}

//3.1 map的操作 声明map
func StateMap() {
	dict := make(map[string]int)
	dict["cui"] = 2
	dict["jia"] = 5

	fmt.Println("StateMap dict \n", dict)

	// 通过字面值创建
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println("StateMap dict \n", dict2)

}
