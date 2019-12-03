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

func printIntcode(data []int) {
	for i := 0; i < len(data); i++ {
		if i == len(data)-1 {
			fmt.Printf("%d\n", data[i])
		} else {
			fmt.Printf("%d,", data[i])
		}
	}
}

func main() {
	var intcode []int

	// getting input
	file, err := os.Open("input.txt")
	check(err)

	// setting up scanner
	scan := bufio.NewScanner(file)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// split function that seperates on commas
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	scan.Split(onComma)

	// adding to intcode
	for scan.Scan() {
		number, err := strconv.Atoi(scan.Text())
		check(err)

		intcode = append(intcode, number)
	}
	fmt.Printf("Intcode before: ")
	printIntcode(intcode)

	// restoring to the "1202 program alarm" state
	intcode[1] = 12
	intcode[2] = 2

	// running intcode
	for i := 0; i < len(intcode); i += 4 {
		// loop through but on an interval of 4
		if intcode[i] == 99 {
			break
		} else {
			pos1 := intcode[i+1]
			pos2 := intcode[i+2]
			pos3 := intcode[i+3]
			if intcode[i] == 1 {
				intcode[pos3] = intcode[pos1] + intcode[pos2]
			} else if intcode[i] == 2 {
				intcode[pos3] = intcode[pos1] * intcode[pos2]
			} else {
				log.Fatal("Unknown opcode!")
			}
		}
	}

	fmt.Printf("Intcode after: ")
	printIntcode(intcode)
}
