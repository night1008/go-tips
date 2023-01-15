package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := "hello"
	s2 := "world"

	t1 := time.Now()
	s3 := s1 + s2
	fmt.Println(s3, time.Since(t1).Nanoseconds())

	t2 := time.Now()
	s4 := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println(s4, time.Since(t2).Nanoseconds())

	// 打印结果：
	// helloworld 365
	// helloworld 2267

	// 结论：
	// 直接使用 + 拼接字符串比使用 fmt.Sprintf 消耗更少的时间
}
