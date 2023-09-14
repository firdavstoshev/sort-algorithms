package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	// Разделение на элементы меньше и больше опорного
	var less, greater []int
	for i, v := range arr {
		if i == pivotIndex {
			continue
		}
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	// Рекурсивная сортировка левой и правой частей
	less = quickSort(less)
	greater = quickSort(greater)

	// Объединение отсортированных частей и опорного элемента
	sorted := append(less, pivot)
	sorted = append(sorted, greater...)

	return sorted
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Создаем случайный неотсортированный массив
	arr := make([]int, 20)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("Неотсортированный массив:", arr)

	// Применяем Quick Sort
	arr = quickSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
