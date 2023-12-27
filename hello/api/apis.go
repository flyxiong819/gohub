package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo"
)

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hi, hello world")
}

func TestChannel() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		for i := 3; i >= 0; i-- {
			ch <- i

			time.Sleep(time.Second)
		}
		// 通过通道通知goroutine
		ch <- 99
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")
	// 等待匿名的goroutine
	for out := range ch {
		fmt.Println("receive done", out)

		if out == 99 {
			break
		}
	}
}

var wg sync.WaitGroup

func PlayBall() {
	// 创建一个吴缓冲通道
	court := make(chan int)

	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 启动两个选手
	go player("First", court)
	go player("Second", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

// 模拟一个选手在打网球
func player(name string, court chan int) {
	// 在程序退出时调用Done，通知已完成
	defer wg.Done()

	for {
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 关闭通道，表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数+1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 将球打向对手
		court <- ball
		time.Sleep(time.Second)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
