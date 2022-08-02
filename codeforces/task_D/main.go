//D. Покрась вхождениями
//https://codeforces.com/contest/1714/problem/D

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
	}
}
