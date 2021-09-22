package main

import "fmt"

// 接口演示
type IF interface {
	getName() string
}

type ZhangSan struct {
	firstName, lastName string
}

type LiSi struct {
	firstName, lastName string
}

func (zs *ZhangSan) getName() string {
	zs.lastName = "张"
	zs.firstName = "三"
	return fmt.Sprintf("My name is %s%s", zs.lastName, zs.firstName)
}

func (ls *LiSi) getName() string {
	ls.lastName = "李"
	ls.firstName = "四"
	return fmt.Sprintf("My name is %s%s", ls.lastName, ls.firstName)
}

func IFMain() {
	People := []IF{}

	ZS := new(ZhangSan)
	LS := new(LiSi)
	People = append(People, ZS)
	People = append(People, LS)

	for _, v := range People {
		fmt.Println(v.getName())
	}
}
