package mydemo

import (
	"fmt"
	"qiniupkg.com/x/errors.v7"
)

func Myerror() {
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i,j: %d,%d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
	fmt.Println(errors.New("this is a error"))

	f()
	fmt.Println("Returned normally from f.")
}

func MyPanic()  {
	panic("不好了")
}

func MySelect()  {
	select {

	}
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g()
	fmt.Println("Returned normally from g.")
}

func g() {
	panic("ERROR")
}
