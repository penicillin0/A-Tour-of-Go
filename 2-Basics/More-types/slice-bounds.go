package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// この辺のノリはpythonと同じ
	// スライスは参照渡しということだけ常に注意
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
