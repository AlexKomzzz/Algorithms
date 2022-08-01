// C. Расстояние
// https://contest.yandex.ru/contest/28412/problems/C/

package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var lenght, sizeWindow int
	fmt.Fscan(os.Stdin, &lenght, &sizeWindow)

	//Создаем слайз для хранения начального массива чисел и заполним его
	StartArr := make([]int, lenght)

	for n := 0; n < lenght; n++ {
		var Number int
		fmt.Fscan(os.Stdin, &Number)

		StartArr[n] = Number
	}

	// Создадим копию слайса - что очень плохо
	NewArr := make([]int, lenght)
	copy(NewArr, StartArr)

	// отсортируем новый слайз
	sort.Slice(NewArr, func(i int, j int) bool { return NewArr[i] < NewArr[j] })

	// Итерируясь по начальному слайсу, для каждого элемента найдем мин длину
	for _, elem := range StartArr {

		// Найдем индекс этого элемента в отсортированном слайсе
		ind := sort.Search(len(NewArr), func(i int) bool { return NewArr[i] >= elem })

		// Создадим два указателя
		Sum := 0
		left, right := 1, 1

		// создадим цикл с k(sizeWindow) шагами
		for l := 0; l < sizeWindow; l++ {

			// Необходимо предусмотреть граничные условия массива NewArr (начальный и коннечный индексы)
			if ind+right >= lenght {
				Sum += NewArr[ind] - NewArr[ind-left]
				left++
				continue
			} else if ind-left < 0 {
				Sum += NewArr[ind+right] - NewArr[ind]
				right++
				continue
			}
			// Рассматриваем ближайшие числа к элементу elem и обновляем сумму
			// Рассматриваем разницу от выбранного elem и ближайшего числа ( с левой или правой стороны).
			// выбираем то число, которое дает меньшую разницу
			if (NewArr[ind] - NewArr[ind-left]) < (NewArr[ind+right] - NewArr[ind]) {
				Sum += NewArr[ind] - NewArr[ind-left]
				left++
			} else {
				Sum += NewArr[ind+right] - NewArr[ind]
				right++
			}
		}

		fmt.Printf("%d ", Sum)

	}
}
