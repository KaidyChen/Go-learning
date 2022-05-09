package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

//1.file.Read()读取文件,文件流读取
func fileRead(fileName string) {
	//1.打开文件
	file, err := os.Open(fileName) //只读方式打开文件
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	//2.读取文件里面的的内容
	var strSlice []byte
	var tempSlice = make([]byte, 128)
	for {
		n, err := file.Read(tempSlice) //读取128字节内容写入到tempSlice
		if err == io.EOF {
			fmt.Println("读取完毕...")
			break
		}
		if err != nil {
			fmt.Println("读取失败...")
			//fmt.Println(err)
			return
		}
		//fmt.Printf("读取到了%v个字节\n", n)
		strSlice = append(strSlice, tempSlice[:n]...)//最后一次读取的时候如果文件内容不足128字节，会导致最后一个tempSlice里面含有重复的内容，所以每次只需拼接n个字节即可
	}
	fmt.Println(string(strSlice))
}

//1.file.WriteString()写入文件
func fileWrite(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败...", err)
		return
	}
	//写入文件
	for i := 0; i< 10; i++ {
		file.WriteString("这是写入的字符串数据" + strconv.Itoa(i) + "\r\n")
	}
}


//bufio读取文件,文件流读取
func bufioRead(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("读取失败...")
		return
	}

	var fileStr string
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')//表示一次读取一行
		if err == io.EOF {
			fmt.Println("读取完毕")
			fileStr += str
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fileStr += str
	}

	fmt.Println(fileStr)
}

////bufio写入文件
func bufioWrite(fileName string) {
	//1.打开文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC,0666)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败...",err)
		return
	}
	//2.创建write对象
	writer := bufio.NewWriter(file)
	//3.将数据写入缓存
	for i := 0; i< 10; i++ {
		writer.WriteString("这是bufio方法写入的数据" + strconv.Itoa(i) + "\r\n")
	}
	//4.将缓存中的内容写入文件
	writer.Flush()
	//5.关闭文件流
}

//3.ioutil读取文件，打开关闭文件的方法都封装好了只需一句话就可以读取,适用于小文件
func ioutilRead(fileName string) {
	byteStr, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("读取失败...", err)
		return
	}
	fmt.Println(string(byteStr))
}

//ioutil写入文件
func ioutilWrite(fileName string) {
	str := "这是ioutil方法写入的数据"
	err := ioutil.WriteFile(fileName, []byte(str), 0666)
	if err != nil {
		fmt.Println("写入文件失败...", err)
	}
}

func main() {
	fileName := "D:/project/go/src/test.txt"
	//fileRead(fileName)
	//bufioRead(fileName)
	//ioutilRead(fileName)
	//fileWrite(fileName)
	//bufioWrite(fileName)
	ioutilWrite(fileName)
}
