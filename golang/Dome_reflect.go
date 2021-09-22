package main

import (
	"fmt"
	"reflect"
)

// 反射机制 reflect

type MyPeople struct {
	FirstName, LastName string
}

func (this *MyPeople) GetName1() string {
	return fmt.Sprintf("[GetName1()] my name is %s%s", this.LastName, this.FirstName)
}

func (this MyPeople) GetName2() string {
	return fmt.Sprintf("[GetName2()] my name is %s%s", this.LastName, this.FirstName)
}

func ReflectMain() {
	people := MyPeople{FirstName: "San", LastName: "Zhang"}
	//people := MyPeople{ FirstName: "San",LastName: "Zhang"}
	v1Field := reflect.ValueOf(people)
	v1Method := reflect.ValueOf(&people)
	for i := 0; i < v1Field.NumField(); i++ {
		fmt.Printf("Field %d:%v\n", i, v1Field.Field(i))
	}

	fmt.Println(v1Method.NumMethod())
	for i := 0; i < v1Method.NumMethod(); i++ {
		fmt.Printf("Method %d:%v\n", i, v1Method.Method(i).String())
		fmt.Println(v1Method.Method(i).Call(nil))
	}
}
