package main

import "fmt"
// go中map是引用类型，必须初始化才能使用
// 定义
// map[KeyType]ValueType
// KeyType:表示键的类型。

// ValueType:表示键对应的值的类型。 默认nil

// 需要使用make函数来分配内存
// make(map[KeyType]ValueType, [cap])


func main() {
	sourceMap := make(map[string]int, 8)
	sourceMap["张三"] = 90
	sourceMap["小明"] = 80
	fmt.Println(sourceMap)
	fmt.Println(sourceMap["小明"])
	fmt.Printf("type of sourceMap: %T\n", sourceMap)

	// map也支持在声明的时候填充元素
	sou := make(map[string]string{
		"usename": "wstreet7",
		"email": "wstreet7@outlook.com"
	})


	// 判断某个键是否存在
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := sourceMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("空")
	}

	// map遍历
	// 使用for range遍历
	for k, v := range sourceMap {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range sourceMap {
		fmt.Println(k)
	}

	// delete()删除健值对
	delete(sourceMap, "张三")

	// 指定顺序遍历map
	sourceMap1 := make(map[string]int, 200)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intin(100) // 0-99随机整数
		sourceMap1[key] = value
	}

	// 取出map中的key放入切片
	keys := make([]string, 0, 200)
	for k := range sourceMap1 {
		keys = append(keys, k)
	}

	// 对切片进行排序
	sort.String(keys)

	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, sourceMap1[key])
	}
}