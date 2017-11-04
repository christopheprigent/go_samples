package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	myFileName := "file_test.txt"
	file, err := os.Create(myFileName)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("hello 42 42é")
	file.Close()

	stream, err := ioutil.ReadFile(myFileName)
	if err != nil {
		fmt.Println("err :")
		log.Fatal(err)
	}
	readString := string(stream)

	fmt.Println("content :", readString)
	// stream.Close()
}
