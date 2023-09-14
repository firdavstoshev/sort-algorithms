package main

import (
	"fmt"
	"math/rand"
	"time"
)

const minMerge = 32 // Минимальный размер для слияния

// Вспомогательная функция для сортировки вставками
func insertionSort(arr []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Вспомогательная функция для слияния двух подпоследовательностей
func merge(arr []int, left, middle, right int) {
	len1 := middle - left + 1
	len2 := right - middle

	// Создаем временные срезы для левой и правой подпоследовательностей
	leftArr := make([]int, len1)
	rightArr := make([]int, len2)

	// Копируем данные во временные срезы
	for i := 0; i < len1; i++ {
		leftArr[i] = arr[left+i]
	}
	for j := 0; j < len2; j++ {
		rightArr[j] = arr[middle+1+j]
	}

	i := 0
	j := 0
	k := left

	// Слияние подпоследовательностей обратно в исходный массив
	for i < len1 && j < len2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	// Копирование оставшихся элементов из левой подпоследовательности (если есть)
	for i < len1 {
		arr[k] = leftArr[i]
		i++
		k++
	}

	// Копирование оставшихся элементов из правой подпоследовательности (если есть)
	for j < len2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// Основная функция Tim Sort
func timSort(arr []int) {
	n := len(arr)

	// Сортировка каждой подпоследовательности длиной minMerge
	for i := 0; i < n; i += minMerge {
		left := i
		right := i + minMerge - 1
		if right >= n {
			right = n - 1
		}
		insertionSort(arr, left, right)
	}

	// Слияние подпоследовательностей
	for size := minMerge; size < n; size = 2 * size {
		for left := 0; left < n; left += 2 * size {
			middle := left + size - 1
			right := left + 2*size - 1

			if middle >= n {
				middle = n - 1
			}
			if right >= n {
				right = n - 1
			}

			if middle < right {
				merge(arr, left, middle, right)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Создаем случайный неотсортированный массив
	arr := make([]int, 20)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("Неотсортированный массив:", arr)
	timSort(arr)
	fmt.Println("Отсортированный массив:", arr)
}
