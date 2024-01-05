package api

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	types "rw-file/types"
)

func HandleBinaryW() {
	fileName := "./files/f-binary.bin"
	filePtr, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建bin文件失败", err.Error())
		return
	}
	defer filePtr.Close()

	for i := 1; i <= 10; i++ {
		info := types.Website2{
			int32(i),
		}

		var binBuf bytes.Buffer
		// 将内容写入Buffer
		err = binary.Write(&binBuf, binary.LittleEndian, info)
		if err != nil {
			fmt.Println("写入Buffer失败", err.Error())
			continue
		}
		// 将Buffer写入文件
		b := binBuf.Bytes()
		_, err = filePtr.Write(b)
		if err != nil {
			fmt.Println("写入bin文件失败", err.Error())
			continue
		}
	}

	fmt.Println("写入bin完成")
}

func HandleBinaryR() {
	fileName := "./files/f-binary.bin"
	filePtr, err := os.Open(fileName)
	if err != nil {
		fmt.Println("bin文件打开失败", err.Error())
		return
	}
	defer filePtr.Close()

	m := types.Website2{}
	for i := 1; i <= 10; i++ {
		data := readNextBytes(filePtr, 4) // int32位为4个字节
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m)
		if err != nil {
			fmt.Println("二进制文件读取失败", err.Error())
			return
		}
		fmt.Println("第", i, "个值为：", m)
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("bin解码失败", err)
	}
	return bytes
}
