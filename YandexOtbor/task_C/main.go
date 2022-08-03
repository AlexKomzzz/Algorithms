// Стажировка весна-лето 2022 | бэкенд
// C. Домашнее задание

// «Даны два целых числа A и B, можно ровно один раз умножить либо A, либо B на любое простое число.
// Какого наибольшего значения НОД можно добиться с помощью такого умножения?»
package main

// Найти НОД для введеных чисел (алгоритм Евклида)
// определить на какое число Х умножить А или В, чтобы получить макс НОД
// определить является ли Х простым числом
// Если является, вывести НОД при значении Х

import (
	"bufio"
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

// Поиск наибольшего числа, на которое умноженить начальные А и В
func SearchX(num, oldNOD int, i *int) int {
	X := num / oldNOD
	*i++
	iOk := false
	for iOk {
		if num%*i == 0 {
			iOk = true
		} else {
			*i++
		}
	}

	return X
}

// Фун-ия поиска наибольшего числа для умножения и его проверка на простое число
func SimpleNum(i *int, num, oldNOD int) bool {
	prime := true
	X := SearchX(num, oldNOD, i)
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
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tasks uint8
	fmt.Fscan(os.Stdin, &tasks)

	for T := uint8(0); T < tasks; T++ {

		// Ввод первого и второго числа
		var A, B int
		fmt.Fscan(os.Stdin, &A)
		fmt.Fscan(os.Stdin, &B)

		var num int
		var l *int
		// Создадим счетчики на которые будем делить А и В для нахождения наибольшего делителя
		var i, j int = 1, 1
		var ok bool

		// Поиск числа на которое нужно увеличить начальное число
		for !ok {

			if A/i > B/j || (A/i == B/j && i <= j) {
				num = A / i
				l = &i
			} else {
				num = B / j
				l = &j
			}

			ok = SimpleNum(l, num, NOD(A, B))
		}
		fmt.Fprintln(out, num)
	}
}
