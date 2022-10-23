package main

import "fmt"

const (
	TypeA = "A"
	TypeB = "B"
	TypeC = "C"
)

type I interface{}

type A struct{}

type B struct{}

type C struct{}

func NewA() I {
	fmt.Println("NewA")
	return &A{}
}

func NewB() I {
	fmt.Println("NewB")
	return &B{}
}

func NewC() I {
	fmt.Println("NewC")
	return &C{}
}

// BAD
func NewWithSwitch(_type string) (I, error) {
	switch _type {
	case TypeA:
		return NewA(), nil
	case TypeB:
		return NewB(), nil
	case TypeC:
		return NewC(), nil
	default:
		return nil, fmt.Errorf("unknown type %s", _type)
	}
}

var funcs map[string]func() I

func init() {
	funcs = make(map[string]func() I)
	funcs[TypeA] = NewA
	funcs[TypeB] = NewB
	funcs[TypeC] = NewC
}

// GOOD
func NewWithFuncMap(_type string) (I, error) {
	f, ok := funcs[_type]
	if !ok {
		return nil, fmt.Errorf("unknown type %s", _type)
	}
	return f(), nil
}

func main() {
	fmt.Println(NewWithSwitch(TypeA))
	fmt.Println(NewWithFuncMap(TypeA))

	// 打印结果：
	// NewA
	// &{} <nil>
	// NewA
	// &{} <nil>

	// 结论：
	// 可以通过 func map 减少 switch type 判断
}
