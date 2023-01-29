package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := make(map[User]string)
	users[User{"a", 1}] = "a"
	users[User{"b", 2}] = "b"
	users[User{"c", 1}] = "c"

	userName, ok := users[User{"a", 1}]
	if ok {
		fmt.Println("find user name", userName)
	} else {
		fmt.Println("cannot find user name", userName)
	}

	// 打印结果：
	// find user name a

	// 结论：
	// 可以指定 struct 作为 map 的 key
}
