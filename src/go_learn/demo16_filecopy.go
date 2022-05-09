package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//ioutil方法拷贝文件
func copy1(srcFileName string, destFileName string) (err error) {
	byteStr, err1 := ioutil.ReadFile(srcFileName)
	if err1 != nil {
		return err1
	}
	err2 := ioutil.WriteFile(destFileName, byteStr, 0666)
	if err2 != nil {
		return err2
	}
	return nil
}

//文件流的方法拷贝文件
func copy2(srcFileName string, destFileName string) (err error) {
	srcfile, err1 := os.Open(srcFileName)
	defer srcfile.Close()
	if err1 != nil {
		fmt.Println("源文件打开失败...", err1)
		return err1
	}
	destfile, err2 := os.OpenFile(destFileName, os.O_CREATE|os.O_RDWR, 0666)
	defer destfile.Close()
	if err2 != nil {
		fmt.Println("源文件打开失败...", err2)
		return err2
	}
	var tempSlice = make([]byte, 1280)
	for {
		n, err := srcfile.Read(tempSlice)
		if err ==io.EOF {
			fmt.Println("文件读取完毕...")
			break
		}
		if err != nil {
			fmt.Println("文件读取失败...")
			return err
		}
		if _, err := destfile.Write(tempSlice[:n]); err != nil {
			fmt.Println("文件写入失败...")
			return err
		}
	}
	return nil
}

func main() {
	srcFileName := "D:/文件/AxureRP-Pro9.0-Chinese.zip"
	dstFileName := "C:/AxureRP-Pro9.0-Chinese.zip"
	err := copy1(srcFileName, dstFileName)
	if err != nil {
		fmt.Println("文件复制失败", err)
		return
	}
	fmt.Println("复制文件成功")
}
