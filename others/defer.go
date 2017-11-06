package main

import "fmt"

func main() {
	defer print2()
	print1()

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}

func print1() {
	fmt.Println("1")
}

func print2() {
	fmt.Println("2")
}
