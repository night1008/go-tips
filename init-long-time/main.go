package main

import (
	"fmt"
	"init-long-time/model1"
	"init-long-time/model2"
)

func main() {
	u1 := model1.User{Name: "model1"}
	fmt.Println(u1)

	u2 := model2.User{Name: "model2"}
	fmt.Println(u2)

	// 打印结果：
	// ===> package model1 init start
	// ===> package model1 init end
	// ===> package model2 init start
	// ===> package model2 init end
	// {model1}
	// {model2}

	// 结论：
	// init 执行有顺序，切允许长时间执行
}
