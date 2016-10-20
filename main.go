package main

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//文件块拆分
const fileChunk = 512 * 1024 * 1024 // 512M

//返回结果
type result struct {
	filePath  string
	sha1Value []byte
	fileSize  int64
}

//byte数组转16进制字符串
func byteToHex(data []byte) string {
	buf := new(bytes.Buffer)
	for _, b := range data {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buf.WriteString("0")
		}
		buf.WriteString(s)
	}
	return buf.String()
}

func main() {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = strings.Replace(dir, "\\\\", "/", -1)
	dir = strings.Replace(dir, "\\", "/", -1)
	var resultName = dir + "/result.txt"
	var regStr = "(.*)"

	fmt.Println("请输入分析目录，结果文件，过滤正则，用空格分开；默认为当前目录")
	fmt.Scanln(&dir, &resultName, &regStr)
	_, err := os.Stat(dir)
	if err != nil && err.Error() == os.ErrNotExist.Error() {
		fmt.Println("分析目录不正确")
		return
	}
	resultName = strings.Replace(resultName, "\\\\", "/", -1)
	resultName = strings.Replace(resultName, "\\", "/", -1)

	fmt.Printf("遍历目录 %s; 结果文件为 %s; 过滤正则 %s\n", dir, resultName, regStr)
	WalkAanlyseFile(dir, resultName, regStr)

}

func WalkAanlyseFile(dir string, resultName string, regStr string) {

	f, err := os.Create(resultName)
	if err != nil {
		fmt.Println("结果文件不合法")
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	c := make(chan string, 20)
	go GetFilePath(dir, c, resultName, regStr)

	tick := time.Tick(3 * time.Second)

	for {
		select {
		case <-tick:
			w.Flush()
			if len(c) < 1 {
				fmt.Println("文件统计结束")
				return
			}
		case filePath := <-c:
			res := AnalyseFile(filePath)
			context := strings.Join([]string{res.filePath, byteToHex(res.sha1Value), strconv.FormatInt(res.fileSize, 10)}, ", ")
			fmt.Println(context)
			if _, err := w.WriteString(context + "\r\n"); err != nil {
				return
			}
		}
	}
}

func GetFilePath(root string, c chan string, resultName string, regStr string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if strings.EqualFold(path, resultName) {
			return nil
		}
		match, _ := regexp.MatchString(regStr, info.Name())
		if match {
			c <- path
		}
		return nil
	})
}

func AnalyseFile(path string) *result {
	f, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		panic(err.Error())
	}

	fileSize := info.Size()
	blocks := fileSize / fileChunk
	hash := sha1.New()

	for i := int64(0); i <= blocks; i++ {
		blockSize := fileSize - i*fileChunk
		if fileChunk < blockSize {
			blockSize = fileChunk
		}
		buf := make([]byte, blockSize)
		f.Read(buf)
		io.WriteString(hash, string(buf))
	}

	return &result{path, hash.Sum(nil), fileSize}
}
