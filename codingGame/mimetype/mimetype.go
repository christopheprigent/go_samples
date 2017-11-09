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

	mimes := make([]_mime, N)
	for i := 0; i < N; i++ {
		m := _mime{}
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &m.ext, &m.mimeValue)
		m.ext = s.ToLower(m.ext)
		mimes[i] = m
	}
	for i := 0; i < Q; i++ {
		scanner.Scan()
		FNAME := s.ToLower(scanner.Text()) // One file name per line.
		currentMimeValue := "UNKNOWN"
		for i := 0; i < N; i++ {
			if s.HasSuffix(FNAME, "."+mimes[i].ext) {
				currentMimeValue = mimes[i].mimeValue
				break
			}
		}
		fmt.Println(currentMimeValue)
	}
}

type _mime struct {
	ext, mimeValue string
}
