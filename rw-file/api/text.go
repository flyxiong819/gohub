package api

import (
	"bufio"
	"fmt"
	"io"
	"os"

	config "rw-file/config"
)

func HandleTextW() {
	var fileName string = "./files/f-text.txt"

	// 创建新文件
	filePtr, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开text文件错误", err.Error())
		return
	}
	defer filePtr.Close()

	// 定义bufio写句柄
	writerHandler := bufio.NewWriter(filePtr)
	// 写内容
	_, err1 := writerHandler.WriteString(config.Str)
	if err1 != nil {
		fmt.Println("写text到缓存失败", err1.Error())
		return
	}
	// 因为writerHandler是带缓存的，因此在调用WriterString方法时，内容是先写入缓存的
	// 调用flush方法，将缓存的数据真正写入到文件
	err = writerHandler.Flush()
	if err != nil {
		fmt.Println("写text失败", err.Error())
	} else {
		fmt.Println("写text成功")
	}
}

func HandleTextR() {
	var fileName string = "./files/f-text.txt"

	// 打开文件
	filePtr, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开text文件错误", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建句柄，带缓冲的
	readerHandler := bufio.NewReader(filePtr)
	// 读文件
	for {
		str, err1 := readerHandler.ReadString('\n')
		if err1 == io.EOF {
			break
		}
		if err1 != nil {
			fmt.Println("读text失败", err1.Error())
		} else {
			fmt.Println("读text成功")
			fmt.Println(str)
		}
	}

}
