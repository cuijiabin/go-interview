package classpath

import "io/ioutil"
import (
	"fmt"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

//文件夹路径
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("新建文件路径 绝对路径；", absDir)
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	fmt.Println("文件路径,读取文件名：", fileName)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
