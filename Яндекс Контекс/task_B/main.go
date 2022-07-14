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

	//start := time.Now()
	// Создаем слайз в котором будем хранить номера карточек
	NumbersSlice := make([]uint, n)

	// Считаем номера карточек с ввода и запишем в слайз
	for i := uint(0); i < n; i++ {
		var NumberOnCard uint
		fmt.Fscan(in, &NumberOnCard)

		NumbersSlice[i] = NumberOnCard
		/*
			// Для проверки больших рандомных массивов
			rand.Seed(time.Now().UnixNano())
			NumbersSlice[i] = uint(rand.Intn(10000 + 1))*/
	}

	// Время создания массива
	//duration1 := time.Since(start)

	// Измерим начало работы
	//start = time.Now()

	// Создадим переменную, в которой будет записана максимальная сумма чисел из начала массива от 0 до k-1 (левая часть массива)
	var MaxSum uint
	for i := uint(0); i < k; i++ {
		MaxSum += NumbersSlice[i]
	}

	ResultSum := MaxSum

	// Если k == n
	if k == n {
		fmt.Fprintln(out, ResultSum)
		return
	}

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

	//duration := time.Since(start)

	fmt.Fprintln(out, ResultSum)
	//fmt.Fprintln(out, duration1)
	//fmt.Fprintln(out, duration)

}
