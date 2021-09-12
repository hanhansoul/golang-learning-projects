package chapter3

import (
	"fmt"
	"testing"
)

var global = 100

func outer(a int) func() (int, int) {
	//a := 10
	closure := func() (int, int) {
		a++
		global++
		return a, global
	}
	return closure
}
func TestClosure1(t *testing.T) {
	f1 := outer(5)
	f2 := outer(10)
	println(f1())
	println(f1())
	println(f2())
	println(f2())
	println(global)
}

func TestClosure2(t *testing.T) {
	var l []func() int
	for val := 1; val <= 3; val++ {
		ff := func() int {
			return val * val
		}
		l = append(l, ff)
	}
	for _, g := range l {
		fmt.Println(g())
	}
}

func TestClosure3(t *testing.T) {
	var l []func() int
	var f = func(val int) func() int {
		return func() int {
			return val * val
		}
	}
	for val := 1; val <= 3; val++ {
		l = append(l, f(val))
	}
	for _, g := range l {
		fmt.Println(g())
	}
}
