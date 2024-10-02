package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time :", now)
	fmt.Println(now.Format("01-02-2006 Monday 03:04:05PM")) // for 12h

	formattedTime := now.Format("02-01-2006 Monday 15:04:05") // for 24h
    fmt.Println("Formatted time:", formattedTime)

	date := time.Date(2024, 9, 14, 0, 0, 0, 0, time.UTC)
    fmt.Println(date)

	year := date.Year()
    month := date.Month()
    day := date.Day()
    fmt.Println(year, month, day)
}