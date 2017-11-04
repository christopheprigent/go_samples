package main

import "fmt"

var i, j = 1, "2"

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	t()
}

func t() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
