package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// スライスの書き方, 範囲はpythonと一緒
	var s []int = primes[1:4]
	fmt.Println(s)

	for p := range primes {
		fmt.Println(p)
	}

}
