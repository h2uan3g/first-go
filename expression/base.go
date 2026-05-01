package expression

import (
	"fmt"
	"unsafe"
)

// go 命令
// go get xxx   下载
// go mod tidy  整理依赖

// 版本更新
// go install golang.org/dl/go1.21.5@latest
// go1.26.2 download
// 		eg: mv /Users/hj/sdk/go1.26.2 /Users/hj/develop/golang
// 			mv go1.26.2 go
// mv $GOROOT-NEW $GOROOT
//
//
// 指定环境打包
// CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o download download
// CGO_ENABLED=0  GOOS=windows  GOARCH=amd64 go build -o face face
// 挂起运行
// sudo nohup ./download > app.log 2>&1 &
//

// - 算术 `+` `-` `*` `/` `%` `++`  `—`
// - 比较 `==` `≠` `>` `<` `≥` `≤`
// - 赋值 `=` `+=` `-=` `*=` `/=` `%=` `>≥` `<≤` `&=` `|=` `^=`
// - 逻辑 `&&` `||` `!`
// - 位运算 `&` `|` `^` `<<` `>>`
// - 其他 `&` `*` (取址、解引用)

// 变量、常量
func VarConstExp() {

	// a := 'a'。     // 简短声明
	// var a int = 10 // 完整

	// 多变量
	var (
		a int
		b string
	)
	fmt.Println(a, b)

	// 常量需要赋初始值
	const c int32 = 10
	fmt.Println(unsafe.Sizeof(c))

	// iota 枚举 偏移量
	const (
		d = iota
		e
		f
	)
	fmt.Println(d, e, f)
}

// 分支
func ConditionalExp() {
	// if 一般形式
	if true {
		fmt.Println("true")
	}

	// if...else
	if false {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// if...else if...else
	if false {
		fmt.Println("true")
	} else if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// 1. 不用 break 语句，Go 的 switch 会自动终止
	// 2. 穿透使用 fallthrough 关键字

	word := "Gophers"
	// 值判断
	switch size := len(word); size {
	case 1, 2, 3, 4: // case 后可以接多个值
		fmt.Println(word, "is a short word!")
	case 5:
		wordLen := len(word)
		fmt.Println(word, "is exactly the right length:", wordLen)
	default: // 补充其他情况 default
		fmt.Println(word, "is a long word!")
	}

	// 表达式判断
	switch {
	case len(word) < 5:
		fmt.Println(word, "is a short word!")
	case len(word) == 5:
		fmt.Println(word, "is exactly the right length!")
	_: // 补充其他情况 _
		fmt.Println(word, "is a long word!")
	}

	// 类型判断
	var i any = 123
	switch v := i.(type) {
	case int:
		fmt.Printf("i is an int, value: %d\n", v)
	case string:
		fmt.Printf("i is a string, value: %s\n", v)
	default:
		fmt.Printf("i is of a different type: %T\n", v)
	}
}

// 循环
// - for
func ForExp() {
	// for 一般形式 fori
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 模拟 while
	// x := 10
	// for x > 1 {
	// 	x--
	// 	fmt.Println(x)
	// }

	// for...range
	for i := range 10 {
		fmt.Println(i)
	}

	// for...range (忽略索引)
	for range 10 {
		fmt.Println("*")
	}
}

// 一般函数
func func_example(a, b int) int {
	return a + b
}

// 表达式
func StatementExp() {}

// 函数
// defer
func FuncExp() {
	// 函数表达式
	// 1 参数 (值传递)
	//    - 可变参数 ...type
	// 2.多返回值
	//    - ([name1] type1,[name2] type2) 命名可选
	// 3.函数类型可以声明、赋值
	// 	  - type HandlerFunc func(ResponseWriter, *Request)
	//    - var x = func () {}
	var func_add = func(a ...int) int {
		sum := 0
		for _, v := range a {
			sum += v
		}
		return sum
	}

	func_add(1, 2)
}

func ponit_example(g *int) {
	i := 10
	// g = &i // 直接指针赋值无效
	*g = i // 非空 指针元素赋值有效
}

// 错误处理
func ErrrorExp() {
	// 错误恢复
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
	// panic("pass")

	//
	// 声明错误
	// errors.New("")

	// 自定义error
	// func (e Myerrer) Error() string {
	// }

	// fmt.Errorf (装包)
	// errors.Is()
	// errors.As()
}
