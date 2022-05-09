package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	timeStr := timeObj.Format("2006/01/02 15:04:05")
	fmt.Println(timeStr)

	unixTime := timeObj.Unix()
	fmt.Println("时间戳:", unixTime)

	//时间戳1587880013
	var tmp = "2006/01/02 15:04:05"
	timeObj1 := time.Unix(1587880013, 0)
	timeStr1 := timeObj1.Format(tmp)
	fmt.Println(timeStr1)

	//日期字符串转换时间戳
	var timeStr2 = "2020/04/26 13:46:53"
	timeObj2, _ := time.ParseInLocation(tmp, timeStr2, time.Local)
	unixTime2 := timeObj2.Unix()
	fmt.Println("时间戳:", unixTime2)

	timeObj3 := time.Now()
	fmt.Printf("程序耗时:%d微秒\n", timeObj3.UnixMicro() - timeObj.UnixMicro())
}
