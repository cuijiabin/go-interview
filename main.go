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
//统计需要遍历的文件夹下面的数量
var fileCount = 0
var rootPath string
var resultPath string

//默认匹配所有文件
var regStr = "(.*)"

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

	argNum := len(os.Args)
	rootPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	rootPath = strings.Replace(rootPath, "\\", "/", -1)
	resultPath = rootPath + "result.txt"

	if argNum >= 4 {
		rootPath = os.Args[1]
		resultPath = os.Args[2]
		regStr = os.Args[3]
	} else if argNum == 3 {
		rootPath = os.Args[1]
		resultPath = os.Args[2]
	} else if argNum == 2 {
		rootPath = os.Args[1]
	}

	CountFile(rootPath)
	fmt.Println("文件数量：", fileCount)

	f, err := os.Create(resultPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	c := make(chan string, 20)
	go GetFilePath(rootPath, c)

	tick := time.Tick(1e8)
	readCount := 0
	for {
		select {
		case <-tick:
			w.Flush()
			if readCount >= fileCount {
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
			readCount += 1
		}
	}
}

//统计文件夹下文件数量
func CountFile(root string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if strings.EqualFold(path, resultPath) {
			return nil
		}
		match, _ := regexp.MatchString(regStr, info.Name())
		if match {
			fileCount++
		}
		return nil
	})
}

func GetFilePath(root string, c chan string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if strings.EqualFold(path, resultPath) {
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
			blockSize = blockSize
		}
		buf := make([]byte, blockSize)
		f.Read(buf)
		io.WriteString(hash, string(buf))
	}

	return &result{path, hash.Sum(nil), fileSize}
}
