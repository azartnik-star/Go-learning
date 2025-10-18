// По данной строке определите, является ли она палидромом. То есть, которое одинаково читается слева направо и справа налево
package main

import f "fmt"

func main() {
	var word string
	flag := true
	f.Scan(&word)
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			flag = false
			break
		}
	}
	if flag {
		f.Print("YES")
	} else {
		f.Print("NO")
	}
}
