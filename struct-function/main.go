package main

import "fmt"

type Item struct {
	Name string
}

// BAD
func (i Item) PrintAddrOfNoPointer() {
	fmt.Printf("struct address in no pointer strcut function: %p\n", &i)
}

// GOOD
func (i *Item) PrintAddrOfPointer() {
	fmt.Printf("struct address in pointer strcut function:    %p\n", i)
}

func printAddrOfNoPointerStruct(i Item) {
	fmt.Printf("struct address in no pointer function:        %p\n", &i)
}

func printAddrOfPointerStruct(i *Item) {
	fmt.Printf("struct address in pointer function:           %p\n", i)
}

func main() {
	item := Item{}
	fmt.Printf("struct address in main:                       %p\n", &item)
	item.PrintAddrOfPointer()
	item.PrintAddrOfNoPointer()
	printAddrOfNoPointerStruct(item)
	printAddrOfPointerStruct(&item)

	// 打印结果：
	// struct address in main:                       0xc000010250
	// struct address in pointer strcut function:    0xc000010250
	// struct address in no pointer strcut function: 0xc000010260
	// struct address in no pointer function:        0xc000010270
	// struct address in pointer function:           0xc000010250

	// 结论：
	// 避免不必要的结构体复制，应该尽量使用 struct point
}
