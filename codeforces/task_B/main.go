// B. Удали префикс
// https://codeforces.com/contest/1714/problem/B
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var CountData int
	fmt.Fscan(in, &CountData)

	for t := 0; t < CountData; t++ {
		var Lenght int
		fmt.Fscan(in, &Lenght)

		MapElem := make(map[int]int)
		EndInd := -1

		for n := 0; n < Lenght; n++ {
			var Elem int
			fmt.Fscan(in, &Elem)

			if Ind, ok := MapElem[Elem]; ok {
				if EndInd < Ind {
					EndInd = Ind
				}
			}
			MapElem[Elem] = n
		}

		fmt.Fprintln(out, 1+EndInd)
	}
}
