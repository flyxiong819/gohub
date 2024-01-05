package api

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func HandleTarW() {
	srcFile := "./files/test-f-for-tar.txt"     // 源文件
	tarf, err := os.Create("./files/f-tar.tar") // 创建一个tar文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tarf.Close()

	tarw := tar.NewWriter(tarf)
	defer tarw.Close()

	fileInfo, err := os.Stat(srcFile) // 获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}
	// 获取源文件的头部信息
	hdr, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		fmt.Println(err)
	}

	err = tarw.WriteHeader(hdr) // 写入头文件信息
	if err != nil {
		fmt.Println(err)
	}

	srcf, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	m, err := io.Copy(tarw, srcf) // 将srcFile文件中的信息写入压缩包中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tar打包完成")
	fmt.Println(m) // m是写入长度（字节数）
}

func HandleTarR() {
	file, err := os.Open("./files/f-tar.tar")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file.Close()

	handlerR := tar.NewReader(file)
	for hdr, err := handlerR.Next(); err != io.EOF; hdr, err = handlerR.Next() {
		if err != nil {
			fmt.Println(err)
			return
		}
		fileInfo := hdr.FileInfo()
		fmt.Println(fileInfo.Name())
		f, err := os.Create("./files/123" + fileInfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		_, err = io.Copy(f, handlerR)
		if err != nil {
			fmt.Println(err)
		}
	}

}
