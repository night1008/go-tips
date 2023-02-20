package model2

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("===> package model2 init start")
	time.Sleep(3 * time.Second)
	fmt.Println("===> package model2 init end")
}

type User struct {
	Name string
}
