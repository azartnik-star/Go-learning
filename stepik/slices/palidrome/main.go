// Дан массив, состоящий из целых чисел. Напишите программу, которая определяет является ли массив палиндромом.
// То есть если перевернуть массив, то получится массив, равный первоначальному.
package main

import (
	f "fmt"
)

func main() {
	var size int
	f.Scan(&size)
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		f.Scan(&numbers[i])
	}
	isPalidrom := true
	for i := 0; i < size/2; i++ {
		if numbers[i] != numbers[size-i-1] {
			isPalidrom = false
			break
		}
	}
	if isPalidrom == true {
		f.Println("YES")
	} else {
		f.Println("NO")
	}
}
