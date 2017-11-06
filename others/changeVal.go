package main

import "fmt"

func main() {
	x := 0

	changeVal(&x)
	fmt.Println("memory address of X :", &x)
	fmt.Println("X :", x)
}

func changeVal(x *int) {
	fmt.Println("memory address of X :", x)
	*x = 2
}
