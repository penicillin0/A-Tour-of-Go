package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // 暗黙的に Y = 0
	v3 = Vertex{}      // 暗黙的に X = 0, Y = 0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3)
	fmt.Println(p)
}
