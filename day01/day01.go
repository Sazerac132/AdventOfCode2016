// day one
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := "R1, R1, R3, R1, R1, L2, R5, L2, R5, R1, R4, L2, R3, L3, R4, L5, R4, R4, R1, L5, L4, R5, R3, L1, R4, R3, L2, L1, R3, L4, R3, L2, R5, R190, R3, R5, L5, L1, R54, L3, L4, L1, R4, R1, R3, L1, L1, R2, L2, R2, R5, L3, R4, R76, L3, R4, R191, R5, R5, L5, L4, L5, L3, R1, R3, R2, L2, L2, L4, L5, L4, R5, R4, R4, R2, R3, R4, L3, L2, R5, R3, L2, L1, R2, L3, R2, L1, L1, R1, L3, R5, L5, L1, L2, R5, R3, L3, R3, R5, R2, R5, R5, L5, L5, R2, L3, L5, L2, L1, R2, R2, L2, R2, L3, L2, R3, L5, R4, L4, L5, R3, L4, R1, R3, R2, R4, L2, L3, R2, L5, R5, R4, L2, R4, L1, L3, L1, L3, R1, R2, R1, L5, R5, R3, L3, L3, L2, R4, R2, L5, L1, L1, L5, L4, L1, L1, R1"

	instructions := strings.Split(input, ", ")

	direction, horizontal, vertical := 'N', 0, 0

	var placesVisited, placesVisitedTwice []string

	for i := 0; i < len(instructions); i++ {
		turningDirection := rune(instructions[i][0])

		distanceString := instructions[i][1:len(instructions[i])]
		distanceInt64, _ := strconv.ParseInt(distanceString, 10, 64)
		distance := int(distanceInt64)

		direction = turn(direction, turningDirection)

		for j := 0; j < distance; j++ {
			switch direction {
			case 'N':
				vertical++
			case 'S':
				vertical--
			case 'E':
				horizontal++
			case 'W':
				horizontal--
			}

			var thisPlace = "V" + strconv.Itoa(vertical) + "H" + strconv.Itoa(horizontal)

			if contains(placesVisited, thisPlace) {
				placesVisitedTwice = append(placesVisitedTwice, thisPlace)
			}

			placesVisited = append(placesVisited, thisPlace)
		}

	}

	fmt.Println("horizontal distance:", horizontal)
	fmt.Println("vertical distance:", vertical)

	fmt.Println("total distance:", math.Abs(float64(horizontal))+math.Abs(float64(vertical)))
	fmt.Println("visited twice:", placesVisitedTwice)
}

func turn(currentDirection rune, turningDirection rune) rune {
	if turningDirection == 'L' {
		switch currentDirection {
		case 'N':
			return 'W'
		case 'W':
			return 'S'
		case 'S':
			return 'E'
		case 'E':
			return 'N'
		}
	}

	switch currentDirection {
	case 'N':
		return 'E'
	case 'E':
		return 'S'
	case 'S':
		return 'W'
	case 'W':
		return 'N'
	}

	return '?'

}

func contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}

	return false
}
