package flow

// select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
// 每个case语句里必须是一个IO操作

// c := make(chan bool) //创建一个无缓冲的bool型Channel 
// c <- x        //向一个Channel发送一个值
// <- c          //从一个Channel中接收一个值
// x = <- c      //从Channel c接收一个值并将其存储到x中
// x, ok = <- c  //从Channel接收一个值，如果channel关闭了或没有数据，那么ok将被置为false

func main() {
	
}