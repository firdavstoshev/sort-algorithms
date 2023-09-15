package main

import "fmt"

func insertionSort(arr []int) {
	length := len(arr)

	for i := 1; i < length; i++ {
		// Запоминаем текущий элемент для вставки
		current := arr[i]
		j := i - 1

		// Сдвигаем элементы в отсортированной части массива, чтобы сделать место для текущего элемента
		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}

		// Вставляем текущий элемент на правильное место
		arr[j+1] = current
	}
}

func main() {
	arr := []int{5, 2, 9, 3, 6}
	fmt.Println("Исходный массив:", arr)

	insertionSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
