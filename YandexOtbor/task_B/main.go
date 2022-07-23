// Стажировка весна-лето 2022 | бэкенд
// B. Шестиугольники
//https://contest.yandex.ru/contest/38818/problems/B/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Центр фигуры. Где Х - это номер строки, Y - столбца
type Centr struct {
	val rune
	X   int8
	Y   int8
}

func Reflection(sl []Centr, n, m int8) {
	for i := range sl {
		sl[i].X = n - sl[i].X + 1
		sl[i].Y = m - sl[i].Y + 1
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	SliseCentr := make([]Centr, 0)

	var n, m int8
	fmt.Fscan(in, &n, &m)

	for i := int8(0); i < n; i++ {

		var InpStr string
		fmt.Fscan(in, &InpStr)

		//Если символ - буква, то сохраняем ее координаты
		for j, Rune := range InpStr {
			if unicode.IsLetter(Rune) {
				SliseCentr = append(SliseCentr, Centr{val: Rune, X: i + 1, Y: int8(j) + 1})
			}
		}
	}

	// Узнали координаты букв
	// Отразим их по вертикали и горизонтали
	Reflection(SliseCentr, n, m)

	// Создаем поле
	//column := make([]string, m)
	field := make([][]string, n)
	for i := int8(0); i < n; i++ {
		field[i] = make([]string, m)
	}

	dx := [6]int8{1, 1, 1, 0, 0, -1}
	dy := [6]int8{-1, 0, 1, -1, 1, 0}
	dval := [6]string{"\\", "_", "/", "/", "\\", "_"}

	// Ставим буквы на места и строим ромб
	for _, point := range SliseCentr {
		field[point.X-1][point.Y-1] = string(point.val)
		// строим ромб вокруг буквы
		for k := uint8(0); k < 6; k++ {
			field[point.X-1+dx[k]][point.Y-1+dy[k]] = dval[k]
		}
	}

	// Заполняем пустые места точками
	for i := int8(0); i < n; i++ {
		for j := int8(0); j < m; j++ {
			if field[i][j] == "" {
				field[i][j] = "."
			}
		}
	}

	// вывод
	for i := int8(0); i < n; i++ {
		for j := int8(0); j < m; j++ {
			fmt.Fprint(out, field[i][j])
		}
		fmt.Fprintln(out)
	}

}
