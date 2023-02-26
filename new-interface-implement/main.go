package main

import (
	"context"
	"fmt"
)

type Job interface {
	Run(context.Context)
}

type ShellJob struct {
	Cmd string
}

func (s *ShellJob) Run(ctx context.Context) {
	fmt.Println("Run", s.Cmd)
}

// GOOD
func NewShellJob1(cmd string) *ShellJob {
	return &ShellJob{
		Cmd: cmd,
	}
}

// BAD
func NewShellJob2(cmd string) Job {
	return &ShellJob{
		Cmd: cmd,
	}
}

func main() {
	j1 := NewShellJob1("ls -l")
	fmt.Println(j1)

	j2 := NewShellJob2("ls -l")
	fmt.Println(j2)
}

// 结论：
// 创建某个接口的实现结构体时，直接返回该结构体指针，而不是接口
