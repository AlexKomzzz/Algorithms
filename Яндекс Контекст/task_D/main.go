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
	verified   bool
	LableStart bool
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
			//Если это верхний правый угол - добавляем соседа слева
			case i == 1 && j == m:
				//Если у вершины значение меньше чем у соседа слева, то добавим этого соседа, иначе предыдущей вершине добывим в соседи эту вершину
				if val < VertexMap[((i-1)*m+j)-1].Value {
					vertex.neighbor = append(vertex.neighbor, (i-1)*m+j-1)
				} else {
					// Если значение больше, чем у предыдущей вершины, добавим ей в соседи эту вершину
					VertexBefore := VertexMap[((i-1)*m+j)-1]
					VertexBefore.neighbor = append(VertexBefore.neighbor, (i-1)*m+j)
					VertexMap[((i-1)*m+j)-1] = VertexBefore
				}
			//Если это нижний левый угол, проверяем верхего соседа. Если наше значение меньше, то добавляем его в соседи, если больше то к нему добавляем нашу вершину
			case i == n && j == 1:
				if val < VertexMap[((i-1)*m+j)-m].Value {
					vertex.neighbor = append(vertex.neighbor, ((i-1)*m+j)-m)
				} else {
					VertexBefore := VertexMap[((i-1)*m+j)-m]
					VertexBefore.neighbor = append(VertexBefore.neighbor, (i-1)*m+j)
					VertexMap[((i-1)*m+j)-m] = VertexBefore
				}
			// Если элемент находится в первой строке рассмотрим только левого соседа
			case i == 1:
				if val < VertexMap[((i-1)*m+j)-1].Value {
					vertex.neighbor = append(vertex.neighbor, (i-1)*m+j-1)
				} else {
					VertexBefore := VertexMap[((i-1)*m+j)-1]
					VertexBefore.neighbor = append(VertexBefore.neighbor, (i-1)*m+j)
					VertexMap[((i-1)*m+j)-1] = VertexBefore
				}
			// Если элемент находится в первом столбце рассмотрим только верхних соседей
			case j == 1:
				if val < VertexMap[((i-1)*m+j)-m].Value {
					vertex.neighbor = append(vertex.neighbor, ((i-1)*m+j)-m)
				} else {
					VertexBefore := VertexMap[((i-1)*m+j)-m]
					VertexBefore.neighbor = append(VertexBefore.neighbor, (i-1)*m+j)
					VertexMap[((i-1)*m+j)-m] = VertexBefore
				}
			// Иначе вершина сравнивается с верхней и левой
			default:
				if val < VertexMap[((i-1)*m+j)-1].Value {
					vertex.neighbor = append(vertex.neighbor, (i-1)*m+j-1)
				} else {
					// Если значение больше, чем у предыдущей вершины, добавим ей в соседи эту вершину
					VertexBefore := VertexMap[((i-1)*m+j)-1]
					VertexBefore.neighbor = append(VertexBefore.neighbor, (i-1)*m+j)
					VertexMap[((i-1)*m+j)-1] = VertexBefore
				}
				if val < VertexMap[((i-1)*m+j)-m].Value {
					vertex.neighbor = append(vertex.neighbor, ((i-1)*m+j)-m)
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

	// Для каждой вершины применяем поиск в глубину
	for i := int32(1); i <= n; i++ {
		for j := int32(1); j <= m; j++ {

			// Если вершина помечена как пройденая, то ее пропускаем, т.к. она уже учавствовала в пути
			if VertexMap[(i-1)*m+j].LableStart {
				continue
			}

			// Счетчик для счета пути
			Count := int32(-1)

			Lable := true
			// при проходе вершины задать  verified = Lable, LableStart = true\

			VertexNow := VertexMap[(i-1)*m+j]
			VertexNow.LableStart = true
			VertexMap[(i-1)*m+j] = VertexNow

			// Создадим стек и добавим текущую вершину
			S := list.New()
			S.PushBack(VertexNow)

			// Если стек не пуст
			for S.Len() != 0 {
				Count++
				//выталкиваем вершину из стека
				VertexNext := S.Back().Value.(Vertex)

				// Если вершина не изведана проходим по ее соседям
				if !VertexNext.verified {
					for _, neighbors := range VertexNext.neighbor {
						VertexNeighbor := VertexMap[neighbors]
						VertexNeighbor.LableStart = true
						VertexMap[neighbors] = VertexNeighbor
						S.PushBack(VertexNeighbor)
					}
					VertexNext.verified = Lable
				}

			}

			if MaxLenght < Count {
				MaxLenght = Count
			}

		}
	}

	fmt.Println(MaxLenght)
}
