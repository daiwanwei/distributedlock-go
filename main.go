package main

import (
	"context"
	"distributedlock-go/pkg/distributedlock"
	"fmt"
	"sync"
	"time"
)

type Hello string

func (h Hello) Key() string {
	return string(h)
}

func main() {
	num := 10
	locker, err := distributedlock.NewLockManager()
	if err != nil {
		return
	}
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		n := i
		go func() {
			defer wg.Done()
			fmt.Println("hello to lock:", n)
			h := Hello("fucker")
			lock, err := locker.Lock(context.Background(), h)
			if err != nil {
				fmt.Println(err, n)
				return
			}
			defer func() {
				err = lock.Unlock(context.Background())
				if err != nil {
					fmt.Println("Unlock err:", err)
				}
			}()
			fmt.Println("hello locking:", n)
			time.Sleep(time.Millisecond * 1000)
			fmt.Println("hello to unlock:", n)
			err = lock.Unlock(context.Background())
			if err != nil {
				fmt.Println(err, n)
				return
			}
			fmt.Println("hello end unlock:", n)
		}()
	}
	wg.Wait()
}
