package main

import "fmt"

func main() {
	val := 0
	for {
		fmt.Print("Enter number:")
		fmt.Scanf("%d", val)

		if val != 0 && val <= 2 {
			fmt.Println("Enter the number greater than 2")
		} else if val > 0 {
			fmt.Println(fibonacci(val))
			break
		}
	}
}

func fibonacci(number int) []int {
	output := []int{1}
	sum := 1
	for i := 1; i < number; i++ {
		output = append(output, sum)
		sum = output[i] + output[i-1]
	}

	return output
}

func fibonacci2(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1

	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums
}
