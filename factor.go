package main

import "fmt"

func main() {
	i := 4

	fmt.Printf("facto %d : %d", i, f(i))
}

func f(i int) int {
	if i < 2 {
		return i
	}

	return i * f(i-1)
}
