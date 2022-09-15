package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dice, sides := 2, 12
	rolls := 2

	for r := 1; r <= rolls; r++ {
		sum := 0

		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("roll #", r, "die #", d, ":", rolled)
		}
		fmt.Println("total rolled :", sum)

		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("Snake eyes!")
			break
		case sum == 7:
			fmt.Println("Lucky 7!")
			break
		case sum%2 == 0:
			fmt.Println("Even!")
			break
		case sum%2 != 0:
			fmt.Println("Odd!")
		}
	}
}

func roll(sides int) int {
	return rand.Intn(sides) + 1
}
