package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	// Внешний цикл: проходим по всем элементам массива.
	for i := 0; i < n-1; i++ {
		swapped := false // Флаг для оптимизации - если на данной итерации не было перестановок, массив уже отсортирован.

		// Внутренний цикл: сравниваем и переставляем соседние элементы.
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Если текущий элемент больше следующего, меняем их местами.
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true // Устанавливаем флаг перестановки.
			}
		}

		// Если на данной итерации не было перестановок, массив уже отсортирован, можно завершить сортировку.
		if !swapped {
			break
		}
	}
}

func main() {
	// Создаем неотсортированный массив.
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Неотсортированный массив:", arr)

	// Вызываем функцию сортировки пузырьком.
	bubbleSort(arr)

	// Выводим отсортированный массив.
	fmt.Println("Отсортированный массив:", arr)
}
