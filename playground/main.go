package main

import "fmt"

func main() {
	// testList := string[]{"One", "Two", "Three"}
	testString := "One"

	// fmt.Println(testString)
	// addressOfTest := &testString
	newAdr := replaceString(testString)
	// testString = "Two"
	// fmt.Println(&testString)
	fmt.Println("out fun ", *newAdr, &newAdr)
	fmt.Println("last fun ", testString, &testString)
}

func replaceString(strPtr string) *string {
	strPtr = "Two"
	fmt.Println("in fun: ", strPtr, &strPtr)
	return &strPtr
}
