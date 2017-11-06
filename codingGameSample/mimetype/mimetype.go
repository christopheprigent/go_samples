package main

import "fmt"
import "os"
import "bufio"
import s "strings"

/**
 * Usage : cat mimetype-input3.txt | go run mimetype.go
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var N, Q int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &Q)

	mimes := make([]Mime, N)
	for i := 0; i < N; i++ {
		m := Mime{}
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &m.ext, &m.mimeValue)
		m.ext = s.ToLower(m.ext)
		mimes[i] = m
	}
	var found = false
	for i := 0; i < Q; i++ {
		scanner.Scan()
		FNAME := s.ToLower(scanner.Text()) // One file name per line.
		found = false
		for i := 0; i < N; i++ {
			if s.HasSuffix(FNAME, "."+mimes[i].ext) {
				fmt.Println(mimes[i].mimeValue)
				found = true
				break
			}
		}
		if found == false {
			fmt.Println("UNKNOWN")
		}

	}
}

type Mime struct {
	ext, mimeValue string
}
type Mimes []Mime
