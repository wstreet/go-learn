package src

import "fmt"
import "unsafe"

func main()  {

	// 常量定义
	const WIDTH int = 30
	const HEIGHT = 50

	const a, b, c = 2, 3, "4"

	// 常量还可以作为枚举
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)
// 内置函数才行
	const (
		a = "qwe"
		b = len(a)
		c = unsafe.Sizeof(a)
	)

// 省略表示相同
	const (
		n1 = 100
		n2
		n3
	)

	// 变量声明
	var age = 26
	var name string = "street7"
	age = 27
	name = "wstreet7"

	// 函数内部还可以使用短变量声明
	short := 3

	println(WIDTH, HEIGHT, a,b,c,age, name)



	// 匿名变量_
	// 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	_, y := foo()
	x, _ := foo()
		

	// 函数外的每个语句都必须以关键字开始（var、const、func等）
	// :=不能使用在函数外。
	// _多用于占位，表示忽略值。


	// iota是go语言的常量计数器
	const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
)
	
}

func foo() (int string) {
	return (3, "sss")
}