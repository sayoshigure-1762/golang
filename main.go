package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"time"
)

// Vertex struct
type Vertex struct {
	X int
	Y int
	S string
} //構造体は他の言語でいうところのclassみたいなもの。ただし、methodを含めることはできない。

func init() {
	fmt.Println("initはmainよりも先に実行される")
}

/*
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
*/

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

	fmt.Println("####################################")
	fmt.Println("switch文")
	tm := time.Now()
	fmt.Println(tm.Hour())
	switch {
	case tm.Hour() < 12:
		fmt.Println("Morning!!")
	case tm.Hour() < 17:
		fmt.Println("Afternoon!!")
	}
	fmt.Println("####################################")
	fmt.Println("defer")
	file, _ := os.Open("./foo.go")
	defer file.Close() //処理の最後に実行する、かつ記述忘れを防止するために最初に書いておきたい場合に使う？
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))

	fmt.Println("####################################")
	fmt.Println("logging")
	LoggingSettings("test.log")

	fmt.Println("####################################")
	fmt.Println("panic/recover")
	save()

	fmt.Println("####################################")
	fmt.Println("演習")
	l1 := []int{100, 300, 500, 400, 200, 50}
	l2 := []int{50, 300, 500, 400, 200, 70}
	l3 := []int{100, 300, 50, 400, 200, 80}
	answer1 := searchMinimumNumber(l1)
	answer2 := searchMinimumNumber(l2)
	answer3 := searchMinimumNumber(l3)
	fmt.Println(answer1, answer2, answer3)

	fmt.Println("####################################")
	fmt.Println("ポインタ")
	var nn int = 100
	var pp *int = &nn //ppという変数にnnのアドレスを代入。*をつけるとポインタ型となる。
	fmt.Println(pp, *pp)

	var ppp *int = new(int) //メモリ領域を確保する。newとmakeの違い: newはポインタ、makeはスライスやmap??
	fmt.Println(ppp, *ppp)

}
