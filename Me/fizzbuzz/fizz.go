package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis 30 aus und ersetzt dabei
// jede durch 3 teilbare Zahl durch "fizz" und jede durch
// 5 teilbare Zahl durch "buzz".
// Bei Zahlen, die sowohl durch 3 als auch durch 5 teilbar sind,
// wird "fizzbuzz" azsgegeben.
func Fizz() {
	for i := 1; i <= 30; i++ {
		// wenn i durch 3 und 5 teilbar ist, gib "fizzbuzz" aus
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")
		} else
		// wenn i durch 3 teilbar ist, gib "fizz" aus
		if i%3 == 0 {
			//fmt.Println(i, "fuzz" )

			fmt.Println("fizz")
		} else
		// wenn i durch 5 teilbar ist, gib "buzz" aus
		if i%5 == 0 {
			fmt.Println("buzz")

			// sonst gib i aus
		} else {
			fmt.Println(i)
		}
	}
}
