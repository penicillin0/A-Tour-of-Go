package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// 本来は(*p).Xとするべきだが↓のように表記可能
	p.X = 1e9
	(*p).Y = 1
	fmt.Println(v)
}
