package main

import (
	"context"
	"fmt"
	"time"
)

// 워커 패턴: 중단 알림이 올 때까지 태스크의 수신 및 처리, 대기를 반복
// select 문과 Done을 사용한 폴링 구현.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	task := make(chan int)
	go func() {
		for {
			// case와 비슷한데 channel용임.
			select {
			case <-ctx.Done():
				return
			case i := <-task:
				fmt.Println("get", i)
			default:
				fmt.Println("중단되지 않음.")
			}
			time.Sleep(1 * time.Millisecond)
		}
	}()
	time.Sleep(time.Second * 1)
	for i := 0; 5 > i; i++ {
		task <- i
	}
	cancel()
}
