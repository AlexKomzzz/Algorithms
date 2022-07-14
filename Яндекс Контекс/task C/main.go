// Яндекс Контекст
// Задача C. Статус 200

// Решим методом

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdin)
	defer out.Flush()

	// Считываем кол-во карточек и кол-во ходов
	var n, k uint
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &k)

	fmt.Fprintln(out)
}
