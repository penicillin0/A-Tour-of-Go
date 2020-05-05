package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum *= 2
		fmt.Println(sum)
	}
}
