package main

import (
		"fmt"
		// "strings"
		// "sort"
		// "os"
		// "log"
		// "io/ioutil"
		"strconv"
		// "net/http"
		"time"
		)

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNum++

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make dough and Send for Sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10 )
}

func addSauce(stringChan chan string) {
	pizza := <- stringChan

	fmt.Println("Add Sauce and Send ", pizza, " for toppings")
}

func addToppings(stringChan chan string) {
	pizza := <- stringChan

	fmt.Println("Toppings Added for ", pizza )
}

func main() {
		stringChan := make(chan string)

		for i := 0; i < 3; i++ {
			go makeDough(stringChan)
			go addSauce(stringChan)
			go addToppings(stringChan)
		}

		time.Sleep(time.Millisecond * 5000)
}