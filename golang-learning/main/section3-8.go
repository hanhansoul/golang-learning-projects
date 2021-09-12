package main

import "fmt"

func main() {
	// 先定义一个数组
	var myArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建一个数组切片
	var mySlice = myArray[:5]
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	mySlice1 := make([]int, 5)
	mySlice2 := make([]int, 5, 10)
	mySlice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice1)
	fmt.Println(mySlice2)
	fmt.Println(mySlice3)

	mySlice4 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice4...)

}
