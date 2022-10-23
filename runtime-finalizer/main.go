package main

import (
	"fmt"
	"runtime"
	"time"
)

type Item struct {
	Name string
}

func (i *Item) Run() {
	runtime.SetFinalizer(i, clear)
	fmt.Println("Print name from run:  ", i.Name)
}

func clear(i *Item) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Print name from clear:", i.Name)
}

func main() {
	item := Item{"hello"}
	item.Run()
	runtime.GC()
	time.Sleep(300 * time.Millisecond)

	// 打印结果：
	// Print name from run:   hello
	// Print name from clear: hello

	// 结论：
	// 该函数可以在对象回收的时候触发回调函数，可以进行资源回收等操作，
	// 但是从文档上看，该函数的使用有不少的限制，
	// 对于和主程序生命周期基本一致的后台 goroutine，一般采用显式的`Stop()`来进行优雅退出，
	// `Stop()` 调用需要对使用者透明时可以使用
	// https://pkg.go.dev/runtime#SetFinalizer
	// https://github.com/patrickmn/go-cache/blob/master/cache.go#L1123
}
