package main

import (
	"context"
	"fmt"
	"time"
)

//超时关闭通知(5秒）

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case <-time.After(time.Second * 5):
			fmt.Println("您已超时")
			return
		}
	}()

	time.Sleep(time.Millisecond * 10000) //留点时间打印

}
