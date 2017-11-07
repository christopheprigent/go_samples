package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	msg := scanner.Text()
	bin := ""
	for i := 0; i < len(msg); i++ {
		var ord = rune(string(msg[i])[0])
		bin += fmt.Sprintf("%b", ord)
	}

	fmt.Println(Bin2Chuck(bin)) // Write answer to stdout
}

// Bin2Chuck helper for chuck
func Bin2Chuck(bin string) (chuck string) {
	// fmt.Println("Bin2Chuck:", bin)
	chuck = ""
	for i := 0; i < len(bin); i++ {
		c := string(bin[i])
		if c == "0" {
			chuck += "00 0"
		} else {
			chuck += "0 0"
		}
		for j := i + 1; j < len(bin) && string(bin[j]) == c; j++ {
			chuck += "0"
			i++
		}
		chuck += " "
	}
	return
}
