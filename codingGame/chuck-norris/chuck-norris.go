package main

import "fmt"
import "os"
import "bufio"

//import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	MESSAGE := scanner.Text()

	fmt.Fprintln(os.Stderr, "Debug messages...", MESSAGE)
	fmt.Println("answer") // Write answer to stdout
}
