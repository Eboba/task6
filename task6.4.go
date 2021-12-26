package main

import (
	"fmt"
)

func main() {

	fmt.Println("Предыдущее занятие на if") // сдается мне там опечатка в названии задания ;)
	fmt.Println("----------")

	var one int
	var two int
	var three int

	for {
		if one <= 10 {
			fmt.Println(one)
			one++

		} else if two <= 100 {
			fmt.Println(two)
			two++

		} else if three <= 1000 {
			fmt.Println(three)
			three++

		} else {
			break
		}
	}
}
