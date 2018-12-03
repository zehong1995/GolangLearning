package main

import (
	"container/list"
	"fmt"
	"sync"
)

func main() {
	// 数组
	// var 数组变量名 [数组数量]T
	// T可是是任意基本类型，包括数组
	var a = [...]string{"hammer", "soldier", "mum"} // ...表示让编译器确定数组大小
	var b = [3]string{"hammer", "soldier", "mum"}   // 常规写法
	_ = b
	// 数组遍历
	for k, v := range a {
		fmt.Println(k, v)
	}

	// 切片slice
	// 动态分配大小的连续空间，用于快速地操作一块数据集合
	// 切片内部结构包括：地址、大小、容量
	// 切片有点像指针，但增加了大小约束内存区域，比指针更安全强大

	fmt.Println("----------slice----------")
	// 声明切片
	var numList1 []int                            // 空切片，但没有分配内存
	var numList2 = []int{}                        // 空切片，但分配了内存
	fmt.Println(numList1, numList2)               // [] []
	fmt.Println(numList1 == nil, numList2 == nil) // true false

	// 从指定范围生成切片
	var num [30]int
	for i := 0; i < 30; i++ {
		num[i] = i
	}
	fmt.Println(num[10:20])
	fmt.Println(num[20:])
	fmt.Println(num[:10])

	// 表示原有切片以及重置切片
	numList1 = []int{1, 2, 3}
	fmt.Println(len(numList1), numList1[:])
	numList1 = numList1[0:0]
	fmt.Println(len(numList1), numList1[:])

	// 使用make()函数
	// make([]T, size, cap) 类型，大小，预分配
	numList1 = make([]int, 2)
	numList2 = make([]int, 2, 10)
	fmt.Println(numList2, len(numList2))

	// 使用append扩容
	// 扩容后首地址会变，因为切片的内存是连续分布的
	var numList3 []int
	fmt.Printf("len: %d\tcap: %d\tpointer:%p\n", len(numList3), cap(numList3), numList3)
	for i := 0; i < 10; i++ {
		numList3 = append(numList3, i)
		fmt.Printf("len: %d\tcap: %d\tpointer:%p\n", len(numList3), cap(numList3), numList3)
	}
	// 一次添加多个元素或是切片
	numList3 = append(numList3, 1, 2, 3)
	numList3 = append(numList3, numList2...) //要加上...

	// 复制切片
	// copy(destSlice, srcSlice []T) int
	// srcSlice为数据来源切片 destSlice为复制的对象
	// 来源和目标数据类型必须一致
	// destSlice必须分配过空间且足够承载复制的元素个数
	// copy返回值表示实际发生复制的元素个数
	srcData := make([]int, 100)
	for i := 1; i < 100; i++ {
		srcData[i] = i
	}
	refData := srcData // 引用
	copyData := make([]int, 100)
	copy(copyData, srcData) // 复制拷贝
	srcData[0] = 999
	fmt.Println(refData[0], copyData[0])
	copy(copyData, srcData[4:6]) // 拷贝从0开始
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
	fmt.Println()

	// 删除切片
	// 没有提供方法封装，只能自己模拟
	srcData = make([]int, 5)
	for i := 0; i < 5; i++ {
		srcData[i] = i
	}
	index := 2 //删除下标
	fmt.Println(srcData[:index], srcData[index+1:])
	srcData = append(srcData[:index], srcData[index+1:]...)
	fmt.Println(srcData)

	// map映射
	// 底层：散列表Hash

	fmt.Println("----------map----------")
	// 创建、映射、查询
	map1 := make(map[string]int)
	map1["hello"] = 1
	map1["hi"] = 2
	map1["ha"] = 3
	fmt.Println(map1["hi"], map1["no"]) // 2 0
	// 删除
	delete(map1, "ha")
	// 遍历
	// 遍历是无序的，若要有序，手动sort后模拟输出
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	// 清空
	map1 = make(map[string]int)

	fmt.Println("----------sync.Map----------")
	// 线程安全的map：sync.Map
	var map2 sync.Map
	// 存储映射Store，键和值以interface{}类型保存
	map2.Store("hello", 123)
	map2.Store(123, "hello")
	// 获取查询Load
	fmt.Println(map2.Load("hello")) // 123 true
	fmt.Println(map2.Load("hi"))    // <nil> false
	// 删除Delete
	map2.Delete("hello")
	// 遍历
	map2.Range(func(k, v interface{}) bool {
		fmt.Println("iterate: ", k, v)
		return true
	})

	fmt.Println("----------List----------")
	// 列表List
	// 列表是一种非连续存储的容器，底层是双链表

	// 声明初始化
	var list1 list.List
	list2 := list.New()
	_ = list2
	// 插入删除
	list1.PushBack(1)              // 1
	list1.PushFront(2)             // 2 1
	element := list1.PushBack("3") // 2 1 3
	list1.InsertAfter(4, element)  // 2 1 3 4
	list1.InsertBefore(5, element) // 2 1 5 3 4
	list1.Remove(element)          // 2 1 5 4
	// 遍历
	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
