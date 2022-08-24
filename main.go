package main

import "fmt"

func main() {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	fmt.Println(arraySign(nums))
}

func arraySign(nums []int) int {
	product := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		} else if nums[i] < 0 {
			product = product * (-1)
		} else {
			product = product
		}

	}
	return product
}
