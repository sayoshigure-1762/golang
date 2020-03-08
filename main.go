package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("initはmainよりも先に実行される")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	fmt.Println("####################################")
	fmt.Println("funcの呼び出し")
	bazz()
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())

	fmt.Println("####################################")
	fmt.Println("変数の宣言")
	var (
		i   int     = 1
		f64 float64 = 1.2
		s   string  = "test"
		t   bool    = true
		f   bool    = false
	)
	fmt.Println(i, f64, s, t, f)
}

//go run: ソースコードのコンパイルと実行を行う。
