package expression

import (
	"fmt"
	"unsafe"
)

// struct 支持嵌套
type T struct {
	Field1 string
	Field2 int
}

// 类型实现一个接口定义的所有方法，它就自动实现该接口
type Method interface {
	MethodName()
}

// receiver参数的基类型本身不能是指针类型或接口类型
func (t *T) MethodName() {
	// 方法体
}

func StructExp() {
	// var t T
	// var t = T{}
	t := T{}
	fmt.Println(unsafe.Sizeof(t))

	// 匿名struct
	// f := struct {
	// 	Name string `json:"name"`
	// 	Age  int    `json:"age"`
	// }{}
	// err := json.Unmarshal([]byte(`{"name":"Bob","age":30}`), &f)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(f)

}

// 接口
//   - 接口 fake nil
func InterfaceExp() {

}

// 接受者
// 调用方法没区别
// 实现接口时有区别
// 对于类型*T，其自动实现了类型T实现的方法；
// 而对于类型T，其没有自动实现*T实现的方法；
// 类型T可以调用类型*T的方法只是go的一个语法糖

// 接口类型非nil
