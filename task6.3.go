package main

import (
	"fmt"
)

func main() {

	fmt.Println("Расчёт суммы скидки")
	fmt.Println("----------")

	fmt.Println("Введите сумму товара:")
	var sum int
	fmt.Scan(&sum)

	fmt.Println("Введите скидку:")
	var dis int
	fmt.Scan(&dis)

	realDis := sum * 30 / 100

	for {
		if sum <= 0 {
			fmt.Println("Сумма должна быть больше 0!")
			fmt.Println("Введите сумму товара:")
			fmt.Scan(&sum)
			continue

		} else if dis < 0 {
			fmt.Println("Скидка должна быть не меньше 0!")
			fmt.Println("Введите скидку:")
			fmt.Scan(&dis)
			continue

		} else if sum <= dis {
			fmt.Println("Скидка должна быть меньше суммы товара!")
			fmt.Println("Введите скидку:")
			fmt.Scan(&dis)
			continue
		} else {
			break
		}
	}

	if realDis > 2000 {
		realDis = 2000
	}

	if realDis > dis {
		realDis = dis
	}

	fmt.Println("Ваша скидка", realDis, "руб.")
}
