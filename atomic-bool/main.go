package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var running atomic.Bool
var wg sync.WaitGroup

func run(d time.Duration) {
	defer wg.Done()

	if wasRunning := running.Swap(true); wasRunning {
		return
	}
	defer running.Swap(false)

	time.Sleep(d)
	fmt.Println("run duration ", d)
}

func main() {
	wg.Add(1)
	go run(3 * time.Second)
	time.Sleep(10 * time.Millisecond)

	wg.Add(1)
	go run(1 * time.Second)

	wg.Wait()
}

// 结论：
// 某些需要保证执行一次的情况下，可以使用 atomic.Bool
