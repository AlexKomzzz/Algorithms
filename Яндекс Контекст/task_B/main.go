// Яндекс Контекст
// Задача B. Card Counter

// Решим методом скользящего окна
package main

import (
	"fmt"
	"os"
)

func Sum(Slice []uint) uint {
	if len(Slice) == 2 {
		return Slice[0] + Slice[1]
	}
	if len(Slice) == 1 {
		return Slice[0]
	} else {
		midle := uint(len(Slice)) / 2
		return Sum(Slice[:midle+1]) + Sum(Slice[midle+1:])
	}

}

func main() {
	/*in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdin)
	defer out.Flush()*/

	// Считываем кол-во карточек и кол-во ходов
	var n, k uint
	fmt.Fscan(os.Stdin, &n)
	fmt.Fscan(os.Stdin, &k)

	//start := time.Now()
	// Создаем слайз в котором будем хранить номера карточек
	NumbersSlice := make([]uint, n)

	// Считаем номера карточек с ввода и запишем в слайз
	for i := uint(0); i < n; i++ {
		var NumberOnCard uint
		fmt.Fscan(os.Stdin, &NumberOnCard)

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

	// Создадим переменную, в которой будет записана максимальная сумма чисел из начала массива от 0 до k-1 (левая часть массива - k левых карт)
	var MaxSum uint
	for i := uint(0); i < k; i++ {
		MaxSum += NumbersSlice[i]
	}

	ResultSum := MaxSum
	// Если k == n
	if k == n {
		fmt.Println(ResultSum)
		return
	}

	// Сдвигаем окно, уменьшая начало(левую часть), увеличия конец(правую часть). Необходимо уменьшать кол-во левых карт и увеличивать правые, при этом запоминать макс сумму
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

	fmt.Println(ResultSum)
	//fmt.Fprintln(out, duration1)
	//fmt.Fprintln(out, duration)

}
