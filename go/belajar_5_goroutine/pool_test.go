package belajar5goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Lutfan")
	pool.Put("Zainul")
	pool.Put("Haq")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
			time.Sleep(1 * time.Second)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}