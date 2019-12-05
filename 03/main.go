package main

import (
	"log"
	"math"
)

// Line defines a line and its coordinates
type Line struct {
	xMin, xMax int
	yMin, yMax int
	alignment  string // either - or |
}

// returns the manhattan distance of the point from the origin
func distance(x, y int) int {
	fx := float64(x) // converting to floats for math clarity
	fy := float64(y)

	dist := math.Abs(fx) + math.Abs(fy)
	return int(dist)
}

// Determines if lines a and b meet
func intersects(a, b Line) (meet bool, x, y int) {
	// first check if both lines are horizontal or vertical
	if a.alignment == b.alignment {
		return false, 0, 0
	}

	if a.alignment == "-" {
		// a will be '-' and b will be '|' so a's y min/max will be equal as will b's x
		if (a.xMin <= b.xMin && a.xMax >= b.xMin) && (a.yMin <= b.yMax && a.yMin >= b.yMin) {
			return true, b.xMin, a.yMin
		}
	} else {
		// a will be '|' and b will be '-' so a's x min/max will be equal as will b's y
		if (a.yMin <= b.yMin && a.yMax >= b.yMin) && (a.yMin <= b.yMax && a.yMin >= b.yMin) {
			return true, 1, 2 // not done yet
		}
	}

	return false, 0, 0
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func main() {

}
