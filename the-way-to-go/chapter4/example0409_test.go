package chapter4

import (
	"fmt"
	"testing"
)

func TestTypeConvert2(t *testing.T) {
	var n int64 = 34
	var m int32
	m = int32(n)
	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
}
