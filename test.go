package main

import (
	"fmt"
)

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		// reg := []string{"fib closure", strconv.Itoa(a), strconv.Itoa(b)}

		fmt.Printf("FIB :%d, %d\n", a, b)

		a, b = b, a+b
		return a
	}
}

func fibRec(a int) int {

	if a < 2 {
		return a
	}
	fmt.Printf("fibrec %d\n", a)
	return fibRec(a-2) + fibRec(a-1)
}

func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-2) + fibo(n-1)
}

func main() {
	// f := fib()
	// // Function calls are evaluated left-to-right.
	// fmt.Println(f(), f(), f(), f(), f())
	fmt.Printf("fibREc RET : %d\n", fibRec(30))
	fmt.Printf("fibo RET : %d\n", fibo(30))
}
