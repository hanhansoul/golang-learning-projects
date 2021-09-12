package chapter10

import "fmt"

type number struct {
	f float32
}

type nr number

func StructTest3() {
	a := number{5.0}
	b := nr{5.0}

	var c = number(b)
	fmt.Println(a, b, c)
}

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func StructTest4() {
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"
}

