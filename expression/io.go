package expression

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

// 文件操作
//   - os     os.Open、os.Create、os.Rename、os.Stat
//   - io
//   - bufio
func IoExp() {
	// os
	// if len(os.Args) < 2 {
	// 	log.Fatal("no file specified")
	// }
	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// data := make([]byte, 2048)
	// for {
	// 	count, err := f.Read(data)
	// 	os.Stdout.Write(data[:count])
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.Fatal(err)
	// 		}
	// 		break
	// 	}
	// }
	//
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// write
	// buf := bufio.NewWriter(file)
	// buf.WriteString("")

	// read
	buf := bufio.NewReader(file) // 参数为 os.Stdin 读取输入
	for {
		readString, err := buf.ReadString('\n')
		if err != nil && err == io.EOF {
			return
		}
		fmt.Println(strings.TrimSpace(readString))
	}
}

// 输入输出
func InOutExp() {

	/*
		// 控制台输入
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("input: ")
		numStr, _ := reader.ReadString('\n')
		nums := strings.Fields(strings.TrimSpace(numStr))
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		fmt.Printf("%d + %d = %d \n", a, b, a+b)
	*/

	// 命令行输入 go run first 12 11
	var args = os.Args
	a, _ := strconv.Atoi(args[1])
	b, _ := strconv.Atoi(args[2])
	fmt.Printf("%d + %d = %d \n", a, b, a+b)
}

func JsonExp() {
	// json.Unmarshal
	user := struct {
		Name string `json:"Name"`
		Age  int    `json:"Age"`
		Sex  string `json:"Sex"`
	}{}
	b1 := []byte(`{"Name":"David","Age":26,"Sex":"Male"}`)
	json.Unmarshal(b1, &user)
	fmt.Println(user)

	// json.Marshal
	user.Name = "test"
	user.Age = 18
	user.Sex = "man"

	b, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))

	// 写入文件
	// json.NewEncoder()
	// 读取文件
	// json.NewDecoder()
}

func TimeExp() {
	// now := time.Now()
	// fmt.Println(now.Format(time.DateTime))

	//
	// time.Parse(time.DateTime, "2006-01-02 15:04:05")
	// time.After(time.Hour * 2)
	// time.NewTimer(time.Hour)
}

func SlogExp() {
	slog.Info("-------")
}
