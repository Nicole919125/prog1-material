package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//my_number := 42
	my_number := rand.Intn(100) + 1

	for i := 0; i < 3; i++ {
		guess := ReadNumber()
		if guess < my_number {
			fmt.Println("Die gesuchte Zahl ist größer.")
		}
		if guess > my_number {
			fmt.Println("Die gesuchte Zahl ist kleiner.")
		}
		if guess == my_number {
			fmt.Println("Das war richtig!")
			return
		}

	}
	fmt.Println("Game Over!")
}

// ReadNumber liefert uns ein int
func ReadNumber() int {
	var n int
	fmt.Println("Bitte gib eine Zahl ein: ")
	fmt.Scan(&n)
	return n

}
