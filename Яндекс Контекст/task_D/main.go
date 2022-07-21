// Задача D. Matrix. Resurrection
// https://contest.yandex.ru/contest/36783/problems/D/
//
// Представим числа в матрице в виде ориентированного ациклического графа. Вершины представим структурой Vertex. Будем хранить в двумерном массиве
// направление ребер определяется от наименьшего значения вершины к наибольшему
// В данной работе вершины, на которые указывает вершина-родитель, называются потомками, соседями(Neighbor) или child.
// Кол-во вершин = n*m, кол-во ребер = (n-1)*m + (m-1)*n
// при добавлении соседей будем смотреть на Value, значение определяет направление ребра (от меньшего к большему).
// Смотрим на левого и верхнего соседа и сравниваем значения.
// Необходимо построить ориентированный граф и применить рекурсивный поиск наибольшего пути.
// Пройденые вершины помечаем в поле Lable и уже не будем их проходить, т.к. в поле Count хранится максимальный путь от них
// Нумерация вершин начинается от 1 до n*m. Формула для нахождения номера вершины (i-1)*m +j

package main

import (
	"fmt"
	"math"
	"os"
)

// Координата вершины

// Структура вершины
// поля: соседи, значение, пометка о проверке вершины, ее максимальный путь по соседям
type Vertex struct {
	Neighbor []*Vertex
	Value    float64
	Lable    bool
	Count    int32
}

// Конструктор для создания матрицы вершин
func NewTableVertex(n, m int32) [][]Vertex {
	TableVertex := make([][]Vertex, n+1)

	for i := int32(0); i <= n; i++ {

		// Слайз для строк в матрице(таблице)
		TableVertex[i] = make([]Vertex, m+1)

		// Для вершин из первой строки и первого столбца установим поле Value в минус бесконечность
		TableVertex[i][0].Value = math.Inf(-1)

		for j := int32(0); j <= m; j++ {
			TableVertex[0][j].Value = math.Inf(-1)
		}
	}

	return TableVertex
}

// Фун-я поиска максим. пути
func MaxLenVert(vert *Vertex) int32 {

	// Помечаем вершину как пройденную
	vert.Lable = true

	// Баховый случай - если нет потомков, то путь равен 1
	if vert.Neighbor == nil {
		vert.Count = 1
		return vert.Count
	}

	// Переменная для хранения макс. пути потомка
	var MaxLenChild int32

	// Раассмотрим всех потомков вершины
	for _, child := range vert.Neighbor {

		// Если потомок уже был пройден, запомним его путь для дальнейшего сравнения, если он больше других потомков
		if child.Lable {
			if MaxLenChild < child.Count {
				MaxLenChild = child.Count
			}
		} else {
			// Если потомок не был пройден, вызовем для него функцию MaxLen() и сравним с другими потомками
			if ChildCount := MaxLenVert(child); MaxLenChild < ChildCount {
				MaxLenChild = ChildCount
			}
		}
	}

	// Макс. путь вершины определяется как максимальный путь потомка + 1
	vert.Count = MaxLenChild + 1

	return vert.Count
}

func main() {
	var n, m int32

	fmt.Fscan(os.Stdin, &n)
	fmt.Fscan(os.Stdin, &m)

	// Слайз для хранения вершин
	TableVertex := NewTableVertex(n, m)

	//Заполним матрицу вершинами
	for i := int32(1); i <= n; i++ {
		for j := int32(1); j <= m; j++ {

			// Считаем число из матрицы и запишем в вершину
			var val float64
			fmt.Fscan(os.Stdin, &val)

			TableVertex[i][j].Value = val

			// Сравниваем значения Value с соседями слева и сверху (сверху находятся соседи с j-1)
			// Если значение Value соседа больше текущей вершины, то добавим соседа в список соседей к текущей вершине,
			// если меньше - к соседней вершине добавим в список соседей текущую вершину
			if val < TableVertex[i-1][j].Value {
				TableVertex[i][j].Neighbor = append(TableVertex[i][j].Neighbor, &TableVertex[i-1][j])
			} else if val > TableVertex[i-1][j].Value {
				TableVertex[i-1][j].Neighbor = append(TableVertex[i-1][j].Neighbor, &TableVertex[i][j])
			}

			if val < TableVertex[i][j-1].Value {
				TableVertex[i][j].Neighbor = append(TableVertex[i][j].Neighbor, &TableVertex[i][j-1])
			} else if val > TableVertex[i][j-1].Value {
				TableVertex[i][j-1].Neighbor = append(TableVertex[i][j-1].Neighbor, &TableVertex[i][j])
			}
		}
	}

	// Ищем наибольший путь рекурсивным методом
	// макс. путь до вершины определяется как максимальный путь ее потомка + 1 и хранится в поле Count
	// в процессе поиска, исследованным вершинам устанавливаем Lable = true, т.к. нет смысла с нее начинать поиск и продолжать поиск, если встретили такую вершину
	// т.е. для вершины с Lable = true максимальный путь уже посчитан.

	// Создадим переменную для хранения максимального пути
	var MaxLenght int32

	// т.к. слайз мы можем проходить только от 0 до n, и если у нас будет матрица содержать оргомное число постоянно убывающих чисел
	// то нам продется для каждой вершины считать путь.
	// Попробуем отдельно посчитать путь для вершины в середине матрицы
	MaxLenght = MaxLenVert(&TableVertex[n/2][m/2])

	// Определим макс путь для каждой вершины и выведем наибольший из них
	for i := int32(1); i <= n; i++ {
		for j := int32(1); j <= m; j++ {

			// Пройденные вершины будем пропускать
			if TableVertex[i][j].Lable {
				continue
			}

			if MaxLenVertex := MaxLenVert(&TableVertex[i][j]); MaxLenght < MaxLenVertex {
				MaxLenght = MaxLenVertex
			}
		}
	}

	fmt.Println(MaxLenght)
}
