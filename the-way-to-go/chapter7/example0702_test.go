package chapter7

import (
	"fmt"
	"testing"
)

//func f(a [3]int) {
//	fmt.Println(a)
//}
//
//func fp(a *[3]int) {
//	fmt.Println(a)
//}

func TestExmaple0702(t *testing.T) {
	//var ar [3]int
	//f(ar) 	// passes a copy of ar
	//fp(&ar) // passes a pointer to ar

	var arr1 = new([5]int)
	var arr2 [5]int
	arr1 = &arr2
	//arr2 = *arr1
	//arr2[2] = 100
	arr1[2] = 100
	fmt.Println(arr1)
	fmt.Println(arr2)
}
