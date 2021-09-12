package chapter4

import (
	"fmt"
	"strings"
	"testing"
)

func TestExample41(t *testing.T) {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}