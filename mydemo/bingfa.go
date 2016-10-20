package mydemo

/*
一些经典的并发编程问题
*/
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

func godemo() {
	go say("world")
	say("hello")
}

func gosum() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 从 c 中获取

	fmt.Println(x, y, x+y)

	cc := make(chan int, 2)
	cc <- 1
	cc <- 2
	fmt.Println(<-cc)
	fmt.Println(<-cc)
}

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func Bcount() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c > 10 {
			break
		}
	}

}

func Add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func MyGo() {
	for i := 0; i < 10; i++ {
		go Add(i, 1)
	}
	time.Sleep(500)
}
func Count1(ch chan int, i int) {
	fmt.Println("Counting")
	ch <- i
}
func MyCount() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count1(chs[i], i)
	}
	for _, ch := range chs {
		fmt.Println(<-ch)
	}

	anych := make(chan interface{}, 10)
	for i := 0; i < 10; i++ {
		anych <- "崔佳彬"
	}
	close(anych)
	x, ok := <-anych
	if ok == false {
		fmt.Println("通道已经关闭")
	} else {
		fmt.Println("x:", x)
	}
	//for chh := range anych {
	//	fmt.Println(chh)
	//}
}

func MySelect1() {
	ch := make(chan int, 1)
	//for {
	select {
	case ch <- 0:
	case ch <- 1:
	}
	i := <-ch
	fmt.Println("接收到的值：", i)
	//}
}
func MyTimeout() {
	ch := make(chan int, 1)
	//ch <- 2
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	select {
	case <-ch:
		fmt.Println("没有超时")
	case <-timeout:
		fmt.Println("超时了")
	}
}

func MyWaitGroup() {
	jobs := make(chan int)
	var wg sync.WaitGroup
	wg.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			jobs <- i
			fmt.Println("producep:", i)
		}
	}()
	go func() {
		for i := range jobs {
			fmt.Println("consume:", i)
			wg.Done()
		}
	}()
	wg.Wait()
}

func MyWaitRace() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	var wg sync.WaitGroup
	var a int = 20
	var p *int = &a
	fmt.Println(*p)
	fmt.Println(&p)
	wg.Add(10)
	fmt.Println("执行")
	wg.Wait()
	fmt.Println("执行")
}

// threads 线程标识创建线程的个数
func quicksort(nums []int, ch chan int, level int, threads int) {
	level=level*2
	if len(nums) == 1 {  ch<- nums[0]; close(ch); return }//ch<-nums[0] 表示将nums[0] 数据写到ch通道中
	if len(nums) == 0 {  close(ch); return }

	less := make([]int, 0)//
	greater := make([]int,0)
	left := nums[0] //快速排序的轴
	nums = nums[1:]

	//从左向右扫描数据 大于轴的放到greater里小于的放到less中
	for _,num_data := range nums{
		switch{
		case num_data <= left:
			less = append(less,num_data)
		case num_data > left:
			greater = append(greater,num_data)
		}
	}

	left_ch := make(chan int, len(less))
	right_ch := make(chan int, len(greater))

	if(level <= threads){
		go quicksort(less, left_ch, level, threads) //分任务
		go quicksort(greater,right_ch, level, threads)
	}else{
		quicksort(less,left_ch, level, threads)
		quicksort(greater,right_ch, level, threads)
	}

	//合并数据
	for i := range left_ch{
		ch<-i;
	}
	ch<-left
	for i := range right_ch{
		ch<-i;
	}
	close(ch)
	return
}

func MyQuickSort() {
	x := []int{3, 1, 4, 1, 5, 9, 2, 6}
	ch := make(chan int)
	go quicksort(x, ch, 0, 0) // 0 0 表示不限制线程个数
	for v := range(ch) {
		fmt.Println(v)
	}
}






