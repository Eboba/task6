package main

import (
	"fmt"
)

func main() {

	fmt.Println("Задача на постепенное наполнение корзин разной ёмкости")

	var basketOne int
	var basketTwo int
	var basketThree int

	for {
		fmt.Println("Введите емкость 1 корзины:")
		fmt.Scan(&basketOne)

		fmt.Println("Введите емкость 2 корзины:")
		fmt.Scan(&basketTwo)

		fmt.Println("Введите емкость 3 корзины:")
		fmt.Scan(&basketThree)

		if basketOne <= 0 || basketTwo <= 0 || basketThree <= 0 {
			fmt.Println("Емкость корзин должна быть больше 0:")
		} else {
			break
		}
	}

	var appleBasketone int
	var appleBaskettwo int
	var appleBasketthree int

	for {

		if appleBasketone != basketOne {
			appleBasketone++
			fmt.Println("В первой корзине", appleBasketone, "из", basketOne)
			continue
		}

		if appleBaskettwo != basketTwo {
			appleBaskettwo++
			fmt.Println("Во второй корзине", appleBaskettwo, "из", basketTwo)
			continue
		}

		if appleBasketthree != basketThree {
			appleBasketthree++
			fmt.Println("В третьей корзине", appleBasketthree, "из", basketThree)
			continue
		}

		fmt.Println("Корзины заполнены! Потребовалось", appleBasketone+appleBaskettwo+appleBasketthree, "яблок.")
		break
	}
}
