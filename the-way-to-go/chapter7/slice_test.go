package chapter7

import (
	"fmt"
	"testing"
)

func TestSlice1(t *testing.T) {
	// array
	//var arr1 [5]int
	//var arr2 = [5]int{18, 20, 15, 22, 16}
	//var arr3 = [...]int{5, 6, 7, 8, 22}
	// slice
	//var slice1 = []int{5, 6, 7, 8, 22}
	//var slice2 = []int{3: 30, 4: 40}
	var slice3 []int
	slice4 := []int{}
	slice5 := make([]int, 0)
	fmt.Println(slice3 == nil)
	fmt.Println(slice4 == nil)
	fmt.Println(slice5 == nil)
}

func TestSlice2(t *testing.T) {
	// generated from array
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	var slice1 = arr[2:5]
	slice2 := arr[3:6]
	fmt.Println(slice1)
	fmt.Println(slice2)

	// nil slice / empty slice
	var slice3 []int
	var slice4 = []int{}
	var slice5 []int
	slice6 := []int{}
	slice7 := make([]int, 0)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
	fmt.Println(slice7)

	//
	var slice8 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice8)
}

func TestSlice3(t *testing.T) {
	//	通过一个已经创建好的数组定义切片
	var arr = [...]int{1, 2, 3, 4, 5}
	var slice1 = arr[1:3]
	fmt.Println(slice1)

	//	通过make创建切片
	//  var slice []type = make([]type, len, [cap])
	var slice2 = make([]int, 5)
	slice2[1] = 5
	slice2[4] = 10
	fmt.Println(slice2)

	//	定义一个切片，并指定具体的数组，其原理类似于make
	var slice3 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice3)

	var slice4 = slice3[2:4]
	fmt.Println(slice4)
}
