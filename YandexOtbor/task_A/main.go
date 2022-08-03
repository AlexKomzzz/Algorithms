// Стажировка весна-лето 2022 | бэкенд
// A. Числовые ребусы
// https://contest.yandex.ru/contest/38818/problems/A/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Фун-ия конвертации из 10-тичной системы в 2-ичную
func ConvertInt10to2(val10 int64) (int64, error) {
	if val10 < 0 {
		val10 = -val10
	}
	NewLetter2Str := strconv.FormatInt(val10, 2)
	NewLetter2, err := strconv.ParseInt(NewLetter2Str, 2, 64)
	if err != nil {
		return -1, err
	}
	return NewLetter2, nil
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var CountEnterNum int
	fmt.Fscan(in, &CountEnterNum)

	// Переменная для запоминания предыдущего числа
	var OldNum int

	for n := 0; n < CountEnterNum; n++ {
		var NewBigNum int
		fmt.Fscan(in, &NewBigNum)

		// чтобы понять, какой разряд изменился, вычтем из нового числа предыдущее
		NewSmallNum := NewBigNum - OldNum
		OldNum = NewBigNum

		// переведем его в двоичную систему
		NewNum_2, err := ConvertInt10to2(int64(NewSmallNum))
		if err != nil {
			fmt.Println("error convert from 10 to 2")
		}

		// создадим счетчик разрядов
		var CountDigit byte

		for NewNum_2 != 1 {
			CountDigit++
			NewNum_2 = NewNum_2 >> 1
		}

		if CountDigit == 26 {
			fmt.Fprint(out, " ")
		} else {
			fmt.Fprint(out, string(CountDigit+97))
		}
	}
}
