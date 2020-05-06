package main

import "fmt"

// 可変スライスは長さが足りなくなったら都度容量が倍になる

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 0)
	printSlice(s)
	s = append(s, 0)
	printSlice(s)
	for i := 0; i < 10; i++ {
		s = append(s, 0)
		printSlice(s)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
