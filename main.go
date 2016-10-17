package main

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

//返回结果
type result struct {
	path string
	sum  []byte
	size int64
	err  error
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	c := make(chan result)
	errc := make(chan error, 1)
	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			wg.Add(1)
			go func() {
				file, err := os.Open(path)
				fileInfo, _ := os.Stat(path)
				h := sha1.New()
				io.Copy(h, file)

				select {
				case c <- result{path, h.Sum(nil), fileInfo.Size(), err}:
				case <-done:
				}
				wg.Done()
			}()
			select {
			case <-done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})

		go func() {
			wg.Wait()
			close(c)
		}()
		errc <- err
	}()
	return c, errc
}

func SHA1All(root string, result string) error {
	done := make(chan struct{})
	defer close(done)
	c, errc := sumFiles(done, root)
	if err := <-errc; err != nil {
		return err
	}

	f, err := os.Create(result)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	for r := range c {
		if r.err != nil {
			return r.err
		}
		context := r.path + "," + ByteToHex(r.sum) + "," + strconv.FormatInt(r.size, 10) + "byte"
		if _, err := w.WriteString(context + "\r\n"); err != nil {
			return err
		}
	}
	w.Flush()

	return nil
}

func ByteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {

		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}

	return buffer.String()
}

func main() {
	err := SHA1All("E:/File/图书/golang/", "E:/File/图书/golang/result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}


