package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println("スライスで切り取った後 a, b => ", a, b)

	b[0] = "XXX"
	fmt.Println("スライスの変数b[0](Paulのとこ)を変更すると a, b => ", a, b)
	fmt.Println("参照渡しなので元の配列も変更される => ", names)
}
