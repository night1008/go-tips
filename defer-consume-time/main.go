package main

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	m    sync.RWMutex
	Name string
}

// BAD
func assignNameWithDefer(i *Item) {
	i.m.Lock()
	defer i.m.Unlock()
	i.Name = "hello"
}

// GOOD
func assignNameWithoutDefer(i *Item) {
	i.m.Lock()
	i.Name = "hello"
	i.m.Unlock()
}

func main() {
	item := Item{}
	t1 := time.Now()
	assignNameWithDefer(&item)
	fmt.Println("call mute with defer:   ", time.Since(t1).Nanoseconds(), "ns")

	t2 := time.Now()
	assignNameWithoutDefer(&item)
	fmt.Println("call mute without defer:", time.Since(t2).Nanoseconds(), "ns")

	// 打印结果：
	// call mute with defer:    2392 ns
	// call mute without defer: 252 ns

	// 结论：
	// 原来使用 defer 对性能会有影响。
}
