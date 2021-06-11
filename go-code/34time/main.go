package main

import (
	"fmt"
	"time"
)

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Println(year, month, day, hour, minute, second)
}

func tickDemo() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}

func main() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)
	fmt.Println("=========时间戳=============")
	timestampDemo2(1622623223000)
	fmt.Println("========时间操作============")
	later := now.Add(time.Hour * 2)
	fmt.Printf("later: %v\n", later)
	diff := later.Sub(now)
	fmt.Printf("diff: %v\n", diff)
	fmt.Println("=======定时器=============")
	// tickDemo()

}
