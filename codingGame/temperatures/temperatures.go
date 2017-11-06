package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var n int
	min := MyNum{neg: false, abs: 5527, val: 0} // -273 - 5526

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		t, _ := strconv.ParseInt(inputs[i], 10, 64)
		current := MyNum{neg: (t < 0), val: int(t), abs: int(t)}
		if current.neg {
			current.abs *= -1
		}
		if min.abs > current.abs || (min.abs == current.abs && min.neg) {
			min = current
		}
	}

	fmt.Println(min.val)
}

// MyNum struct demo
type MyNum struct {
	neg bool
	abs int
	val int
}
