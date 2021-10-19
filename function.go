package src

func main() {
	// 定义函数类型
	type FormatFunc func (s string, x, y int) (string int) 

	func format(fn FormatFunc, s string, x, y int) (string int)  {
		return fn(s, x, y)
	}

	// /值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	// 引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
	func swap(x, y *int)  { // 交换地址
		var temp int
		temp = *x // 保存x的值
		*x = *y
		*y = temp
	}

	// 无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
	// map、slice、chan、指针、interface默认以引用的方式传递。

	// 没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。

	func calc(a, b int) (sum int, avg int) {
    sum = a + b
    avg = (a + b) / 2

    return
	}

	sum, avg := calc(a, b)
}