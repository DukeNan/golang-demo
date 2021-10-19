package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func write_file() {
	file, err := os.OpenFile("./aaa.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello world"
	file.Write([]byte(str))
	file.WriteString("hello 小王子")
}

func bufio_write_file() {
	file, err := os.OpenFile("./aaa.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n")
	}
	writer.Flush()
}

func ioutil_write_file() {
	str := "hell 冰冰"
	err := ioutil.WriteFile("./aaa.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
