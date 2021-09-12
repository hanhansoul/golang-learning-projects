package chapter4

import (
	"fmt"
	"testing"
)

func TestTypeConvert1(t *testing.T) {
	var a int
	var b int32
	a = 15
	//b = a + a
	b = b + 5
	fmt.Println(a, b)
}
