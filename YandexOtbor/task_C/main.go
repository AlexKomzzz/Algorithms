// Стажировка весна-лето 2022 | бэкенд
// C. Домашнее задание

// «Даны два целых числа A и B, можно ровно один раз умножить либо A, либо B на любое простое число.
// Какого наибольшего значения НОД можно добиться с помощью такого умножения?»
package main

// Найти НОД для введеных чисел
// определить на какое число Х умножить А или В, чтобы получить макс НОД
// определить является ли Х простым числом

import (
	"fmt"
	"math"
	"os"
)

// Поиск НОД
func NOD(num1, num2 int) int {
	// Большее число сделаем первым
	if num1 < num2 {
		num1, num2 = num2, num1
	}

	// Найдем НОД
	for {
		if num1%num2 == 0 {
			return num2
		} else {
			num2 = num1 - num2*(num1/num2)
		}
	}
}

// Проверка на простое число
func SimpleNum(X int) bool {
	prime := true

	if X <= 2 {
		prime = false
		return prime
	}

	for i := float64(2); i <= math.Sqrt(float64(X)); i++ {
		if X%int(i) == 0 {
			prime = false
			break
		}
	}

	return prime
}

func main() {
	var tasks uint8
	fmt.Fscan(os.Stdin, &tasks)

	Result := make([]int, 0)

	for T := uint8(0); T < tasks; T++ {

		// Ввод первого и второго числа
		var A, B int
		fmt.Fscan(os.Stdin, &A)
		fmt.Fscan(os.Stdin, &B)

		//Найдем НОД
		oldNOD := NOD(A, B)

		num1, num2 := A, B
		// Большее число сделаем первым
		//if num1 < num2 {
		//	num1, num2 = num2, num1
		//}

		var ok bool
		var X int = 2
		var i, j int = 1, 1

		for {
			//for k := 0; k < 4; k++ {
			//fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
			//fmt.Println("NOD = ", NOD(num1, num2))

			if num1 > num2 {

				// if num1%NOD(num1, num2) == 0 ????
				X = num1 / oldNOD
				i++
				iOk := false
				for iOk {
					if num1%i == 0 {
						iOk = true
					} else {
						i++
					}
				}
			} else if num1 < num2 {
				// if num2%NOD(num1, num2) == 0 ????
				X = num2 / oldNOD
				j++
				jOk := false
				for jOk {
					if num2%j == 0 {
						jOk = true
					} else {
						j++
					}
				}
			}
			// if num1 == num2 ???

			ok = SimpleNum(X)
			if ok {
				break
			}

			if A/i > num2 {
				num1 = A / i
			} else {
				num2 = B / j
			}
			// if num1 == num2 ????

		}

		NOD1 := NOD(A*X, B)
		NOD2 := NOD(A, B*X)
		if NOD1 > NOD2 {
			Result = append(Result, NOD1)
		} else {
			Result = append(Result, NOD2)

		}

	}

	for _, ResultNODs := range Result {
		fmt.Println(ResultNODs)
	}

}
