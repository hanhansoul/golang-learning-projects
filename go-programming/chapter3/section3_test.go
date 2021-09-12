package chapter3

import (
	"fmt"
	"testing"
)

/*
1. 一般调用
2. 方法值
3. 方法表达式：提供一种语法将类型方法调用显式地转换为函数调用，接收者必须显式地传递进去。

方法集
命名类型方法接收者有两种类型：
1. 值类型
2. 指针类型

将接收者为值类型T的方法集合记录为S，将接收者为指针类型*T的方法集合记录为*S：
1. T类型的方法集为S
2. *T类型的方法集为S和*S

在使用类型实例调用类型的方法时，无论值类型变量还是指针类型变量，都可以调用类型的所有方法，编译器在编译期间能够识别这种调用关系，并做自动转换。

组合结构的方法集规则：
1. 若类型S包含匿名字段T，则S的方法集包含T的方法集；
2. 若类型S包含匿名字段*T，则S的方法集包含T和*T的方法集；
3. 不管类型S中嵌入的匿名字段是T还是*T，*S的方法集总是包含T和*T的方法集。

*/

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(i int) {
	t.a = i
}

func TestExample31(tt *testing.T) {
	var t = &T{}
	t.Set(2)
	t.Get()
}

func TestExample32(tt *testing.T) {
	var t = &T{}
	f := t.Set
	f(2)
	fmt.Printf("%p, %v, %d, \n", t, t, t.a)
	f(3)
	fmt.Printf("%p, %v, %d, \n", t, t, t.a)
}

func TestExample33(tt *testing.T) {
	t := T{a: 1}
	t.Get()
	T.Get(t)

	f1 := T.Get
	f1(t)

	f2 := T.Get
	f2(t)

	(*T).Set(&t, 1)
	f3 := (*T).Set
	f3(&t, 1)
}

type Int int

func (a Int) Max(b Int) Int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func (i *Int) Set(a Int) {
	*i = a
}

func (i Int) Print() {
	fmt.Printf("value=%d\n", i)
}

func TestExample34(t *testing.T) {
	var a Int = 10
	var b Int = 20

	c := a.Max(b)
	c.Print()
	(&c).Print()

	a.Set(20)
	a.Print()

	(&a).Set(30)
	a.Print()
}

type Data struct{}

func (Data) TestValue() {}

func (*Data) TestPointer() {}

func TestExample35(t *testing.T) {
	d := Data{}
	p := &Data{}

	// 通过实例调用方法
	d.TestPointer()
	d.TestValue()
	p.TestPointer()
	p.TestValue()

	(*Data)(&struct{}{}).TestPointer()
	(*Data)(&struct{}{}).TestValue()
	(Data)(struct{}{}).TestValue()
	//(Data)(struct{}{}).TestPointer()

	//(*Data).TestValue(d)
	//(*Data).TestPointer(d)
	(*Data).TestValue(p)
	(*Data).TestPointer(p)

	//Data.TestPointer(d) // runtime error
	//Data.TestPointer(p)
	Data.TestValue(d)
	//Data.TestValue(p)
}

type X struct {
	a int
}
type Y struct {
	X
}
type Z struct {
	*X
}

func (x X) Get() int {
	return x.a
}
func (x *X) Set(v int) {
	x.a = v
}

func TestExample36(t *testing.T) {
	x := X{a: 1}
	y := Y{X: x}
	println(y.Get())
	y.Set(2)
	println(y.Get())

	(*Y).Set(&y, 3)
	//Y.Set(y, 3) // runtime error
}
