package mydemo

import (
	"strconv"
	"strings"
)

func GenIp(ip string) []string {
	slice := [][]string{}
	slice = append(slice, []string{ip})
	s := GenerateIp(slice, 4)

	resultSlice := []string{}
	for _, v := range s {
		if JudgeLegal(v) {
			var genIp string
			for i := 0; i < len(v); i++ {
				if i == 3 {
					genIp += v[i]
				} else {
					genIp += v[i] + "."
				}
			}
			resultSlice = append(resultSlice, genIp)
		}
	}
	return resultSlice

}

//把字符串切成数组列表
func Split(str string) [][]string {
	slice := [][]string{}
	for idx := 1; idx < len(str); idx++ {
		slice = append(slice, []string{str[0:idx], str[idx:]})
	}
	return slice
}

//递归切分
func GenerateIp(strList [][]string, num int) [][]string {
	if num == 1 {
		return strList
	}

	slice := [][]string{}
	for _, v := range strList {
		leftSlice := v[0 : len(v)-1]
		next := v[len(v)-1]
		subSlice := Split(next)

		for _, val := range subSlice {
			var tmpSlice []string
			tmpSlice = append(tmpSlice, leftSlice...)
			tmpSlice = append(tmpSlice, val...)
			slice = append(slice, tmpSlice)
		}
	}

	return GenerateIp(slice, num-1)
}

//判断是否合法
func JudgeLegal(list []string) bool {
	for _, v := range list {
		if strings.HasPrefix(v, "0") {
			return false
		}
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			panic(err)
		}
		if i < 0 || i > 255 {
			return false
		}

	}
	return true
}
