package flow

// 可省略break
// case必须是相同类型
// 可以使用fallthrough强制执行后面的case代码

func main() {
	var s = "www"
	switch s {
		case "ww":
			//
		case "wee":
			//
		default:
		//
	}

	// type switch
	var x interface {}

	switch i := x.(type) {
	case nil:
		// 
	case int:
		// 
	default:

		
	}


}