package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	go func(c context.Context) {
		// time.Sleep(time.Second * 2)
		fmt.Println(ctx.Deadline())
		for {
			select {
			case <-c.Done():
				fmt.Println("done")
				time.Sleep(time.Second * 3)
				break
			case <-time.After(time.Second):
				fmt.Println("过去一秒了")
			}
		}
	}(ctx)
	// ctx.Done()
	fmt.Println("yi")
	time.Sleep(time.Second * 3)
	// fmt.Println(ctx.Deadline())
	fmt.Println("er")
	ctx.Done()
	fmt.Println("san")

	time.Sleep(time.Second * 5)
	cancel()
	fmt.Println("si")
	time.Sleep(time.Second * 2)
}
