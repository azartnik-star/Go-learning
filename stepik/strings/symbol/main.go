// Дана строка; Удалите k-ый сивол в ней
package main

import f "fmt"

func main() {
	var word, newWord string
	f.Scan(&word)
	number := 0
	f.Scan(&number)
	number -= 1
	for i := 0; i < len(word); i++ {
		if number != i {
			newWord += string(word[i])
		}
	}
	f.Print(newWord)
}
