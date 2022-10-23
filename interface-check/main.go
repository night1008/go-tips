package main

type Reader interface {
	Count() int64
}

type Read struct{}

func (r *Read) Count() int64 {
	return 0
}

// GOOD
var _ Reader = (*Read)(nil)

func main() {
	// 结论：
	// 可以通过 `var _ Reader = (*Read)(nil)` 的方式，检查某个结构体是否实现了某个接口
}
