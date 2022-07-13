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

	var n, k uint
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &k)
	Cards, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(out, "error read string cards")
		return
	}

	fmt.Fprintln(out, Cards)
}
