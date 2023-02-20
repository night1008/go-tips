package model1

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("===> package model1 init start")
	time.Sleep(3 * time.Second)
	fmt.Println("===> package model1 init end")
}

type User struct {
	Name string
}
