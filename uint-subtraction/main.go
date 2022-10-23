package main

import "fmt"

func main() {
	var a uint = 10
	var b uint = 20
	fmt.Printf("%d - %d = %d\n", a, b, a-b)

	// 打印结果：
	// 10 - 20 = 18446744073709551606

	// 结论：
	// 首先 uint 不能赋值负数
	// 其次，应该小心 uint 数值相减后可能得到一个很大的数字
	// 尤其在 increase 和 decrease 的场景下需要注意，不能简单地通过乘以 -1
}
