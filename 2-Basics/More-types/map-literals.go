package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		}, // <= ,がないと構文エラーになる、意外
	}
	fmt.Println(m)
}
