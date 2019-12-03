package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func main() {
	var numbers []int
	var fuel int

	file, err := os.Open("input.txt")
	check(err)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		check(err)
		numbers = append(numbers, num)
	}

	total := 0
	for _, number := range numbers {
		fuel = number/3 - 2
		total += fuel
	}

	fmt.Println(total)
}
