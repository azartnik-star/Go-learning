// Определить, является ли введенное слово индентификатором.
// Идентификатор должен начинается с английской буквы (в любом регистре) или знака подчеркивая,
// а так же введенное слово НЕ должно содержать других символов, КРОМЕ букв английского алфавита (в любом регистре), цифр и знака подчеркивая.
package main

import f "fmt"

func main() {
	var word string
	flag := true
	firstflag := false
	f.Scan(&word)

	if (word[0] >= 65 && word[0] <= 90) || (word[0] == 95) || (word[0] >= 97 && word[0] <= 120) {
		firstflag = true
	}

	for i := 1; i < len(word); i++ {
		if !((word[i] >= 48 && word[i] <= 57) || (word[i] >= 65 && word[i] <= 90) || (word[i] == 95) || (word[i] >= 97 && word[i] <= 122)) {
			flag = false
			break
		}
	}
	if flag && firstflag {
		f.Print("YES")
	} else {
		f.Print("NO")
	}
}
