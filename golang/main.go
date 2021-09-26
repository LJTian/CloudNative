package main

import (
	"fmt"
)

// type 细节区别
type myString string
type AliasString = string

func TypeMain() {

	fmt.Println("Type 细节实现 测试【重点在 “=” 号】")
	var str1 string
	var str2 myString
	var str3 AliasString
	str1 = "String"
	str2 = "String"

	fmt.Printf(" str1 is [%T] \n str2 is [%T]\n str3 is [%T]\n", str1, str2, str3)

}

// 练习题1 替换数据
func Exercise() {

	fmt.Println("练习题1 替换数据")
	a := []string{"I", "am", "stupid", "and", "weak"}
	for k, v := range a {
		if v == "stupid" {
			a[k] = "smart"
		}

		if v == "weak" {
			a[k] = "strong"
		}
	}
	fmt.Println(a)
}

func main() {

	//TypeMain()
	//Exercise()
	//IFMain()
	//ReflectMain()
	//MarshalMain()
	//PanicRecoverMain()
	//GoroutineChannelMain()
	SyncGoSliceMain(1000, 50, 20)
}
