package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func calPoints(operations []string) int {
	sum := 0

	score := []int{}

	for idx, op := range operations {
		fmt.Println(idx, op)
		switch op {
		case "+":
			num1, num2 := score[len(score)-1], score[len(score)-2]
			score = append(score, num1+num2)
		case "C":
			score = score[:len(score)-1]
		case "D":
			fmt.Println("Length of stack: ", len(score), len(score)-1)
			score = append(score, score[len(score)-1]*2)
		default:
			num, _ := strconv.Atoi(op)
			score = append(score, num)
		}
	}

	for _, num := range score {
		sum += num
	}
	return sum
}

// // Convert the string to an integer if it's a valid number
// if num, err := strconv.Atoi(char); err == nil {
// 	fmt.Println(char, "is a number")
// 	append(score, num)
// } else {
// 	fmt.Println(char, "is not a number")
// }

// // Handling the character "C"
// if char == "C" {
// 	fmt.Println("Invalidate previous score: ", sum)
// 	if num, err := strconv.Atoi(operations[idx-1]); err == nil {
// 		sum -= num
// 	} else {
// 		fmt.Println(char, "is not a number")
// 	}
// 	fmt.Println("New score: ", sum)
// }
