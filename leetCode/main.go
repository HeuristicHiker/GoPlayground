package main

import (
	"fmt"
	"sort"
)

/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.


*/

func main() {
	// plusOne([]int{9})
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
}

func plusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i >= 0 && carry > 0; i-- {
		newDigit := (digits[i] + carry) % 10
		carry = (digits[i] + carry) / 10
		digits[i] = newDigit
		fmt.Println("i: ", digits[i], " newDigit: ", newDigit, " carry: ", carry)
	}
	if carry > 0 {
		// Could just make a new array with len + 1 and set [0] to 1 but meh
		digits = append([]int{carry}, digits...)
	}
	return digits
}

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	dif := arr[1] - arr[0]

	for i := 1; i < len(arr); i++ {
		if dif != arr[i]-arr[i-1] {
			return false
		}
	}
	return true
}
