package main

import "fmt"

// オブジェクト指向言語でいうところのclass
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
