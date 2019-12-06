package main

import (
	"fmt"
)

// check if a number is a password
func checkPassword(password int) bool {
	var check bool
	numbers := fmt.Sprint(password) // convert number to string

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
		if numbers[i] == numbers[i+1] {
			check = true
		}
	}

	return check
}

// check the amount of passwords in a range
func checkPasswords(min, max int, c chan int) {
	var amount int

	for i := min; i <= max; i++ {
		if checkPassword(i) {
			amount++
		}
	}

	c <- amount
}

func main() {
	var amount int
	inputMin := 236491
	inputMax := 713787
	c := make(chan int)

	// parallel version
	amount = 0
	go checkPasswords(inputMin, inputMax/2, c)
	go checkPasswords(inputMax/2, inputMax, c)
	amount += <-c
	amount += <-c
	fmt.Println(amount)
}
