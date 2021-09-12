package chapter7

import "fmt"

func Test1() {
	var arr1 = new([5]int)
	arr2 := *arr1
	arr2[2] = 100
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func f(a [3]int) {
	a[0] = 100
	//fmt.Println(a)
}

func fp(a *[3]int) {
	a[0] = 100
	//fmt.Println(a)
}

func Test2() {
	var ar [3]int
	f(ar)
	fmt.Println(ar)
	fp(&ar)
	fmt.Println(ar)
}

func Test3() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(arrAge)
	fmt.Println(arrLazy)
	fmt.Println(arrKeyValue)
}
