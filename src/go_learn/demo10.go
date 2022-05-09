package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	//fmt.Printf("%T\n", timeObj)
	fmt.Println(timeObj)
	Year := timeObj.Year()
	Month := timeObj.Month()
	Day := timeObj.Day()
	Hour := timeObj.Hour()
	Minute := timeObj.Minute()
	Sec := timeObj.Second()
	//时间戳
	UnixTime := timeObj.Unix()
	fmt.Printf("year:%v month:%v day:%v hour:%v minute:%v second:%v\n", Year, Month, Day, Hour, Minute, Sec)
	fmt.Printf("时间戳:%v\n", UnixTime)

	//时间戳转换成日期格式
	timeObj2 := time.Unix(1645063867, 0)
	var str = timeObj2.Format("2006-01-02 15:04:05")
	fmt.Println(str)

	//格式化日期字符串转换成时间戳存储到数据库，在实际项目中数据库一般存储时间戳
	var timestr = "2022-02-17 10:11:07"
	var tmp = "2006-01-02 15:04:05" //模板要与实际时间格式保持一致
	timeObj3, _ := time.ParseInLocation(tmp, timestr, time.Local)
	//timeObj4, _ := time.Parse(tmp, timestr)
	fmt.Println(timeObj3.Unix())
	//fmt.Println(timeObj4)

	//时间操作函数
	timeObj4 := timeObj3.Add(time.Hour)
	fmt.Println(timeObj3)
	fmt.Println(timeObj4)

	//定时器
	ticker := time.NewTicker(time.Second) //创建一个间隔一秒触发一次的定时器
	//fmt.Println(ticker)
	var n = 5
	for t := range ticker.C {
		n--
		fmt.Println(t)
		if n < 0 {
			ticker.Stop()
			break
		}
	}

	for {
		time.Sleep(time.Second)
		fmt.Println("执行定时任务...")
	}
}