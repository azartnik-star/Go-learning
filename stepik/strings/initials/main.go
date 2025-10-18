// По данным ФИО сформируйте строку, содержащую фамилию с инициалами
// Например, по 'Ivanov Ivan Ivanovich' получите 'Ivanov I.I.'
package main

import f "fmt"

func main() {
	var surname, firstname, middlename string
	f.Scan(&surname, &firstname, &middlename)
	firstnameR, middlenameR := []rune(firstname), []rune(middlename)
	f.Print(surname, " ", string(firstnameR[0]), ".", string(middlenameR[0]), ".")
}
