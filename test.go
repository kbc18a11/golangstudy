package main // mainパッケージであることを宣言

// fmtモジュールをインポート

func main() { // 最初に実行されるmain()関数を定義
	const homo = 'aaa'

	name, age := "Yamada", 26

	homo = name 
	print(name)
	print(age)
}
