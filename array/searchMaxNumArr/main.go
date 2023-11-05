// Поиск наибольшего числа в массиве стандартными средствами
package main

import (
	"errors"
	"fmt"
)

// searchMax - ищет наибольшее число в массиве целых чисел
func searchMax(arr []int) int {
	if len(arr) < 1 {
		errors.New("Массив пустой!")
	}
	maxNum := arr[0]
	for i := 0; i < len(arr); i++ {
		maxNum = max(maxNum, arr[i])
	}
	return maxNum
}

func main() {
	nums := []int{4, 9, 2, 3, 8, 22, 18}
	fmt.Printf("Наибольшее число в последовательности: %d\n", searchMax(nums))
}
