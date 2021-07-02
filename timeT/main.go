/*
 * @Author: your name
 * @Date: 2020-05-21 18:59:19
 * @LastEditTime: 2020-10-28 16:31:14
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /G/src/studyGo/timeT/main.go
 */
package main

import (
	"fmt"
	"time"
)

// 以下为时间转化的一些参数，固定的使用，如果自行换的话会出问题
//2006-01-02 15:04:05
// const (
// 	ANSIC       = "Mon Jan _2 15:04:05 2006"
// 	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
// 	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
// 	RFC822      = "02 Jan 06 15:04 MST"
// 	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
// 	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
// 	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
// 	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
// 	RFC3339     = "2006-01-02T15:04:05Z07:00"
// 	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
// 	Kitchen     = "3:04PM"
// 	// Handy time stamps.
// 	Stamp      = "Jan _2 15:04:05"
// 	StampMilli = "Jan _2 15:04:05.000"
// 	StampMicro = "Jan _2 15:04:05.000000"
// 	StampNano  = "Jan _2 15:04:05.000000000"
// )

// //显示当前的时间,格式"2006-01-02 15:04:05"
// time.Now().Format("2006-01-02 15:04:05")

// //显示当前的时间,格式"2006-01-02 15:04:05.232"
//  time.Now().Format("2006-01-02 15:04:05.000")

//  //当前的时间戳
//  time.Now().Unix()

//  //把时间戳转换为"2006-01-02 15:04:05"
//  time.Unix(1470017531, 0).Format("2006-01-02 15:04:05")

//  //五天前的时间
//  time.Unix(time.Now().Unix()-3600*24*5, 0).Format("2006-01-02 15:04:05")

//  //"2016-11-11 15:08:05"转换为时间戳
//  tm,_:=time.ParseInLocation("2006-01-02 15:04:05", "2016-11-11 15:08:05", time.Local) //前一个参数是时间格式，后一个参数是需要转换的时间
//  fmt.Println(tm.Unix())

//  //获取下个月的时间
// 	 t := time.Now()
// 	 startTime := time.Date(t.Year(), t.Month()+1, t.Day(), 0, 0, 0, 0, t.Location()).Format("2006-01-02 15:04:05")
//  //output:2017-05-07 00:00:00

//  //从数字20171102转为时间字符串2017-11-02 00:00:00
//  date:=20171102
//  t, _ = time.Parse("20060102", strconv.Itoa(date))
//  startDate := t.Format("2006-01-02") + " 00:00:00"
//  //output: 2017-11-02 00:00:00

//  //从“2017-11-02 00:00:00”转化为“20171102”
//  t, _ := time.Parse("2006-01-02 15:04:05", "2017-11-02 00:00:00")
//  startDateInt := t.Format("20060102")

func uni() {
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)
}

func strToUtc() {
	createTime, _ := time.Parse("2006-01-02 15:04:05", "2020-06-27 23:34:38")

	fmt.Println(createTime.Unix())
}

func utcToStr() {
	fmt.Println(time.Now().Format("060102"))
}

func bijiao() {
	time1 := "2015-03-20 08:50:29"
	time2 := "2015-03-21 09:04:25"
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		fmt.Println("true")
	}
	if t2.After(t1) {
		fmt.Println(fmt.Sprintf("%v 在(After) %v之后", t2, t1))
	}
}

func ticker() {
	var tmr = time.NewTimer(time.Second)
	tmr.Reset(123)
}

func main() {
	// uni()
	bijiao()
}
457720