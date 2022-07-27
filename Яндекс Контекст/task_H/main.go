// H. Сизиф
// https://contest.yandex.ru/contest/36783/problems/H/
// В этой задаче вы будете перекладывать камни. Изначально есть n кучек камней. Кучка i весит ai килограммов. Кучки можно объединять.
// При объединении кучек i и j затрачивается ai + aj единиц энергии, при этом две исходные кучки пропадают и появляется кучка весом ai + aj.
// Определите наименьшее количество энергии, которое надо затратить для объединения всех кучек в одну.

package main

// При решении это задачи используем такую структуру данных, как куча.
// 1. Создаем кучу из всех элементов
// 2. Извлекаем два минимальных элемента из кучи
// 3. Их сумму запоминаем и вставляем в кучу
// 4. Повторяем пункт 2, пока в куче не останется элементов

import (
	"container/heap"
	"fmt"
	"os"
)

// IntHeap это минимальная куча целых чисел.
type IntHeap []int

// Len, Less, Swap для реализации интерфейса sort.Interface
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push и Pop используют приемники указателей,
	// потому что они изменяют длину среза,
	// не только его содержимое.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	var amount int
	fmt.Fscan(os.Stdin, &amount)

	MyHeap := &IntHeap{}

	// Заполним кучу
	for n := 0; n < amount; n++ {
		var bunch int
		fmt.Fscan(os.Stdin, &bunch)

		heap.Push(MyHeap, bunch)
	}

	// Переменная для хранения затраченной энергии
	var Energy int

	for n := 0; n < amount-1; n++ {

		// Извлекаем два минимальных числа
		NewEnergy := heap.Pop(MyHeap).(int) + heap.Pop(MyHeap).(int)
		Energy += NewEnergy

		// Выход из цикла, если куча пустая
		//if MyHeap.Len() == 0 {
		//	break
		//}

		// Их сумму добавляем в кучу
		heap.Push(MyHeap, NewEnergy)
	}

	fmt.Print(Energy)
}
