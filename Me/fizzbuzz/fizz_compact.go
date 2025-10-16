package fizzbuzz

import "fmt"

func FizzCompact() {
	for i := 1; i <= 30; i++ {
		//Wenn i weder durch 3 noch durch 5 teilbar ist
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			//continue
		}
		if i%3 == 0 {
			fmt.Println("fizz")
		}
		if i%5 == 0 {
			fmt.Println("buzz")
		}
		fmt.Println()
	}
}
