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
}
