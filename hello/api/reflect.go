package api

import (
	"fmt"
	"reflect"

	"github.com/codegangsta/inject"
)

func HandleReflect() {
	// 声明整型变量并赋值
	var a int = 1024

	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)

	// 获取interface{}类型的值，通过类型断言转换
	var getA int = valueOfA.Interface().(int)

	// 获取64位的值，强制类型转换为int类型
	var getA2 int = int(valueOfA.Int())

	// 获取64位的值，通过类型断言转换
	// var getA3 int = valueOfA.Int().(int)	// 没有类型断言方法

	fmt.Println(getA, getA2)
}

// 定义结构体
type dummy struct {
	a int
	b string

	// 嵌入字段
	float32
	bool

	next *dummy
}

func HandleReflectStruct() {
	// 值包装结构体
	d := reflect.ValueOf(dummy{
		next: &dummy{
			next: &dummy{},
		},
	})

	// 根据索引查找值中，next字段的int字段的值
	fmt.Println("FieldByIndex([] int{4, 0}).Type(): ", d.FieldByIndex([]int{4, 4, 0}).Type())
	fmt.Println("FieldByIndex([] int{4, 0}).Type(): ", d.FieldByIndex([]int{4, 4, 1}).Type())
	fmt.Println("FieldByIndex([] int{4, 0}).Type(): ", d.FieldByIndex([]int{4, 4, 2}).Type())

	// fmt.Println("FieldByIndex([] int{1, 0}).Type(): ", d.FieldByIndex([]int{1, 0}).Type())
}

type SpecialString interface{}

type TestStruct struct {
	Name   string `inject`
	Nick   []byte
	Gender SpecialString `inject`
	uid    int           `inject`
	Age    int           `inject`
}

func HandleInject() {
	s := TestStruct{}

	inj := inject.New()
	inj.Map("张三")
	inj.MapTo("男", (*SpecialString)(nil))
	inj2 := inject.New()
	inj2.Map(26)
	inj.SetParent(inj2)
	inj.Apply(&s)
	fmt.Println("s.Name =", s.Name)
	fmt.Println("s.Gender =", s.Gender)
	fmt.Println("s.Age =", s.Age)
}
