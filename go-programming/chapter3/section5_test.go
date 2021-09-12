package chapter3

import "testing"

/**
函数类型
1. 函数字面量类型：`func (InputTypeList) OutputTypeList`。“有名函数”和“匿名函数”都是函数字面量类型。
2. 函数命名类型：`type NewFuncType FuncLiteral`。

函数类型的意义：
1. 函数也是一种类型，函数字面量类型上可以定义各种函数命名类型
2. 有名函数和匿名函数的函数签名与命名函数类型的底层类型相同，可以进行类型转换
3. 可以为有名函数类型添加方法，提供了一种装饰设计模式
4. 为有名函数类型添加方法，使其与接口建立联系，使用接口的地方可以传递函数类型变量
*/

// 有名函数定义，函数名为add
// add的类型是函数字面量类型func (int, int) int
func add(a, b int) int {
	return a + b
}

// 函数声明语句
//func add(int, int) int

//func (int, int) int

type ADD func(int, int) int

func TestExample51(t *testing.T) {
	var add func(int, int) int
	add = func(a, b int) int {
		return a + b
	}
	add(1, 2)
}

func TestExample52(t *testing.T) {
	var g ADD = add
	res := g(1, 2)
	println(res)
}
