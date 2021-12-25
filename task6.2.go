package main

import (
	"fmt"
)

func main() {

	fmt.Println("Сумма двух чисел по единице")
	fmt.Println("----------------")

	fmt.Println("Введите первое число:")
	var numOne int
	fmt.Scan(&numOne)

	fmt.Println("Введите второе число:")
	var numTwo int
	fmt.Scan(&numTwo)

	for {
		if numOne <= numTwo {
			fmt.Println(numOne)
			numOne++
		} else {
			break
		}
	}
}
