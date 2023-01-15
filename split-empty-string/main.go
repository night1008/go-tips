package main

import (
	"fmt"
	"strings"
)

func main() {
	emptyStr := ""
	parts := strings.Split(emptyStr, ",")
	fmt.Printf("empty string split parts length is: %d\n", len(parts))

	// 打印结果：
	// empty string split parts length is: 1

	// 结论：
	// 需要注意 go 语言拆分空字符串竟然得到一个包含空字符串的数组
}
