package main

import "fmt"

type Role string

func (r Role) Change1(role string) {
	r = Role(role)
}

func (r *Role) Change2(role string) {
	*r = Role(role)
}

func (r *Role) Change3(role string) {
	newRole := Role(role)
	r = &newRole
}

func main() {
	r := Role("name")
	r.Change1("Name")
	fmt.Printf("change1 %s\n", r)

	r.Change2("Name")
	fmt.Printf("change2 %s\n", r)

	r.Change3("Name")
	fmt.Printf("change3 %s\n", r)

	// 打印结果：
	// change1 name
	// change2 Name
	// change3 Name

	// 结论：
	// 想在结构体方法中覆盖结构体，可以使用 Change2 或 Change3 中方式。
}
