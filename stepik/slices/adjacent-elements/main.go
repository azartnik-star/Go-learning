// напишите программу, которая переставляет соседние элементы массива
// (1-й элемент поменять с 2-м, 3-й с 4-м, и т.д. если элементов нечетное, то последний элемент остается на своем месте)

package main

import f "fmt"

func main() {
	var size int
	f.Scan(&size)
	numbers := make([]int, size)
	for i := 0; i < len(numbers); i++ {
		f.Scan(&numbers[i])
	}
	for z := 0; z+1 < len(numbers); z += 2 {
		numbers[z], numbers[z+1] = numbers[z+1], numbers[z]
	}
	for x := 0; x < len(numbers); x++ {
		f.Print(numbers[x], " ")
	}
}
