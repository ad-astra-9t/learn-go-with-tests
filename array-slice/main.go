package main

import "fmt"

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(all ...[]int) []int {
	sums := make([]int, 0)
	for _, nums := range all {
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumAllTails(all ...[]int) []int {
	sums := make([]int, 0)
	for _, nums := range all {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(nums[1:]))
		}
	}
	return sums
}

func main() {
	fmt.Println("Hello world!")
}
