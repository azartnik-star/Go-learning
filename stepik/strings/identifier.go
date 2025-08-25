// Определить, является ли введенное слово идентификатором.
// Идентификатор должен начинаться с английской буквы (в любом регистре) или знака подчёркивания,
// а так же введенное слово НЕ должно содержать других символов, КРОМЕ букв английского алфавита (в любом регистре), цифр и знака подчеркивания.
package main

import (
	f "fmt"
)

func main() {
	var word string
	flag := true
	firstFlag := false
	f.Scan(&word)

	if (word[0] >= 65 && word[0] <= 90) || (word[0] == 95) || (word[0] >= 97 && word[0] <= 122) {
		firstFlag = true
	}

	for i := 1; i < len(word); i++ {
		if !((word[i] >= 48 && word[i] <= 57) || (word[i] >= 65 && word[i] <= 90) || (word[i] == 95) || (word[i] >= 97 && word[i] <= 122)) {
			flag = false
			break
		}
	}
	if flag && firstFlag {
		f.Print("YES")
	} else {
		f.Print("NO")
	}
}
