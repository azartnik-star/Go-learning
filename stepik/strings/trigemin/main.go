// Сейчас в городе Арторзе популярна троичная система счисления. Для телеграфной передачи чисел, записанных в троичной системе счисления,
// используется азбука Артрозе. Цифра 0 передается как ., 1 как -., 2 как --
package main

import f "fmt"

func main() {
	var a string
	f.Scan(&a)
	sym := []rune(a)
	var newSym []rune
	for i := 0; i < len(sym)-1; {
		if sym[i] == '.' {
			newSym = append(newSym, '0')
			i++
		} else if sym[i] == '-' && sym[i+1] == '.' {
			newSym = append(newSym, '1')
			i += 2
		} else if sym[i] == '-' && sym[i+1] == '-' {
			newSym = append(newSym, '2')
			i += 2
		}
	}
	f.Print(string(newSym))
}
