package main

import (
	api "test/api"
)

func main() {
	// e := echo.New()

	// e.GET("/", api.HelloWorld)
	// go e.Start(":1323")

	// go api.Running()

	// api.TestChannel()
	// 接受命令行输入
	// var input string
	// fmt.Scanln(&input)

	// fmt.Println("my input: ", input)

	// api.PlayBall()

	// api.HandleReflect()
	// api.HandleReflectStruct()
	api.HandleInject()
}
