package cs

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string // 对外发送消息的单向通道

var (
	entering = make(chan client) // 通道传递的内容为另一个通道
	leaving  = make(chan client)
	messages = make(chan string) // 所有链接的客户端都是通过这个通道进行信息传递
)

func Broadcaster() {
	clients := make(map[client]bool) // 创建map容器，key为clien，value为bool

	for {
		select {
		case msg := <-messages:
			// 把所有接收到的消息广播给所有客户端
			// 发送消息通道
			for cli := range clients { // range clients意思是遍历clients的key
				cli <- msg
			}
		case cli := <-entering: // 新加入一个客户端
			clients[cli] = true
			log.Printf("当前聊天用户数: %d\n", len(clients))
		case cli := <-leaving: // 有一个客户端离开
			delete(clients, cli)
			close(cli)
			log.Printf("当前聊天用户数: %d\n", len(clients))
		}
	}
}

func HandleConn(conn net.Conn) {
	ch := make(chan string) // 对外发送客户消息的通道

	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "欢迎" + who
	messages <- who + "上线"
	log.Println(who + "上线")
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() { // 当前client可以一直输入，除非ctrl+c中断退出
		messages <- who + ": " + input.Text()
	}
	// 注意：忽略input.Err()中可能的错误

	log.Println(who + "下线")
	leaving <- ch
	messages <- who + "下线"
	conn.Close()
}

// 建立conn和通道的关系
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch { // 直到ch被关闭，这里才会结束
		fmt.Fprintln(conn, msg) // 注意：忽略网络层面的错误
	}
}
