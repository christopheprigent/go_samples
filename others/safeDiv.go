package main

import "fmt"

func main() {
	fmt.Printf("safeDiv %d\n", safeDiv(1, 0))
}

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Printf("recover : %s", recover())
	}()
	s := num1 / num2

	return s
}
