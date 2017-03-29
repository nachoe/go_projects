package main

import (
		"fmt"
		// "strings"
		// "sort"
		// "os"
		// "log"
		// "io/ioutil"
		// "strconv"
		"net/http"
		)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r * http.Request) {
	fmt.Fprint(w, "Hello more World\n")
}