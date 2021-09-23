package main

import "fmt"

// Panic 和 recover
func PanicFunc() {
	panic("core err!")
}

func PanicRecoverMain() {
	defer func() {
		fmt.Println("程序结束~")
		if err := recover(); err != nil {
			fmt.Println("发生panic内容:", err)
		}
	}()
	PanicFunc()
}
