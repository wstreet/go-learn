package flow

import "fmt"

// 可省略条件表达式括号。
// 持初始化语句，可定义代码块局部变量。
// • 代码块左 括号必须在条件表达式尾部。

// if 布尔表达式 {
// /* 在布尔表达式为 true 时执行 */
// }

func main() {
	n := 3
	if x := "abc"; n > 0 {
		fmt.Println(x)
	}
}