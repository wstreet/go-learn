package src

import "fmt"


// 切片不是数组
// 切片是引用类型

func main() {
	// 声明切片
	var s1 []int

	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}


	s2 := []int{}

	var s3 = make([]int, 0)

	// 初始化
	var s4 = make([]int, 0, 0)

	s5 := []int{1, 2, 3}

	// 从数组切片
	arr := [5]int{1, 2, 3, 4, ,5}
	var s6 = []int

	// 	前包后不包
	s6 = arr[1:4]

	// append  ：向 slice 尾部添加数据，返回新的 slice 对象。
	s7 := append(s5, 6, 7, 8)

	s8 := []int{1, 2, 3, 4, 5}
   
    s9 := make([]int, 10)
   
    copy(s9, s8)
}