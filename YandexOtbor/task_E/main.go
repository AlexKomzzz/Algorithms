package main

import (
	"fmt"
	"os"
	"sort"
)

type Rectangle struct {
	number, x1, y1, x2, y2, intersections int32
}

func main() {
	var amount int32
	fmt.Fscan(os.Stdin, &amount)

	// Слайз в котором храним прямоугольники
	RectangleSlice := make([]Rectangle, amount)

	// Сохраним все координаты прямоугольников
	for n := int32(0); n < amount; n++ {
		var x1, y1, x2, y2 int32
		fmt.Fscan(os.Stdin, &x1, &y1, &x2, &y2)

		RectangleSlice[n].x1 = x1
		RectangleSlice[n].y1 = y1
		RectangleSlice[n].x2 = x2
		RectangleSlice[n].y2 = y2
		RectangleSlice[n].number = n + 1
	}

	// условия для определения пересекаемости прямоугольников (* - координата второго прямоугольника)
	// X1 < *x2
	// X2 > *x1
	// Y1 < *y2
	// Y2 > *y1

	// Отсортируем прямоугольники по координате x2
	sort.Slice(RectangleSlice, func(i, j int) bool { return RectangleSlice[i].x2 < RectangleSlice[j].x2 })

	for k, rectangle := range RectangleSlice {
		// не будем рассматривать прямоугольники, которые уже рассмотрели
		NewSlice := RectangleSlice[k+1:]

		// Будем рассматривать только ту часть слайса, где у прямоугольников х2 больше или равно х1 текущего прямоугольника
		start := sort.Search(len(NewSlice), func(i int) bool { return NewSlice[i].x2 >= rectangle.x1 })
		NewRectangleSlice := NewSlice[start:]

		for l, AnotherRectangle := range NewRectangleSlice {

			// проверим остальные условия
			if rectangle.x2 <= AnotherRectangle.x1 || rectangle.y1 >= AnotherRectangle.y2 || rectangle.y2 <= AnotherRectangle.y1 {
				continue
			}

			RectangleSlice[k].intersections++
			NewRectangleSlice[l].intersections++

		}
	}

	// Отсортируем прямоугольники по номеру
	sort.Slice(RectangleSlice, func(i, j int) bool { return RectangleSlice[i].number < RectangleSlice[j].number })

	for i := range RectangleSlice {
		fmt.Println(RectangleSlice[i].intersections)
	}

}
