package main

import (
	"fmt"
)

func main() {

	fmt.Println("Движение лифта")

	coordinate := 1
	lift := 1
	var passenger int
	signal4 := true
	signal7 := true
	signal10 := true

	for {

		fmt.Println("Этаж:", coordinate, "   Пассажиров:", passenger)

		if (signal4 == false && signal7 == false && signal10 == false) && coordinate == 1 {
			break
		}

		coordinate = lift + coordinate

		if coordinate == 24 {
			lift = -1

		} else if coordinate == 1 {
			lift = 1
			passenger = 0
		}

		if coordinate == 10 && lift == -1 && passenger != 2 && signal10 {
			passenger++
			signal10 = false
		}

		if coordinate == 7 && lift == -1 && passenger != 2 && signal7 {
			passenger++
			signal7 = false
		}

		if coordinate == 4 && lift == -1 && passenger != 2 && signal4 {
			passenger++
			signal4 = false
		}
	}
}
