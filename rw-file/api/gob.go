package api

import (
	"encoding/gob"
	"fmt"
	"os"

	config "rw-file/config"
	types "rw-file/types"
)

func HandleGobW() {
	var fileName string = "./files/f-gob.gob"

	// 创建文件
	filePtr, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("创建gob文件失败", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建gob句柄
	gobHandler := gob.NewEncoder(filePtr)
	err = gobHandler.Encode(config.Info)
	if err != nil {
		fmt.Println("写入gob失败", err.Error())
	} else {
		fmt.Println("写入gob成功")
	}
}

func HandleGobR() {
	var fileName string = "./files/f-gob.gob"

	// 打开文件
	filePtr, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开gob文件失败", err.Error())
		return
	}
	defer filePtr.Close()

	var info []types.Website
	// 创建gob句柄
	gobHandler := gob.NewDecoder(filePtr)
	err = gobHandler.Decode(&info)
	if err != nil {
		fmt.Println("读取gob失败", err.Error())
	} else {
		fmt.Println("读取gob成功")
		fmt.Println(info)
	}
}
