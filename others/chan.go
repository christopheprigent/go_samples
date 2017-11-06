package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("--------------")
	fmt.Println("Make Dough for", pizzaName, "and send to the sauce")
	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 1)
}

func makeSauce(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("Make Sauce and send", pizza, "for toppings")
	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 1)
}

func addToppings(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("Add toppings to", pizza, "and ship")
	time.Sleep(time.Millisecond * 1)
}

func main() {

	stringChan := make(chan string)

	for i := 0; i < 4; i++ {
		go makeDough(stringChan)
		go makeSauce(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond * 50)
	}

}
