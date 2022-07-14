// Яндекс Контекст
// Задача B. Card Counter

// Решим методом скользящего окна
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

	// Считываем кол-во карточек и кол-во ходов
	var n, k uint
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &k)

	// Создаем слайз в котором будем хранить номера карточек
	NumbersSlice := make([]uint, n)

	// Считаем номера карточек с ввода и запишем в слайз
	for i := uint(0); i < n; i++ {
		var NumberOnCard uint
		fmt.Fscan(in, &NumberOnCard)

		NumbersSlice[i] = NumberOnCard
	}

	// Создадим переменную, в которой будет записана максимальная сумма чисел из начала массива от 0 до k-1 (левая часть массива)
	var MaxSum uint
	for i := uint(0); i < k; i++ {
		MaxSum += NumbersSlice[i]
	}

	ResultSum := MaxSum

	// Сдвигаем окно, уменьшая начало(левую часть), увеличия конец(правую часть).
	// Из суммы вычитаем последний элемент из левой части и прибавляем новый элемент из правой
	for i := uint(0); i < k; i++ {

		MaxSum -= NumbersSlice[k-1-i]
		MaxSum += NumbersSlice[n-1-i]
		// Если новая сумма больше, обновляем результат
		if ResultSum < MaxSum {
			ResultSum = MaxSum
		}
	}

	fmt.Fprintln(out, ResultSum)
}
