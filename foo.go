package main

import (
	"fmt"
)

/*
func init() {
	fmt.Println("initはmainよりも先に実行される")
}
*/

func bazz() {
	fmt.Println("Bazz")
}

func add(x, y int) (result int) {
	fmt.Println("Add function.")
	return result
} //関数の中でresultという変数を定義して返すことができる。

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	} //定義された変数xはメモリ上で保持され続けるため、この関数を呼び出すたびにreturnされるxの値は1ずつ増える。
}

/*
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

	fmt.Println("####################################")
	fmt.Println("配列の宣言")
	n := []int{1, 2, 3, 4, 5, 6} //配列を可変長にしたい場合、[]として宣言する。
	fmt.Println(n)
	fmt.Println(n[1:2]) //n[2]の要素のみ表示される。スタート位置がn[1]とn[2]の間であるとイメージする。
	fmt.Println(n[:4])  //n[0]からn[3]まで表示される。
	n = append(n, 7)    //配列に要素を追加する。
	fmt.Println(n[:])

	fmt.Println("####################################")
	fmt.Println("スライスの宣言")
	var c []int
	//c = make([]int, 5)
	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	fmt.Println("####################################")
	fmt.Println("map: pythonの辞書型と似ているもの")
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"]) //対応するvalueが返される。
	m["new"] = 500
	fmt.Println(m)

	fmt.Println("####################################")
	fmt.Println("型変換")
	var ii int = 1
	iii := string(ii)
	fmt.Printf("%T, %v, %s\n", iii, iii, iii)

	var ss string = "14"
	sss, _ := strconv.Atoi(ss) //str型からint型への変換はpythonのように記述できない。専用のメソッドを使用する。
	fmt.Printf("%T, %v, %d\n", sss, sss, sss)
}
*/
