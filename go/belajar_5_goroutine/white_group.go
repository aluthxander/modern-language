package belajar5goroutine

import (
	"fmt"
	"sync"
	"time"
	"testing"
)

func RunAsynchronous(group *sync.WaitGroup)  {
	defer group.Done()
	group.Add(1)
	
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T)  {
	group := &sync.WaitGroup{}
	group.Add(2)
	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}
	group.Wait()
	fmt.Println("Selesai")
}