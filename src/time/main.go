package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now()
	fmt.Println(date.Format("2006-01-02 15:04:05"))
	var t = date.Add(time.Minute * 5).Sub(date)
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())
	fmt.Println(t.Minutes())
}
