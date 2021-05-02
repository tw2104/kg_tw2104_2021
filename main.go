package main

import (
	"fmt"
	"os"
)

func main() {

    // get termial agrs input
	argsWithoutProg := os.Args[1:]

	// test print
	fmt.Println(argsWithoutProg);
}



