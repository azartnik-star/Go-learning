// Дано число n, n<=100. Создайте массив n x n и заполните его по следующему правилу:
// числа на диагонали, идущей из правого верхнего в левый нижний угол, равны 1;
// числа, стоящие выще этой диагонали, равны 0
// числа, стоящие ниже этой диагонали, равны 2
package main

import f "fmt"

func main() {
	var rowsCount int
	f.Scan(&rowsCount)
	columsCount := rowsCount
	numbers := make([][]int, rowsCount)
	for i := 0; i < rowsCount; i++ {
		numbers[i] = make([]int, columsCount)
		for x := 0; x < columsCount; x++ {
			if i+x == rowsCount-1 {
				numbers[i][x] = 1
			} else if i+x < rowsCount-1 {
				numbers[i][x] = 0
			} else {
				numbers[i][x] = 2
			}
		}
	}
	for x := 0; x < rowsCount; x++ {
		for j := 0; j < columsCount; j++ {
			f.Print(numbers[x][j], " ")
		}
		f.Println()
	}
}
