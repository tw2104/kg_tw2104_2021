package main

import (
	"fmt"
	"os"
)

var (
	// map interges to their corresponding phonetic string
	byteToStr = map[byte]string {
        '0' : "Zero",
        '1' : "One",
        '2' : "Two",
        '3' : "Three",
        '4' : "Four",
        '5' : "Five",
        '6' : "Six",
        '7' : "Seven",
        '8' : "Eight",
        '9' : "Nine",
    }
)

func main() {

    // get termial agrs input
	argsWithoutProg := os.Args[1:]

	for i := 0; i < len(argsWithoutProg); i++ {
		// print comma as divider except for the first
		if (i != 0) {
			fmt.Print(",")
		}

		// print the current number as phonetic string
		printAsStr(argsWithoutProg[i])
	}

	// print a newline after printing all the results
	fmt.Println();
}

func printAsStr(num string) {

	//print the input number as phonetic string
	for i := 0; i < len(num); i++ {
		fmt.Print(byteToStr[num[i]])
	}
}