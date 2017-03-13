package classpath

import "os"
import "strings"

// :(linux/unix) or ;(windows) 路径分离符号
const pathListSeparator = string(os.PathListSeparator)

/**
类路径接口
*/
type Entry interface {
	// className: fully/qualified/ClassName.class
	readClass(className string) ([]byte, Entry, error)
	String() string
}

/**
类路径构造方法
*/
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		//合成类路径
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		//通配符路径
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		//归档文件路径
		return newZipEntry(path)
	}
	//文件夹路径
	return newDirEntry(path)
}
