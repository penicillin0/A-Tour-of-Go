package main

import "fmt"

func main() {
	i, j := 42, 2701

	// & でポインタ型を生成できる
	p := &i
	// * でポインタが指している変数を取得
	fmt.Println("*p", *p)
	// *p も変数なので代入可能
	*p = 21
	fmt.Println("*p に21を代入した後のi", i)

	p = &j
	*p = *p / 37
	fmt.Println("*p を37で割った後のj", j)
}
