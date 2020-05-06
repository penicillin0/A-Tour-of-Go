package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ans := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		ans[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			ans[y][x] = uint8(x ^ y)
		}
	}
	return ans
}

func main() {
	pic.Show(Pic)
}
