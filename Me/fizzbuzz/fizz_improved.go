package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis n aus und ersetzt dabei
// jede durch x teilbare Zahl durch "fizz" und jede durch
// y teilbare Zahl durch "buzz".
// Bei Zahlen, die sowohl durch 3 als auch durch y teilbar sind,
// wird "fizzbuzz" azsgegeben.
func FizzImproved(x, y, n int) {
	for i := 1; i <= n; i++ {
		// wenn i durch 3 und 5 teilbar ist, gib "fizzbuzz" aus
		if i%y == 0 && i%x == 0 {
			fmt.Println("fizzbuzz")
		} else
		// wenn i durch 3 teilbar ist, gib "fizz" aus
		if i%x == 0 {
			//fmt.Println(i, "fuzz" )

			fmt.Println("fizz")
		} else
		// wenn i durch 5 teilbar ist, gib "buzz" aus
		if i%y == 0 {
			fmt.Println("buzz")

			// sonst gib i aus
		} else {
			fmt.Println(i)
		}
	}
}
