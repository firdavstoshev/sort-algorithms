package main

import "fmt"

func shakerSort(arr []int) {
	n := len(arr)   // Получаем длину массива
	swapped := true // Флаг, указывающий, были ли сделаны обмены на текущей итерации

	for swapped {
		swapped = false

		// Проход слева направо
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				// Если текущий элемент больше следующего, меняем их местами
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		if !swapped {
			break // Если на этой итерации не было обменов, массив уже отсортирован
		}

		// Проход справа налево
		for i := n - 2; i >= 0; i-- {
			if arr[i] > arr[i+1] {
				// Если текущий элемент больше следующего, меняем их местами
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Неотсортированный массив:", arr)

	shakerSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
