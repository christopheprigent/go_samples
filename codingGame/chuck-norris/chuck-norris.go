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

	for _, char := range msg {
		bin += fmt.Sprintf("%07b", char)
	}
	fmt.Println(Bin2Chuck(bin)) // Write answer to stdout
}

// Bin2Chuck : Chuck Norris' keyboard has 2 keys: 0 and white space.
func Bin2Chuck(bin string) (chuck string) {
	chuck = ""
	for i := 0; i < len(bin); i++ {
		c := string(bin[i])
		if i > 0 {
			chuck += " "
		}
		if c == "0" {
			chuck += "00 0"
		} else {
			chuck += "0 0"
		}
		for j := i + 1; j < len(bin) && string(bin[j]) == c; j++ {
			chuck += "0"
			i++
		}

	}
	return
}
