package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	logFile, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志……", v)
	log.SetPrefix("[日志前缀]")
	log.Fatalln("这是一条会触发fatal的日志")
	log.Panicln("这是一条会触发panic的日志……")
}
