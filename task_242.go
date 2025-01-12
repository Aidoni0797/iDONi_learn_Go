package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m) // Считываем размеры массива

	// Создаём массив
	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
	}

	// Заполняем массив по правилам
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				A[i][j] = 1 // Первый ряд и первый столбец заполняем 1
			} else {
				A[i][j] = A[i-1][j] + A[i][j-1] // Остальные элементы — сумма сверху и слева
			}
		}
	}

	// Вывод массива
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(A[i][j], " ") // Выводим элементы строки через пробел
		}
		fmt.Println() // Переход на новую строку
	}
}
