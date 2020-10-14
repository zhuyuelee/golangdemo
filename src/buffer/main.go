package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	var buffer bytes.Buffer
	index := 0
	for {
		buffer.WriteString(getNextString())
		if index == 20 {
			break
		}
		index++
	}
	fmt.Print(buffer.String(), "\n")
}

func getNextString() string {
	date := time.Now()
	return fmt.Sprintf("%d%d%d%d%d\n", date.Year(), date.Month(), date.Day(), date.Hour(), date.Nanosecond())
}
