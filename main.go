package main

import (
	"fmt"
	"test-in-go/tests"
)

func main() {
	add := tests.IsOdd(10)
	fmt.Println(add)
}
