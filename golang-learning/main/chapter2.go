package main

import "fmt"

var v1 int
var v2 string
var v3 [10]int
var v4 []int

const ( // iota被重设为0
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

func main() {
	var value2 int32
	value1 := 64
	value2 = int32(value1)
	fmt.Println(value2)

	i, j := 1, 2
	if i == j {
		fmt.Println("i and j are equal")
	}

	str := "Hello, 中文"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	for i, ch := range str {
		fmt.Println(i, ch)
	}

}
