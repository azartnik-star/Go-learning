package main

import f "fmt"

// Fact2 вычисляет двойной факториал заданного числа
func Fact2(n int) int {
	fac := 1
	if n == 0 || n == 1 {
		return 1
	} else {
		for i := n; i > 1; i++ {
			fac *= i
		}
	}
	return fac
}

func main() {
	var a, b, c int
	f.Scan(&a, &b, &c)
	x := Fact2(a)
	z := Fact2(b)
	y := Fact2(c)
	f.Print(x, z, y)
}
