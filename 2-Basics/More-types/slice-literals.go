package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// ここ少しわかりにくい
	// 構造体を要素にもつスライス
	// Hoge{1 , true}みたいに作成しているのと同じこと↓
	// Hoge ≡ struct {
	//	i int
	//	b bool
	//}
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
