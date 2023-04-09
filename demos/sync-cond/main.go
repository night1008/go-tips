package main

import (
	"log"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)
	time.Sleep(2 * time.Second)
	go read("reader4", cond)

	time.Sleep(time.Second * 3)
}

// 好像可以用来做查询去重，比如多个用户发起了一样的查询，
// 用这个方式只需要查一次，再把结果广播出去

type Query struct {
	sql  string
	done chan struct{}
}
