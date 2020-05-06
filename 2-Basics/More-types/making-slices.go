package main

import "fmt"

// スライスの作り方
// make関数で動的なスライスを作成
// 第三引数で、capを指定できる

func main() {
	a := make([]int, 5) // len(a) = 5
	printSlice("a", a)

	b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	for 
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
