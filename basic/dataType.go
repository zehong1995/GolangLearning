package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	// 基本声明
	// var 变量名 变量类型
	var a bool
	// var b int
	// var c string
	// var d []float32
	// var e func() bool
	var f struct {
		x int
	}
	fmt.Println(a, f.x)

	// bool 类型无法参与数值运算，也无法与其他类型进行转换
	// fmt.Println(int(a) * 2)

	// 全写：var 变量名 变量类型 = 表达式
	// 简写：var 变量名 = 表达式
	// 根据表达式自动推到变量名
	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17
	// 类型需要统一 var damage = damageRate * (attack - defence)
	var damage = damageRate * float32(attack-defence)
	fmt.Println(damage)

	// 再简写
	// hp := 10 等价于 var hp int = 10
	// hp必须是一个没有声明过的变量
	hp1 := 10
	fmt.Println(hp1)
	// 短变量声明和赋值中至少有一个新声明即可
	hp1, hp2 := 3, 4
	// 匿名变量
	_ = hp2

	// 字符类型
	var ch1 byte = 'a' // byte /unit8
	var ch2 rune = '你' // UTF-8/int32
	fmt.Printf("%d %T\n", ch1, ch1)
	fmt.Printf("%d %T\n", ch2, ch2)

	// 切片：能动态分配的空间，可变长度的序列
	dv := make([]int, 3)
	fmt.Println(dv[0])

	// 模拟枚举
	// 给myType写一个String()函数打印字符串值
	const (
		id1 myType = iota
		id2
		id3
		id4
	)
	fmt.Printf("%s %d %d %d\n", id1, id2, id3, id4)

	// 类型别名&类型定义
	type DefineInt int  // 类型定义
	type AliasInt = int // 类型别名
	var int1 DefineInt
	var int2 AliasInt
	fmt.Printf("DefineInt type: %T\n", int1)
	fmt.Printf("AliasInt type:  %T\n", int2)
}

type myType int

func (c myType) String() string {
	if c == 0 {
		return "id1"
	} else if c == 1 {
		return "id2"
	} else if c == 2 {
		return "id3"
	} else if c == 3 {
		return "id4"
	} else {
		return "NULL"
	}
}
