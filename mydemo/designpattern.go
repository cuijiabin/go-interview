package mydemo

import (
	"fmt"
	"math/rand"
	"time"
)

//生成器
func xrange() chan int { // xrange用来生成自增的整数
	var ch chan int = make(chan int)

	go func() { // 开出一个goroutine
		for i := 0; ; i++ {
			ch <- i // 直到信道索要数据，才把i添加进信道
		}
	}()

	return ch
}

func Pxgen() {
	generator := xrange()
	for i := 0; i < 1000; i++ { // 我们生成1000个自增的整数！
		fmt.Println(<-generator)
	}
}

//服务化
func GetNotice(user string) chan string {
	/*
	 * 此处可以查询数据库获取新消息等等..
	 */
	notifications := make(chan string)

	go func() { // 悬挂一个信道出去
		notifications <- fmt.Sprintf("Hi %s, welcome to weibo.com!", user)
	}()

	return notifications
}

//多路合并
func do_stuff(x int) int { // 一个比较耗时的事情，比如计算
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) //模拟计算
	return 100 - x                                              // 假如100-x是一个很费时的计算
}

func branch(x int) chan int { // 每个分支开出一个goroutine做计算并把计算结果流入各自信道
	ch := make(chan int)
	go func() {
		ch <- do_stuff(x)
	}()
	return ch
}

func fanIn(chs ...chan int) chan int {
	ch := make(chan int)

	for _, c := range chs {
		// 注意此处明确传值
		go func(c chan int) { ch <- <-c }(c) // 复合
	}

	return ch
}

func MergeGo() {
	result := fanIn(branch(1), branch(2), branch(3))

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}

func foo(i int) chan int {
	c := make(chan int)
	go func() {
		c <- i
		time.Sleep(time.Duration(5) * time.Duration(4-i) * time.Second)
	}()
	return c
}

func SelectDemo() {
	c1, c2, c3 := foo(1), foo(2), foo(3)

	c := make(chan int)

	go func() { // 开一个goroutine监视各个信道数据输出并收集数据到信道c
		for {
			select { // 监视c1, c2, c3的流出，并全部流入信道c
			case v1 := <-c1:
				c <- v1
			case v2 := <-c2:
				c <- v2
			case v3 := <-c3:
				c <- v3
			}
		}
	}()

	// 阻塞主线，取出信道c的数据
	for i := 0; i < 3; i++ {
		fmt.Println(<-c) // 从打印来看我们的数据输出并不是严格的1,2,3顺序
	}
}

//菊花链生成素数
func xrange2() chan int{ // 从2开始自增的整数生成器
	var ch chan int = make(chan int)

	go func() { // 开出一个goroutine
		for i := 2; ; i++ {
			ch <- i  // 直到信道索要数据，才把i添加进信道
		}
	}()

	return ch
}


func filter(in chan int, number int) chan int {
	// 输入一个整数队列，筛出是number倍数的, 不是number的倍数的放入输出队列
	// in:  输入队列
	out := make(chan int)

	go func() {
		for {
			i := <- in // 从输入中取一个

			if i % number != 0 {
				out <- i // 放入输出信道
			}
		}
	}()

	return out
}


func PrimeNumber() {
	const max = 100 // 找出100以内的所有素数
	nums := xrange2() // 初始化一个整数生成器
	number := <-nums  // 从生成器中抓一个整数(2), 作为初始化整数

	for number <= max { // number作为筛子，当筛子超过max的时候结束筛选
		fmt.Println(number) // 打印素数, 筛子即一个素数
		nums = filter(nums, number) //筛掉number的倍数
		number = <- nums  // 更新筛子
	}
}

func rand01() chan int {
	ch := make(chan int)

	go func () {
		for {
			select { //select会尝试执行各个case, 如果都可以执行，那么随机选一个执行
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	return ch
}


func GenRand() {
	generator := rand01() //初始化一个01随机生成器

	//测试，打印10个随机01
	for i := 0; i < 10; i++ {
		fmt.Println(<-generator)
	}
}