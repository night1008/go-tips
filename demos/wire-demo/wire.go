//go:build wireinject
// +build wireinject

// 不加上面两行执行不了
// go run main.go 执行不了，得通过 go run .

package main

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
