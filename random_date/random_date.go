package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)

	year := 1990 + rand.Intn(30)
	month := rand.Intn(12) + 1
	day := rand.Intn(30) + 1

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	fmt.Println(date.Format("2006年1月2日"))
}
