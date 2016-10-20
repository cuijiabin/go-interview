package mydemo

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//1.通过路径获取文件
func OpenFile(path string) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
}

//2.全部读取文件
func ReadAll(path string) {
	fi, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	fmt.Println(string(fi))
}

//3.自定义缓冲切片
func BufRead(path string) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	buf := make([]byte, 1024)
	chunks := make([]byte, 1024, 1024)

	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
}

//4.使用 bufio
func BFIOReader(path string) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	buf := make([]byte, 1024)
	chunks := make([]byte, 1024, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
}

//文件的写入 都有哪些方式? 新建 删除 覆盖追加 修改
//创建文件
func CreateFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println(path + " 创建成功")
}

//判断文件是否存在
func CheckExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	if exist {
		fmt.Println(path + " 存在")
	} else {
		fmt.Println(path + " 不存在")
	}
	return exist
}

//文件写入 io.WriteString ioutil.WriteFile File(Write,WriteString) bufio.NewWriter
func IoWrite(path string, content string, wway int) {
	fi, err := os.OpenFile(path, os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	switch wway {
	case 1:
		if _, err := io.WriteString(fi, content); err != nil {
			panic(err)
		}
	case 2:
		if err := ioutil.WriteFile(path, []byte(content), 0666); err != nil {
			panic(err)
		}
	case 3:
		if _, err := fi.WriteString(content); err != nil {
			panic(err)
		}
	case 4:
		w := bufio.NewWriter(fi)
		if _, err := w.WriteString(content); err != nil {
			panic(err)
		}
		w.Flush()
	}
	fmt.Println("文件写入成功")
	fi.Close()
}
