package chapter4

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("hello")
}

func TestA(t *testing.T) {
	main()
}