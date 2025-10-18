// Есть два дерева и нужно проверить одинаковые ли они
package main

import (
	f "fmt"

	"golang.org/x/tour/tree"
)

// Walk Совершает обход дерева и посылает значение в каналы
func Walk(t *tree.Tree, ch chan int) {
	// Рекурсивный обход дерева
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)

}

// Same проверяет схожесть двух деревьев
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// если один закрывался, а другой нет/если значения не совпадают, соответственно
		if ok1 != ok2 || v1 != v2 {
			return false
		} else if !ok1 && !ok2 {
			return true
			// если деревья одинаковы и заканчиваются в одно время
		}
	}
}

func main() {
	// Тест функции Walk
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()
	for v := range ch {
		f.Println(v)
	}
	f.Println(Same(tree.New(1), tree.New(1)))
	f.Println(Same(tree.New(1), tree.New(2)))
}
