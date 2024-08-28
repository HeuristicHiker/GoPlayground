package main

import (
	"fmt"
	"strings"
)

/* I've got my key value pair
What about IV
	need to know if I have to subtract
Between 5 and 10 do I subtract?
	5 no - it's a key
	6 no - I add
	7 no - I add
	8 no - I add
	9 yes - I sub
Between 50 and 100?
	41 - XLI
	67 - LXVII
	71 - LXXI
	77 - LXXVII
	78 - LXXVIII
	79 - LXXIX
*/

func main() {
	// input := "MCMXCIV"
	// romanToInt(input)

}

func romanToInt(s string) int {
	total := 0
	largestValue := 0

	for _, numeral := range s {
		value := romanNumeralToInt(string(numeral))

		// is subtracting
		if value > largestValue {
			// subtract * 2 because it was already added
			total += value - 2*largestValue
		} else {
			total += value
		}

		largestValue = value

	}
	return total
}

func romanNumeralToInt(rn string) int {
	romanToInt := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	return romanToInt[rn]
}

type numInt struct {
	value   int
	numeral string
}

func intToRomanNumeral(i int) string {
	numInts := []numInt{
		{1000, "M"},
		{500, "D"},
		{100, "C"},
		{50, "L"},
		{10, "X"},
		{5, "V"},
		{1, "I"},
	}

	asString := ""

	/*
		imaginary loop

		element = 9
		Find 5V
		Try to add smal num
		At 3 tries, stop
			Add back
			Try large num
		See if we good

	*/

	for index, irn := range numInts[:6] {
		if i == 0 {
			break
		}
		if i == irn.value {
			asString += irn.numeral
			return asString
		}
		if i > irn.value {
			// give it a shot
			temp := asString
			for strikes := 3; strikes >= 0; strikes-- {
				if strikes == 3 {
					asString += irn.numeral
					i -= irn.value
				}
				if strikes == 0 {
					// need to sub
					i = -numInts[index-1].value
					asString = numInts[index+1].numeral + temp
				}
				lowerNum := numInts[index+1]
				additionalValue := strikes * lowerNum.value
				fmt.Println("strikes: ", strikes, " * ", lowerNum.value, " = ", additionalValue)

				if additionalValue == i {
					fmt.Println("i ", i, " = ", additionalValue)
					asString = asString + strings.Repeat(lowerNum.numeral, strikes)
					i -= strikes
					break
				}
				// if i+lowerNum.value < i {
				// 	i += lowerNum.value
				// 	asString += lowerNum.numeral
				// } else {
				// 	break
				// }
			}
			if i < 0 {
				asString += numInts[index-1].numeral
			}
			// asString = numInts[index+1].numeral + asString
		}
	}

	return asString
}

func digitToNumeral(i int) string {
	// numInts := []numInt{
	// 	{1, "I"},
	// 	{5, "V"},
	// 	{10, "X"},
	// 	{50, "L"},
	// 	{100, "C"},
	// 	{500, "D"},
	// 	{1000, "M"},
	// }

	// diff := 1

	// fmt.Println(i)

	// for _, ni := range numInts {
	// 	val := ni.value
	// 	s := ni.numeral

	// }

	return "Hi"

}
