package main

import "fmt"

/*
 *类型指针:
 *允许对这个指针类型的数据进行修改。传递数据使用指针，而无须拷贝数据
 *指针不能进行偏移和运算
 *
 *切片：
 *由指向起始元素的原始指针、元素数量和容量组成
 */

func swap(a, b *int) {
	*b, *a = *a, *b
}
func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

	// 创建方式1
	ptr1 := &x
	_ = ptr1
	// 创建方式2
	ptr2 := new(int)
	ptr2 = &x
	fmt.Printf("PtrValue:%p\nValue:%d\nType:%T\n", ptr2, *ptr2, ptr2)
}
