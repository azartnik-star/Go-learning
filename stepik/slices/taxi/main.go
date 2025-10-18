// После окончания уроков n групп школьников вышли на улицу и собрались ехать домой к Артуру на празднование его дня рождения.
// Известно, что i-ая группа состоит из s друзей (1 <= si <= 4), которые не хотят расставаться по пути к артуру.
// Решено ехать на такси. Каждая машина может вместить не более четырех пассажиров.
// Какое минимальное количество машин потребуется школьникам. если каждая группа должна целиком находиться в одной из машин такси
// (но одна машина может вмещать более чем одну группу)
package main

import f "fmt"

func main() {
	var group, cnt4, cnt3, cnt2, cnt1, taxi, min3, min2 int
	f.Scan(&group)
	numbers := make([]int, group)
	for i := 0; i < group; i++ {
		f.Scan(&numbers[i])
		switch {
		case numbers[i] == 4:
			cnt4++
		case numbers[i] == 3:
			cnt3++
		case numbers[i] == 2:
			cnt2++
		case numbers[i] == 1:
			cnt1++
		}
	}
	taxi = cnt4
	if cnt3 < cnt1 {
		min3 = cnt3
	} else {
		min3 = cnt1
	}
	cnt3 -= min3
	cnt1 -= min3
	taxi += min3
	taxi += cnt3

	min2 = cnt2 / 2
	taxi += min2
	cnt2 %= 2

	if cnt2 == 1 {
		if cnt1 >= 2 {
			taxi++
			cnt1 -= 2
		} else {
			taxi++
			cnt1 = 0
		}
	}
	taxi += cnt1 / 4
	if cnt1%4 != 0 {
		taxi++
	}
	f.Println(taxi)
}
