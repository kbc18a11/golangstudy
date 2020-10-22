package main // mainパッケージであることを宣言
import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(5, 3)) // => 8
}
