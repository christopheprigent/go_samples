package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	horses := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&horses[i])
	}
	sort.Ints(horses)
	prev := -1
	minDiff := horses[1] - horses[0]
	for _, h := range horses {
		current := h - prev
		if prev != -1 && minDiff > current {
			minDiff = h - prev
		}

		prev = h
	}
	fmt.Println(minDiff)
}
