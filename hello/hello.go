package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	go func() {
		fmt.Println("다른 고루틴!")
	}()
	go func() {
		fmt.Println("다른 고루틴!!")
	}()
	go func() {
		fmt.Println("다른 고루틴!!!")
	}()
	fmt.Println("STOP")
	<-ctx.Done()
	fmt.Println("재시작")
}
