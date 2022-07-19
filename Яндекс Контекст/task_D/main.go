// Задача D. Matrix. Resurrection
// https://contest.yandex.ru/contest/36783/problems/D/
//
// Представим числа в матрице в виде графа. Граф представим с помощью матрицы смежности или списка???
// Кол-во вершин = n*m, кол-во ребер = (n-1)*m + (m-1)*n
// при добавлении соседей будем смотреть на Value, значение определяет направление ребра (от меньшего к большему)
// Необходимо построить ориентированный граф и применить топографическую сортировку
// Нумерация вершин начинается от 1 до n*m. Формула для нахождения номера вершины (i-1)*m +j

package main

import (
	"container/list"
	"fmt"
	"os"
)

// Координата вершины

// Структура вершины
// поля: ее координата, соседи, значение, пометка о проверке вершины
type Vertex struct {
	neighbor   []int32
	Value      uint32
	LableStart bool
	Count      int32
}

func main() {
	var n, m int32

	fmt.Fscan(os.Stdin, &n)
	fmt.Fscan(os.Stdin, &m)

	// Вершины храним в мапе
	VertexMap := make(map[int32]Vertex)

	// Заполним мапу
	for i := int32(1); i <= n; i++ {
		for j := int32(1); j <= m; j++ {

			// Считаем число из матрицы и запишем в вершину
			var val uint32
			fmt.Fscan(os.Stdin, &val)

			// Создание слайса, в котором хранятся соседи
			verifiedSlice := make([]int32, 0)
			//verifieds := []int32{-m, m, -1, 1}

			// Создадим вершину
			vertex := Vertex{
				neighbor: verifiedSlice,
				Value:    val,
			}

			// Если вершина находится в первой или последней строке матрицы, она имеет 2(вершина вы углу) или 3 соседа
			// Рассмотрим соседей для угловых вершин
			switch {

			// Если это первый элемент пропускаем цикл
			case i == 1 && j == 1:
			// Если элемент находится в первой строке рассмотрим только левого соседа
			case i == 1:
				if val < VertexMap[((i-1)*m+j)-1].Value {
					vertex.neighbor = append(vertex.neighbor, (i-1)*m+j-1)
				} else if val == VertexMap[((i-1)*m+j)-1].Value {
				} else {
					VertexBefore := VertexMap[((i-1)*m+j)-1]
					VertexBefore.neighbor = append(VertexBefore.neighbor, (i-1)*m+j)
					VertexMap[((i-1)*m+j)-1] = VertexBefore
				}
			// Если элемент находится в первом столбце рассмотрим только верхних соседей
			case j == 1:
				if val < VertexMap[((i-1)*m+j)-m].Value {
					vertex.neighbor = append(vertex.neighbor, ((i-1)*m+j)-m)
				} else if val == VertexMap[((i-1)*m+j)-m].Value {
				} else {
					VertexBefore := VertexMap[((i-1)*m+j)-m]
					VertexBefore.neighbor = append(VertexBefore.neighbor, (i-1)*m+j)
					VertexMap[((i-1)*m+j)-m] = VertexBefore
				}
			// Иначе вершина сравнивается с верхней и левой
			default:
				if val < VertexMap[((i-1)*m+j)-1].Value {
					vertex.neighbor = append(vertex.neighbor, (i-1)*m+j-1)
				} else if val == VertexMap[((i-1)*m+j)-1].Value {
				} else {
					// Если значение больше, чем у предыдущей вершины, добавим ей в соседи эту вершину
					VertexBefore := VertexMap[((i-1)*m+j)-1]
					VertexBefore.neighbor = append(VertexBefore.neighbor, (i-1)*m+j)
					VertexMap[((i-1)*m+j)-1] = VertexBefore
				}
				if val < VertexMap[((i-1)*m+j)-m].Value {
					vertex.neighbor = append(vertex.neighbor, ((i-1)*m+j)-m)
				} else if val == VertexMap[((i-1)*m+j)-m].Value {
				} else {
					VertexBefore := VertexMap[((i-1)*m+j)-m]
					VertexBefore.neighbor = append(VertexBefore.neighbor, (i-1)*m+j)
					VertexMap[((i-1)*m+j)-m] = VertexBefore
				}
			}

			// Добавим вершину в мапу
			VertexMap[(i-1)*m+j] = vertex

		}
	}
	//fmt.Println(VertexMap)

	// ИЩем наибольший путь методомо поиска в глубину
	// перебираем каждую вершину у которой verified = false
	// в процессе поиска, исследованным вершинам устанавливаем verified = true, т.к. нет смысла с нее начинать поиск
	// обрабатываем соседей у которых значение Value больше чем у предшествующей вершины

	// Создадим переменную для хранения максимального пути
	var MaxLenght int32

	// Для каждой вершины применяем поиск в ширину
	for NumVertex, VertexByKey := range VertexMap {
		//fmt.Println("Num vert: ", NumVertex)
		//var MaxVertLable int32

		// Если вершина помечена как пройденая, то ее пропускаем, т.к. она уже учавствовала в пути
		if VertexByKey.LableStart {
			continue
		}

		// Помечаем вершину как уже посчитанную в предыдущих разборах ребер
		VertexByKey.LableStart = true
		VertexMap[NumVertex] = VertexByKey

		VertexByKey.Count = 1

		// Создадим очередь и добавим текущую вершину
		S := list.New()
		S.PushFront(VertexByKey)

		// Если очередь не пуста
		for S.Len() != 0 {

			//выталкиваем вершину из очереди
			VertexNext := S.Remove(S.Back()).(Vertex)

			// Рассмотрим всех соседей
			for _, neighbors := range VertexNext.neighbor {

				VertexNeighbor := VertexMap[neighbors]

				//fmt.Println("Vert count", VertexNeighbor.Count)

				/*for _, numVert := range SliceVertex {
					if neighbors == numVert {
						VertexNeighbor.Count = VertexNext.Count + 1
						fmt.Println("вновь берем вершину ", VertexNeighbor.Count)
						VertexMap[neighbors] = VertexNeighbor

						// И добавляем в очередь
						S.PushFront(VertexNeighbor)
						continue
					}
				}

				// Проверяем, была ли пройдена вершина в предыдущих циклах
				if VertexNeighbor.LableStart {
					if MaxVertLable < VertexNeighbor.Count+VertexNext.Count {
						MaxVertLable = VertexNeighbor.Count + VertexNext.Count
					}
					continue
				}*/

				// Прибавляем счетчик на один от родительского
				VertexNeighbor.Count = VertexNext.Count + 1

				VertexNeighbor.LableStart = true
				VertexMap[neighbors] = VertexNeighbor

				//SliceVertex = append(SliceVertex, neighbors)

				// И добавляем в очередь
				S.PushFront(VertexNeighbor)

			}

			// Если очередь пуста, значит мы дошли до последнего элемента
			if S.Len() == 0 {
				// Если максимальный путь, вычисленный ранним циклом, меньше, чем путь, найденный в этой итерации
				/*if MaxVertLable < VertexNext.Count {
					// то первоначальной вершине присвоим больший путь
					VertexByKey.Count = VertexNext.Count
				} else {
					VertexByKey.Count = MaxVertLable
				}
				VertexMap[NumVertex] = VertexByKey

				fmt.Println(MaxLenght, VertexByKey.Count)*/

				// если наибольший путь меньше, чем у это вершины, значит это то что мы ищем
				if MaxLenght < VertexNext.Count {
					MaxLenght = VertexNext.Count
				}

				//fmt.Println("max len", MaxLenght)
			}

		}

	}

	fmt.Println(MaxLenght)
	//fmt.Println(VertexMap)
}
