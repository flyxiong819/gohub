package api

import (
	"encoding/xml"
	"fmt"
	"os"
	config "rw-file/config"
	types "rw-file/types"
)

func HandleXmlW() {
	var fileName string = "./files/f-xml.xml"

	// 创建文件
	filePtr, err := os.Create(fileName)
	if err != nil {
		fmt.Println("XML创建文件失败", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建xml句柄
	xmlHandler := xml.NewEncoder(filePtr)
	err = xmlHandler.Encode(config.Info)
	if err != nil {
		fmt.Println("xml写入失败")
	} else {
		fmt.Println("xml写入成功")
	}
}

func HandleXmlR() {
	var fileName string = "./files/f-xml.xml"

	// 打开文件
	filePtr, err := os.Open(fileName)
	if err != nil {
		fmt.Println("XML打开失败", err.Error())
		return
	}
	defer filePtr.Close()

	var info []types.Website
	// 创建xml句柄
	xmlHandler := xml.NewDecoder(filePtr)
	err = xmlHandler.Decode(&info)
	if err != nil {
		fmt.Println("读取xml失败", err.Error())
	} else {
		fmt.Println("读取xml成功")
		fmt.Println(info)
	}

}
