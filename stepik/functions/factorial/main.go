// по данным наутральным n и k вычислите значение Cnk=n!/k!(n-k)! число сочетаний из n элементов по k

package main

import f "fmt"

func factor(n int) int {
	fac := 1
	for i := 1; i <= n; i++ {
		fac *= i
	}
	return fac
}

func main() {
	var a, b int
	f.Scan(&a, &b)
	facA := factor(a)
	facB := factor(b)
	facC := factor(a - b)
	f.Print(facA / (facB * facC))
}
