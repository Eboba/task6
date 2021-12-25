package main

import (
	"fmt"
)

func main() {

	fmt.Println("Написание простого цикла")
	fmt.Println("----------------")

	fmt.Println("Введите число:")
	var num int
	fmt.Scan(&num)

	for num <= 0 {
		fmt.Println("Введите число, больше нуля:")
		fmt.Scan(&num)
	}

	for sum := 0; sum <= num; sum = sum + 1 {
		fmt.Println(sum)
	}
}
