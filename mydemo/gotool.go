package mydemo

import (
	"fmt"
	"time"
)

/*
格式化时间用 t.Format
解析时间用 time.Parse
*/
func Now() {
	now := time.Now()
	fmt.Println(now)

	secs := now.Unix() //精确到秒
	fmt.Println(secs)

	nanos := now.UnixNano()
	millis := nanos / 1000000
	fmt.Println(millis)

	fmt.Println(nanos)

	// 反过来，你也可以将一个整数秒数或者微秒数转换
	// 为对应的时间
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(now.Format("02/01/2006 15:04:05 PM"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	tm, _ := time.Parse("01/02/2006", "02/08/2015")
	fmt.Println(tm.Unix())
}
