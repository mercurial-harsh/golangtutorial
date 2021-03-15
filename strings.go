package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("suman", "harsh"))
	fmt.Println(strings.Compare("Nitya", "nitya"))
	fmt.Println(strings.Contains("harsh", "harsh"))
	fmt.Print(strings.ContainsAny("harsh", "a"))
}
