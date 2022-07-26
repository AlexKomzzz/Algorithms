// H. Сизиф
// https://contest.yandex.ru/contest/36783/problems/H/
// В этой задаче вы будете перекладывать камни. Изначально есть n кучек камней. Кучка i весит ai килограммов. Кучки можно объединять.
// При объединении кучек i и j затрачивается ai + aj единиц энергии, при этом две исходные кучки пропадают и появляется кучка весом ai + aj.
// Определите наименьшее количество энергии, которое надо затратить для объединения всех кучек в одну.

package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
)

func main() {

	var amount int
	fmt.Fscan(os.Stdin, &amount)

	BunchSlice := make([]*big.Int, amount)

	for n := 0; n < amount; n++ {
		var bunch int64
		fmt.Fscan(os.Stdin, &bunch)

		BunchSlice[n] = big.NewInt(bunch)
	}

	// сортируем кучи по неубываю весов
	sort.Slice(BunchSlice, func(i int, j int) bool { return BunchSlice[i].Int64() <= BunchSlice[j].Int64() })

	bunchNew := big.NewInt(0)

	for n := 1; n < amount; n++ {

		sum := big.NewInt(0)
		bunchNew.Add(bunchNew, sum.Add(BunchSlice[n], BunchSlice[n-1]))
		BunchSlice[n] = bunchNew
	}

	fmt.Print(bunchNew.Uint64())
}
