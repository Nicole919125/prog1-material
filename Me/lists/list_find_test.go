package lists

import "fmt"

// Sucht x in l und liefert die Position
// Falls x nicht in l vorkommt, wird -1 geliefert.
func Find(l []int, x int) int {
	// TODO
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}

	return -1
}
func ExampleLists() {
	l1 := []int{17, 5, 42, 25, 3, -4, 8, -23}

	pos1 := Find(l1, 42)
	pos2 := Find(l1, 200)

	fmt.Println(pos1)
	fmt.Println(pos2)

	// Output:
	// 2
	// -1
}
