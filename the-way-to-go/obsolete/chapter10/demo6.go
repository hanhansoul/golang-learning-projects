package chapter10

import (
	"fmt"
	"math"
)

type TwoInts struct {
	a int
	b int
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func StructTest7() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

type B struct {
	thing int
}

func (b *B) change() int {
	b.thing = 1
	return b.thing
}

func (b B) write() string {
	return fmt.Sprint(b)
}

func (b B) changeVar() {
	b.thing = 20
}

func (b *B) changePoint() {
	b.thing = 20
}

func StructTest8() {
	//var b1 B // b1是值
	//b1.change()
	//fmt.Println(b1.write())
	//
	//b2 := new(B) // b2是指针
	//b2.change()
	//fmt.Println(b2.write())

	b := new(B)
	b.thing = 10
	b.changeVar()
	fmt.Println(b)
	b.changePoint()
	fmt.Println(b)
}

type Point3 struct{ x, y, z float64 }

// A method on Point3
func (p Point3) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func StructTest9() {
	p := &Point3{3, 4, 5}
	fmt.Println(p.Abs())
	fmt.Println((*p).Abs())
}

func StructTest10() {
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.change())

	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.change())
}
