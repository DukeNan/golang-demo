package main

import (
	"fmt"
	"io"
	"os"
)

// CopyFile 拷贝文件函数
func copy_file(dstName, srcName string) (written int64, err error) {
	// 已读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 已创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) // 调用io.Copy()拷贝内容
}
