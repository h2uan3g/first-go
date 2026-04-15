package expression

import (
	"fmt"
	"reflect"
	"strings"
)

// 整数类型
// ┌────────────┬──────────────────────┬──────────┬───────────────────┐
// │ 类型        │ 说明                  │ 大小      │ 示例              │
// ├────────────┼────────────────────── ┼──────────┼───────────────────┤
// │ int        │ 有符号整数(平台相关)     │ 32或64位  │ var i int = 42    │
// │ int8       │ 8位有符号整数        	│ 8位      │ var i int8 = 42   │
// │ int16      │ 16位有符号整数        	│ 16位     │ var i int16 = 42  │
// │ int32/rune │ 32位有符号整数        	│ 32位     │ var r rune = 'A'  │
// │ int64      │ 64位有符号整数        	│ 64位     │ var i int64 = 42  │
// │ uint       │ 无符号整数(平台相关)     │ 32或64位  │ var i uint = 42   │
// │ uint8/byte │ 8位无符号整数        	│ 8位      │ var b byte = 'A'  │
// │ uint16     │ 16位无符号整数        	│ 16位     │ var i uint16 = 42 │
// │ uint32     │ 32位无符号整数        	│ 32位     │ var i uint32 = 42 │
// │ uint64     │ 64位无符号整数        	│ 64位     │ var i uint64 = 42 │
// └────────────┴──────────────────────┴──────────┴───────────────────┘

// 浮点类
// ┌─────────┬─────────────────┬──────┬──────────────────────┐
// │ 类型     │ 说明            │ 大小  │ 示例                 │
// ├─────────┼─────────────────┼──────┼──────────────────────┤
// │ float32 │ IEEE-754 浮点数  │ 32位  │ var f float32 = 3.14 │
// │ float64 │ IEEE-754 浮点数  │ 64位  │ var f float64 = 3.14 │
// └─────────┴─────────────────┴──────┴──────────────────────┘

// 复数类型
// ┌────────────┬──────────────────────┬───────┬───────────────────────────┐
// │ 类型       │ 说明                 │ 大小  │ 示例                      │
// ├────────────┼──────────────────────┼───────┼───────────────────────────┤
// │ complex64  │ 实部虚部均为 float32 │ 64位  │ var c complex64 = 1 + 2i  │
// │ complex128 │ 实部虚部均为 float64 │ 128位 │ var c complex128 = 1 + 2i │
// └────────────┴──────────────────────┴───────┴───────────────────────────┘

//  布尔和字符串类型
// ┌────────┬────────────────┬─────────────┬────────────────────────┐
// │ 类型   │ 说明           │ 可能值      │ 示例                   │
// ├────────┼────────────────┼─────────────┼────────────────────────┤
// │ bool   │ 布尔类型       │ true, false │ var b bool = true      │
// │ string │ 不可变字节序列 │ 任意字符    │ var s string = "hello" │
// └────────┴────────────────┴─────────────┴────────────────────────┘

func StrExp() {
	res := strings.Map(func(r rune) rune {
		return r + 32
	}, "HELLE")
	fmt.Println(res)
}

// 复合数据类
// ┌───────────┬─────────────┬─────────────┬──────────────────────────────┐
// │ 类型      │ 说明        │ 特性        │ 示例                         │
// ├───────────┼─────────────┼─────────────┼──────────────────────────────┤
// │ 数组(Array) │ 固定长度... │ 长度固定... │ var arr [5]int               │
// │ 切片(Slice) │ 动态长度... │ 动态扩容... │ var slice []int = []int{1,2} │
// │ 映射(Map)   │ 键值对无... │ 哈希表实现  │ var m map[string]int         │
// │ 结构体(Struct) │ 自定义复... │ 字段集合    │ `type Person struct{Name ... │
// └───────────┴─────────────┴─────────────┴──────────────────────────────┘

func ArrayAndSliceExp() {
	// Array 类型长度同, 可比较
	// var arr [2]int
	// var arr = [2]int{1, 2}
	// var arr = [...]int{1, 2}
	// var arr = [2]int{1} // 默认初值
	// fmt.Println(unsafe.Sizeof(arr))

	// Slice 类型可变
	// 1. 直接声明
	// var nums = []int{1, 2, 3, 4, 5, 6} // 默认nil  零值可用
	// 2. make 声明
	// nums := make([]byte, 6, 10) // 其中10为cap值，即底层数组长度，6为切片的初始长度
	// 3. 从数组切片
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums := arr[3:7:9]

	// append
	nums = append(nums, 1, 2)

	// copy 深拷贝

}

func MapExp() {
	// var m map[string]int // 一个map[string]int类型的变量   零值不可用

	// m1 := make(map[int]string)    // 未指定初始容量
	// m2 := make(map[int]string, 8) // 指定初始容量为8

	m := map[int]string{}

	if _, ok := m[1]; ok {
	}
	// 删除 delete

	delete(m, 1)
}

// 引用类型
// ┌────────────┬───────────┬──────────────┬───────────────────────────────┐
// │ 类型        │ 说明      │ 用途         │ 示例                          │
// ├────────────┼───────────┼──────────────┼───────────────────────────────┤
// │ 指针(Point) │ 内存地址  │ 修改变量,... │ var p *int = &i               │
// │ 通道(Chanel) │ gorout... │ 并发通信     │ ch := make(chan int)          │
// │ 函数(Func)  │ 一等公民  │ 参数/返回... │ var f func(int) int           │
// │ 接口(Interface) │ 方法集... │ 多态         │ `type Stringer interface{S... │
// │ 切片(Slice)  │ 动态数... │ 灵活操作数组 │ slice := arr[1:3]             │
// │ 映射(Map)  │ 键值对... │ 快速查找     │ m := make(map[string]int)     │
// └────────────┴───────────┴──────────────┴───────────────────────────────┘

func TypeAssert() {
	// 类型断言
	// 方式一
	// var i any = 1
	// if v, ok := i.(int); ok {
	// 	fmt.Println("i is an int, value:", v)
	// }

	// 方式二
	// var i any = []string{"a", "b", "c"}
	// switch t := i.(type) {
	// case []string:
	// 	fmt.Println("i is an slice, value:", t)
	// case map[int]string:
	// 	fmt.Println("i is an array, value:", t)
	// default:
	// 	fmt.Printf("default: %T \n", t)
	// 	return
	// }

	// 类型转换
	// v1 := T(v2)

	// strconv.Itoa()、strconv.Atoi()、
	// strconv.ParseXxx()
	// strconv.FormatXxx()
}

// 指针
func PointExp() {
	// 指针大小
	// var a int = 10
	// var p *int = &a
	// fmt.Println(unsafe.Sizeof(a))
	// fmt.Println(unsafe.Sizeof(p))

	// 指针创建
	// 	- &
	//  - new
	// var a = new(int)
	// fmt.Println(a)
	// fmt.Println(*a)

	// 结构体指针
	// type Point struct {
	// 	x *string
	// }
	// s := ""
	// p := Point{
	// 	x: &s,
	// }
	// fmt.Println(p)

	// 函数传nil 重新赋值无效
	// var g *int
	// var g = new(int)
	// ponit_example(g)
	// fmt.Println(*g)

	//

}

// 接口
//   - 接口 fake nil
func InterfaceExp() {

}

type Foo struct {
	A int `myTag:"value"`
}

// Reflect
func ReflectExp() {
	// reflect.TypeOf
	// it := reflect.TypeFor[int]()
	// ipt := reflect.TypeFor[*int]()
	// fpt := reflect.TypeFor[*Foo]()

	// fmt.Println(it.Name())
	// fmt.Println(ipt.Name())
	// fmt.Println(fpt.Name())

	// fmt.Println(it.Kind())
	// fmt.Println(ipt.Kind())
	// fmt.Println(fpt.Kind())

	// fmt.Println(it.Elem())
	// fmt.Println(ipt.Elem())
	// fmt.Println(fpt.Elem())

	// ft := reflect.TypeFor[Foo]()
	// for curFiled := range ft.Fields() {
	// 	fmt.Println(curFiled.Name, curFiled.Type, curFiled.Tag)
	// }

	// reflect.ValueOf
	// s := []string{"a", "b"}
	// v := reflect.ValueOf(s)
	// s2 := v.Interface().([]string)
	// fmt.Println(v, s2)

	// 修改值
	// i := 10
	// iv := reflect.ValueOf(&i)
	// ivv := iv.Elem()
	// ivv.SetInt(20)
	// fmt.Println(iv, ivv)

	// 创建值
	var stringValue = reflect.TypeOf((*string)(nil)).Elem()
	sv := reflect.New(stringValue).Elem()
	sv.SetString("hell0")
	fmt.Println(sv)

}

// Unsafe
func UnsafeExp() {

}
