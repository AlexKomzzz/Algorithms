//C. Минимальное разнообразное число
//https://codeforces.com/contest/1714/problem/C

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

		var StartNum int
		fmt.Fscan(in, &StartNum)

		ReturnSlice := make([]int, 0)

		for i := 9; i > 0; i-- {
			if StartNum > i {
				StartNum -= i
				ReturnSlice = append(ReturnSlice, i)
			} else {
				ReturnSlice = append(ReturnSlice, StartNum)
				break
			}
		}

		for i := len(ReturnSlice) - 1; i >= 0; i-- {
			fmt.Fprint(out, ReturnSlice[i])
		}

		fmt.Fprintln(out)
	}
}
