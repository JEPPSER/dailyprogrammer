package main

import (
	"fmt"
)

var print = fmt.Println;

func main() {
	print(yahtzee_upper([]int{2, 3, 5, 5, 6}))
	print(yahtzee_upper([]int{1, 1, 1, 1, 3}))
	print(yahtzee_upper([]int{1, 1, 1, 3, 3}))
	print(yahtzee_upper([]int{1, 2, 3, 4, 5}))
	print(yahtzee_upper([]int{6, 6, 6, 6, 6}))

	print(yahtzee_upper([]int{1654, 1654, 50995, 30864, 1654, 50995, 22747,
		1654, 1654, 1654, 1654, 1654, 30864, 4868, 1654, 4868, 1654,
		30864, 4868, 30864}))
}

func yahtzee_upper(values []int) int {
	var checked[] int

	max := 0

	for i := 0; i < len(values); i++ {
		if contains(checked, values[i]) { continue }
		sum := values[i]
		checked = append(checked, values[i])
		for j := 0; j < len(values); j++ {
			if j == i { continue }
			if values[j] == values[i] {
				sum += values[j]
			}
		}
		if sum > max { max = sum }
	}

	return max
}

func contains(slice []int, val int) bool {
	for _, a := range slice {
		if a == val {
			return true
		}
	}
	return false
}