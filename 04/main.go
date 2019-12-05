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

func main() {
	var passwords int
	inputMin := 236491
	inputMax := 713787

	for i := inputMin; i <= inputMax; i++ {
		if checkPassword(i) {
			passwords++
		}
	}

	fmt.Println(passwords)
}
