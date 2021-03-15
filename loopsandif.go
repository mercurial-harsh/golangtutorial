package main

import "fmt"

func main() {

	var name string = "nitya"

	if name == "nity" {
		fmt.Println("nity")
	} else if name == "nitya" {
		fmt.Println("nitya")
	}

	const i int = 6

	for k := 0; k < i; k++ {
		fmt.Println(k)
	}

	names := []string{"harsh", "nitya", "seema", "us"}
	for j, l := range names {
		fmt.Println(l, j)
	}

	arraydisplay(names)

}

func arraydisplay(x []string) {
	for i, j := range x {
		fmt.Println(i, j)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i)
	}
}
