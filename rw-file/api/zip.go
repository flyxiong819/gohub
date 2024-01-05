package api

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
)

func HandleZipW() {
	// 创建一个缓冲区用来保存压缩文件内容
	buf := new(bytes.Buffer)

	// 创建一个压缩文档
	zipw := zip.NewWriter(buf)

	// 将文件加入压缩文档
	var files = []struct {
		Name, Body string
	}{
		{"Golang.txt", "http://c.biancheng.net/golang/"},
	}
	for _, file := range files {
		f, err := zipw.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}

	// 关闭压缩文档
	err := zipw.Close()
	if err != nil {
		fmt.Println(err)
	}

	// 将压缩文档内容写入文件
	f, err := os.OpenFile("./files/f-zip.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	buf.WriteTo(f)
}

func HandleZipR() {
	// 打开一个zip格式文件
	zipf, err := zip.OpenReader("./files/f-zip.zip")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer zipf.Close()

	// 迭代压缩文件中的文件，打印出文件中的内容
	for _, f := range zipf.File {
		fmt.Printf("文件名：%s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = io.CopyN(os.Stdout, rc, int64(f.UncompressedSize64))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println()
		rc.Close()
	}
}
