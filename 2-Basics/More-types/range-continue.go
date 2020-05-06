package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// インデックだけ欲しい時
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// 値だけ欲しいとき
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
