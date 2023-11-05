// Ряд Фибоначчи. Несколько реализаций алгоритма
package main

import (
	"fmt"
)

// fibOneVersion - первая версия вычисления ряда Фибоначчи
func fibOneVersion(position int) int {
	seq := []int{0, 1}

	for len(seq) <= position {
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
	}
	return seq[len(seq)-1]
}

func main() {
	fmt.Println("Старт программы для вычисления ряда Фибоначчи")

	var numFib int
	fmt.Println("Введите целое число для вычисления ряда Фибоначчи:")
	fmt.Scan(&numFib)

	fmt.Printf("Версия алгоритма 1. Число Фибоначчи: %d\n", fibOneVersion(numFib))
}
