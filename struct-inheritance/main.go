package main

import "fmt"

type person struct {
	Name string
}

func (p *person) PrintName() {
	fmt.Println(p.Name)
}

// BAD
type PersonEmbedNopointer struct {
	person
	Age int
}

func (p *PersonEmbedNopointer) PrintName() {
	fmt.Println("PersonEmbedNopointer PrintName", p.Name)
}

// GOOD
type PersonEmbedPointer struct {
	*person
	Age int
}

func (p *PersonEmbedPointer) PrintName() {
	fmt.Println("PersonEmbedPoint PrintName", p.Name)
}

func main() {
	p := person{"hello"}
	fmt.Printf("inner person field address:         %p\n", &p)

	p1 := PersonEmbedNopointer{p, 10}
	fmt.Printf("person embed nopointer field address: %p\n", &(p1.person))

	p2 := PersonEmbedPointer{&p, 10}
	fmt.Printf("person embed pointer field address:   %p\n", p2.person)

	// 调研覆盖后的方法
	p.PrintName()
	p1.PrintName()

	p.PrintName()
	p2.PrintName()

	// 打印结果：
	// inner person field address:           0xc000010250
	// person embed nopointer field address: 0xc00000c030
	// person embed pointer field address:   0xc000010250

	// hello
	// PersonEmbedNopoint PrintName hello
	// hello
	// PersonEmbedPoint PrintName hello

	// 结论：
	// go 中没有继承，但是有内嵌，达到和继承一样的效果。
	// 和 struct-function 中的代码一样，如果不内嵌指针的，会进行结构体复制，因此需要使用内嵌指针的方式。
	// https://golang.org/doc/effective_go.html#embedding
}
