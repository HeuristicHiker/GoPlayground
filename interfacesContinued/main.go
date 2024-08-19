package main

import (
	"io"
	"log"
	"os"
)

/*

	Read contents of text file
	Print contents

	File as argument to program when executed in terminal

	ex: go run main.go myfile.txt

	Look at
		os.Args
		os/#Open
		Interface for File type implement

		If File type implements reader interface
		Might be able to reuse io.Copy function

*/

type file interface {
	// be able to read in a file
	// be able to copy contents
	// be able to display those contents
	readFile() ([]byte, error)
	Write([]byte)
}

type sourceFile struct {
	path string
}

func main() {
	fileName := os.Args[1]
	// file[0] is the compiled temp executable file
	// file[1] is the file name

	// Read the file

	bs, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// s := strings.Split(string(bs), ",")
	// fmt.Println(s)

	io.Copy(os.Stdout, bs)

}

// Challenge: create your own write file
