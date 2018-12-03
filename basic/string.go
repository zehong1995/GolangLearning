package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 计算长度
	str1 := "hello world!"                    // string 12 == 1*12
	str2 := "你好 世界！"                          // string 16 == 1*1+3*5
	fmt.Println(len(str1), len(str2))         // 12 16
	fmt.Println(utf8.RuneCountInString(str2)) // 6

	// asc2码遍历用for遍历下标
	for i := 0; i < len(str1); i++ {
		fmt.Printf("ascii: %c %d\n", str1[i], str1[i])
	}
	// Unicode字符遍历用for range
	for _, s := range str2 {
		fmt.Printf("Unicode: %c\t%d\n", s, s)
	}
	// 搜索字串
	str3 := "狼来了，狼跑了"
	comma := strings.Index(str3, "，")
	// 以"，"开始的字串中搜索子串"狼"
	pos := strings.Index(str3[comma:], "狼")
	// 结果：9 3 狼跑了
	fmt.Println(comma, pos, str3[comma+pos:])

	// Go语言字符串不可变
	// 修改字符串：转换为[]byte进行修改
	// []byte和string可以通过强制类型转换互转
	str4 := "Heros never die"
	angleBytes := []byte(str4)
	for i := 5; i <= 10; i++ {
		angleBytes[i] = ' '
	}
	str4 = string(angleBytes)

	// 高效字符串连接
	str5 := "abc"
	str6 := "def"
	var strBuilder bytes.Buffer
	strBuilder.WriteString(str5)
	strBuilder.WriteString(str6)
	str7 := strBuilder.String()
	fmt.Println(str7) // abcdef

	// base64编码
	msg := "zehong1995@gmail.com"
	// 编码
	encodeMsg := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encodeMsg)
	// 解码
	decodeMsg, err := base64.StdEncoding.DecodeString(encodeMsg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(decodeMsg))
	}
}
