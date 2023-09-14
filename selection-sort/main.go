package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr) // Получаем длину массива.

	// Внешний цикл: проходим по всем элементам массива.
	for i := 0; i < n-1; i++ {
		minIndex := i // Предполагаем, что текущий элемент - минимальный.

		// Внутренний цикл: сравниваем текущий элемент с остальными элементами.
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j // Если найден элемент меньше текущего минимального, обновляем индекс минимального элемента.
			}
		}

		// После завершения внутреннего цикла меняем местами текущий элемент и минимальный элемент.
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	// Создаем неотсортированный массив.
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Неотсортированный массив:", arr)

	// Вызываем функцию сортировки выбором.
	selectionSort(arr)

	// Выводим отсортированный массив.
	fmt.Println("Отсортированный массив:", arr)
}
