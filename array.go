package src

import "fmt"

// 数组定义,数组长度是类型的组成部分
var a [5]int = [5]int{1, 2, 3, 4, 5}

for i:= 0; i < len(a); i ++ {

}

// 数组是值类型，赋值和传参数会复制真个数组，而不是指针，改变副本的值，不会改变原始的值
// 指针数组 [n]*T, 数组指针 *[n]T

// 数组初始化
// 全局
var b = [2]string = [2]string{"q", "w"}
// 函数内部
c := [3]int{1, 2} // 未初始化的值为0
d := [...]int{1, 2, 3} // 通过初始化值确定数组长度。
e := [4]int{2: 100 } // 使用索引号初始化值

type Person struct {
	name string
	age int
}

f := [2]Person{
	{ "user1", 26 },
	{ "user2", 27 },
}

// 多维数组
// 全局
var arr0 = [2][2]int
var arr1 = [2][3]int{{1, 2, 3}, { 1, 2, 3}}
// 局部
arr3 := [...][3]int{{1, 2, 3}, { 1, 2, 3}} //  第 2 纬度不能用 "..."。


func test(x [2]int)  {
	x[1] = 100
	fmt.Printlf("x: %p", &x)
}

func main() {
	a := [2]int{1, 2}
	fmt.Printlf("a%p", &a)

	test(a)
}

// 多维数组遍历
var arr1 = [2][3]int{{1, 2, 3}, { 1, 2, 3}}

for k1, v1 := range arr1 {
	for k2, v2 := range v1 {
		fmt.Printlf("%d", v2)
	}
}