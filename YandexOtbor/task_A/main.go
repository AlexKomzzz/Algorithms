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

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func ConvertInt10to2(val10 int64) (int64, error) {
	NewLetter2Str := strconv.FormatInt(val10, 2)
	NewLetter2, err := strconv.ParseInt(NewLetter2Str, 2, 64)
	if err != nil {
		return -1, err
	}
	return NewLetter2, nil
}

func main() {

	MapNumb := make(map[int]int64)
	// заполним мапу
	Map := make(map[int]string)

	for i := 0; i < 26; i++ {
		Map[i] = string(byte(97 + i))
	}
	Map[26] = " "

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var result string

	//Кол-во чисел
	var num uint16
	fmt.Fscan(in, &num)

	for n := 1; uint16(n) <= num; n++ {
		var number string
		fmt.Fscan(in, &number)

		Numb2Str, err := ConvertInt(number, 10, 2)
		if err != nil {
			fmt.Fprintln(out, err.Error())
			return
		}

		Numb2, err := strconv.ParseInt(Numb2Str, 2, 32)
		if err != nil {
			fmt.Fprintln(out, err.Error())
			return
		}

		Numb10, err := strconv.ParseInt(number, 10, 32)
		if err != nil {
			fmt.Fprintln(out, err.Error())
			return
		}

		MapNumb[n] = Numb10

		if n == 1 {
			Count := 0
			for j := 0; j <= 26; j++ {
				if Numb2 == 1 {
					result = Map[Count]
					break
				}
				Numb2 = Numb2 >> 1
				Count++
			}
		} else {

			NewLetter10 := Numb10 - MapNumb[n-1]
			NewLetter2, err := ConvertInt10to2(NewLetter10)
			if err != nil {
				fmt.Fprintln(out, err.Error())
				return
			}

			if NewLetter10 > 0 {
				Count := 0
				for j := 0; j <= 26; j++ {
					if NewLetter2 == 1 {
						result = result + Map[Count]
						break
					}
					NewLetter2 = NewLetter2 >> 1
					Count++
				}
			} else if NewLetter10 < 0 {
				NewLetter10 = -NewLetter10
				NewLetter2, err := ConvertInt10to2(NewLetter10)
				if err != nil {
					fmt.Fprintln(out, err.Error())
					return
				}

				Count := 0
				for j := 0; j <= 26; j++ {
					if NewLetter2 == 1 {
						result = result + Map[Count]
						break
					}
					NewLetter2 = NewLetter2 >> 1
					Count++
				}

			} else {
				fmt.Fprintln(out, "Error work")
			}
		}

	}
	fmt.Fprintln(out, result)
}
