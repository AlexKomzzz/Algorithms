// Стажировка весна-лето 2022 | бэкенд
// D. Двоичная медиана
// https://contest.yandex.ru/contest/38818/problems/D/

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pref struct {
	Count0, Count1 int32
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var LenStr int32
	fmt.Fscan(os.Stdin, &LenStr)

	var Numbers []byte
	fmt.Fscan(os.Stdin, &Numbers)

	//var Count0, Count1 uint32
	SliceNum := make([]Pref, LenStr+1) // первый ноль не будем учитывать

	for N := int32(1); N <= LenStr; N++ {

		var Result int32 = -1

		// Считаем каждое число последовательно в виде байт
		// 0-48   1-49 и тд.
		NewNum := Numbers[N-1]

		if NewNum == 48 {
			// если это 0, увеличиваем счетчик нулей
			SliceNum[N].Count0++
			for i := int32(1); i < N; i++ {
				SliceNum[i].Count0++

				if SliceNum[i].Count0 > SliceNum[i].Count1 {
					Result = i
				}
			}
		} else {
			// если это 1, то увеличиваем счетчик единиц
			SliceNum[N].Count1++
			for i := int32(1); i < N; i++ {
				SliceNum[i].Count1++

				if SliceNum[i].Count1 > SliceNum[i].Count0 {
					Result = i
				}
			}
		}

		// По определению
		if N == 1 {
			Result = -1
		}

		fmt.Fprint(out, Result, " ")
	}
}
