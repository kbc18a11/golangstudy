package main // mainパッケージであることを宣言
import "fmt"

// fmtモジュールをインポート

func main() { // 最初に実行されるmain()関数を定義
	const homo = "a"

	age := 26.0

	a := 2
	age += float64(a)
	fmt.Println(age)
}
