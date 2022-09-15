package main

import "fmt"

func main() {
	// loops
	// for
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += 1
	}

	// go version of while
	sum2 := 1
	for sum2 < 1000 {
		sum2 += 1
	}

	// exercise
	// print numbers from 1 to 50
	// if number is divisible by 3 print "Fizz"
	// if number is divisible by 5 print "Buzz"
	// if number is divisible by 3 and 5 print "FizzBuzz"

	for i := 1; i <= 50; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
