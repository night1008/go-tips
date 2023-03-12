package main

import "fmt"

// Declare a type constraint
type Number interface {
	int64 | float64
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func mapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	r := mapKeys(m) // 自己回推断类型
	fmt.Println(r)
}

// 结论：
// 使用类型参数泛型可以共用很多实现
