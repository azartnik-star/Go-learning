// напишите программу, которая циклически сдвигает элементы массива вправо. Например, если элементы нумеруются,
// начиная с нуля, то 0-й элемент становится 1-м, 1-й становится 2-м, ..., последний становится 0-м, то есть массив
// [3,5,7,9] превращается в массив [9,3,5,7]
package main

import f "fmt"

func main() {
	var size int
	f.Scan(&size)
	numbers := make([]int, size)
	last := 0
	for i := 0; i < size; i++ {
		f.Scan(&numbers[i])
		last = numbers[len(numbers)-1]
	}
	for z := len(numbers) - 1; z >= 1; z-- {
		numbers[z] = numbers[z-1]
	}
	numbers[0] = last
	for x := 0; x < len(numbers); x++ {
		f.Print(numbers[x], " ")
	}
}
