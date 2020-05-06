package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = make(map[string]Vertex)

func main() {
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Google"] = Vertex{37.42202, -122.08408}
	fmt.Println(m)
}
