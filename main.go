package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, upper string) {
	defer fmt.Println("I'm done")
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func main() {
	len, upper := lenAndUpper("Hello")
	fmt.Println(len, upper)
}
