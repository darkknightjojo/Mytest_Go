package multi_process

import (
	"context"
	"fmt"
	"time"
)

func TestWithCancel(t int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("TestWithCancel.Done", ctx.Err())
	case e := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("TestWithCancel:", e)
	}
	return
}

func TestWithDeadline(t int) {
	ctx := context.Background()
	dl := time.Now().Add(time.Duration(t) * time.Second)
	ctx, cancel := context.WithDeadline(ctx, dl)
	defer cancel()
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("TestWithDeadline.Done", ctx.Err())
	case e := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("TestWithDeadline:", e)
	}
	return
}
