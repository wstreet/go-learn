package main

import "fmt"
// go使用struct实现面向对象
// 定义
// type 类型名 struct {
// 	字段名 字段类型
// 	字段名 字段类型
// 	…
// }
// 1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
// 2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
// 3.字段类型：表示结构体字段的具体类型。



func main() {
	type Person struct {
		neme string
		age int8
	}

	type Person1 struct {
		neme, city string
		age int8
	}

	// 实例化结构体
	var p Person
	p.name = "wstreet7"
	p.age = 26


	// 匿名结构体
	var user struct {
		name string
		age int8
	}

	user.name = "wstreet7"
	user.age = 26

	// 指针型结构体
	// 通过使用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(Person)
	p2.name = "wstreet7"
	p2.age = 26

	// 使用&对结构体进行取地址操作，相当于对结构体进行了一次new()
	p3 := &Person{}

	// 使用健值对初始化
	p4 := Person{
		name: "wstreet7",
		age: 26
	}

	// 也可以对结构体指针进行健值对初始化
	p5 := &Person{
		name: "wstreet7",
		age: 26
	}
	// 使用值的列表初始化，省略健,必须初始化所有字段
	p6 := Person{
		"wstreet7",
		26
	}

	func NewPerson(name string, age int8) *Person {
    return &Person{
        name: name,
        age:  age,
    }
}

	func (p Person) dream() {
		fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
	}
	p7 := NewPerson("wstreet7, 26")
	p7.dream()

	// 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。


	//指针类型的接收者
	// 什么时候应该使用指针类型接收者
	// 1.需要修改接收者中的值
	// 2.接收者是拷贝代价比较大的大对象
	// 3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
	func (p *Person) setAge(age int8) {
		p.age = age
	}

	p7.setAge(25)

	//值类型的接收者
	//当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份


	// 结构体字段的可见性
	// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
}

