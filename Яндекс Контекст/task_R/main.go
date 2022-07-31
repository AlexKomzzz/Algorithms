package main

import (
	"fmt"
	"os"
)

func MaxD(numb int) int {
	i := numb / 2
	for numb%i != 0 {
		i--
	}

	return i
}

func MinD(numb int) int {
	i := 2
	for numb%i != 0 {
		i++
	}

	return i
}

func Woner(numb int, count *int) {
	if numb == 1 {
		return
	}
	*count++
	d := MinD(numb)

	//d := MaxD(numb)
	fmt.Println(d, " ", numb-d)
	Woner(numb-d, count)
}

func main() {
	var numb int
	fmt.Fscan(os.Stdin, &numb)
	var count int

	Woner(numb, &count)

	if count%2 == 0 || count == 0 {
		fmt.Println("Mark")
	} else {
		fmt.Println("Pasha")

	}
}
