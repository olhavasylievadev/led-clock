package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {

	screen.Clear()
	for {
		screen.MoveTopLeft()
		now := time.Now()
		currentHour, currentMinute, currentSecond := now.Hour(), now.Minute(), now.Second()
		sSecond := now.Nanosecond() / 1e8

		fmt.Printf("hour: %d, min: %d, sec: %d\n", currentHour, currentMinute, currentSecond)

		clock := [...]Placeholder{
			digits[currentHour/10], digits[currentHour%10],
			colon,
			digits[currentMinute/10], digits[currentMinute%10],
			colon,
			digits[currentSecond/10], digits[currentSecond%10],
			dot,
			digits[sSecond],
		}

		//alarmed := currentSecond%10 == 0

		for line := range clock[0] {
			//if alarmed {
			//	clock = alarm
			//}
			// Print a line for each Placeholder in digits
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && currentSecond%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		const splitSecond = time.Second / 10
		time.Sleep(splitSecond)
	}
}
