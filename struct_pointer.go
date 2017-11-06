/**
 * Les pointeurs vers des structs
 * Les champs struct peuvent être accessibles via un pointeur de struct.
 * L'indirection via le pointeur est transparente.
 **/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 42
	v.X = 43
	fmt.Println(v)
}
