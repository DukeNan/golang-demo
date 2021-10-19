package main

import (
	"fmt"
	"time"
)

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func main() {
	now := GetNowTime()
	fmt.Printf("现在的时间是：%s\n", now)
}
