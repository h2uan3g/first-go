package expression

import (
	"fmt"
	"reflect"
)

// 通过类型信息调用函数
func hello(a, b string) {
	fmt.Println(a, b)
}

func ReflectReg() {
	x := 1

	// 接口变量转反射对象
	// rt := reflect.TypeOf(x)
	rt := reflect.TypeFor[int]()
	fmt.Printf("%v : %T \n", rt, rt)

	// xv.Name() 声明类型  xv.Kind() 底层类型
	xv := reflect.ValueOf(x)
	fmt.Printf("%v : %T \n", xv, xv)
	xr := reflect.ValueOf(&x).Elem()
	if b := xr.CanSet(); b {
		xr.SetInt(8)
		fmt.Printf("%v : %T \n", x, x)
	}

	// 获取tag

	// 通过类型信息调用函数
	h := reflect.ValueOf(hello)
	params := []reflect.Value{reflect.ValueOf("hello"), reflect.ValueOf("world")}
	h.Call(params)
}
