package main

import "fmt"

//スライスのゼロ値はnil
// len 0 cap 0

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
