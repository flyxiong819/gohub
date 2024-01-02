package api

import (
	"encoding/json"
	"fmt"
	"os"

	config "rw-file/config"
	types "rw-file/types"
)

func HandleJsonW() {

	// 创建文件
	filePtr, err := os.Create("./files/f-json.json")
	if err != nil {
		fmt.Println("创建json文件失败", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建Json句柄
	jsonHandler := json.NewEncoder(filePtr)
	err = jsonHandler.Encode(config.Info)
	if err != nil {
		fmt.Println("JSON写入失败", err.Error())
	} else {
		fmt.Println("JSON写入成功")
	}
}

func HandleJsonR() {
	// 打开文件
	filePtr, err := os.Open("./files/f-json.json")
	if err != nil {
		fmt.Println("打开json文件失败", err.Error())
		return
	}
	defer filePtr.Close()

	var info []types.Website
	// 创建json句柄
	jsonHandler := json.NewDecoder(filePtr)
	err = jsonHandler.Decode(&info)
	if err != nil {
		fmt.Println("JSON读取失败", err.Error())
	} else {
		fmt.Println("JSON读取成功")
		fmt.Println(info)
	}
}
