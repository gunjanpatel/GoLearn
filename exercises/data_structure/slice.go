package main

import "fmt"

func main() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("First Month is", months[0])
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	quarter1 := months[0:3]
	quarter2 := months[3:6]

	// Slice extended to one more elm
	quarter2Extended := quarter2[:4]

	// Slice reduced to start from 1
	quarter2Extended2 := quarter2[1:]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))
	fmt.Println(quarter2Extended2, len(quarter2Extended2), cap(quarter2Extended2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

	// Append slice
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)

		// Notice cap 4 and then 8 and then 16, why?
		// When a slice doesn't have enough capacity to hold more elements, Go doubles its capacity.
		// 0       cap=1   [0]
		// 1       cap=2   [0 1]
		// 2       cap=4   [0 1 2]
		// 3       cap=4   [0 1 2 3]
		// 4       cap=8   [0 1 2 3 4]
		// 5       cap=8   [0 1 2 3 4 5]
		// 6       cap=8   [0 1 2 3 4 5 6]
		// 7       cap=8   [0 1 2 3 4 5 6 7]
		// 8       cap=16  [0 1 2 3 4 5 6 7 8]
		// 9       cap=16  [0 1 2 3 4 5 6 7 8 9]
	}

	letters := []string{"A", "B", "C", "D", "E"}
	unset(letters, 2)
}

// GO doesn't have native function to remove element from slice. Instead we can use s[i:p] method for creating new slice.
func unset(data []string, position int) {
	if position < len(data) {

		fmt.Println("Before", data, "Remove ", data[position])

		data = append(data[:position], data[position+1:]...)

		fmt.Println("After", data)
	}
}
