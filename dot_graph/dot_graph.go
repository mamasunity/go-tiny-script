package main

import (
	"os"
	"strings"
)

func main() {
	f, err := os.Create("./output.dot")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := strings.ReplaceAll(os.Args[1], " ", ";")
	s = strings.ReplaceAll(s, ",", "--")
	s = "graph{" + s + "}"

	_, err = f.WriteString(s)
	if err != nil {
		panic(err)
	}
}
