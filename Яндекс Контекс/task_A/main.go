package main

import (
	"fmt"
	"os"
)

func main() {

	// Считываем кол-во карточек и кол-во ходов
	var n, k int
	fmt.Fscan(os.Stdin, &n)
	fmt.Fscan(os.Stdin, &k)

	fmt.Println(n + k)
}
