package src
import(
	"fmt"
	"time"
)

// 在调用函数或者 方法时在前面加一个go关键字，可以让一个新的协程并发的运行
 

func hello(done chan int)  {
	fmt.Println("Hello world goroutine")
	done <- true
}

func main() {
	done := make(chan int)
	go hello(done)
	<- done
	
	fmt.Println("main function")
}
