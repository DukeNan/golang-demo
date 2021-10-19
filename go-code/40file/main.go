package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func read_file() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	// 使用Read方法读取数据
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节的数据\n", n)
	fmt.Println(string(tmp[:n]))
}

func loop_read_file() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了！")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

func bufio_read_file() {
	file, err := os.Open("./aaa.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

func ioutil_read_file() {
	centent, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(centent))
}

// func main() {
// 	// read_file()
// 	// loop_read_file()
// 	// bufio_read_file()
// 	// ioutil_read_file()
// 	// write_file()
// 	// bufio_write_file()
// 	// ioutil_write_file()
// 	_, err := copy_file("dst.txt", "src.txt")
// 	if err != nil {
// 		fmt.Println("copy file failed, err:", err)
// 		return
// 	}
// 	fmt.Println("copy done!")
// }
