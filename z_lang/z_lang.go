package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		if strings.Contains(os.Args[1], "z") {
			fmt.Println("yes")

			return
		}
	}

	fmt.Println("no")
}
