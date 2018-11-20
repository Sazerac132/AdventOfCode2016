// day two
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./day02/input.txt")
	part := 2 // part 1 or 2
	var grid [][]rune

	if part == 1 {
		grid = [][]rune{
			{'0', '0', '0', '0', '0'},
			{'0', '1', '2', '3', '0'},
			{'0', '4', '5', '6', '0'},
			{'0', '7', '8', '9', '0'},
			{'0', '0', '0', '0', '0'}}
	} else {
		grid = [][]rune{
			{'0', '0', '0', '0', '0', '0', '0'},
			{'0', '0', '0', '1', '0', '0', '0'},
			{'0', '0', '2', '3', '4', '0', '0'},
			{'0', '5', '6', '7', '8', '9', '0'},
			{'0', '0', 'A', 'B', 'C', '0', '0'},
			{'0', '0', '0', 'D', '0', '0', '0'},
			{'0', '0', '0', '0', '0', '0', '0'}}
	}

	var x, y int

	if err != nil {
		fmt.Println(err)
	}

	inputLines := strings.Split(string(input), "\n")

	for i := 0; i < len(inputLines); i++ {
		y = 2

		if part == 1 {
			x = 2
		} else {
			x = 4
		}

		for j := 0; j < len(inputLines[i]); j++ {
			var direction = string(inputLines[i][j])

			switch direction {
			case "U":
				if grid[y-1][x] != '0' {
					y--
				}
			case "D":
				if grid[y+1][x] != '0' {
					y++
				}
			case "L":
				if grid[y][x-1] != '0' {
					x--
				}
			case "R":
				if grid[y][x+1] != '0' {
					x++
				}
			}
		}

		fmt.Println("Row ends on ", string(grid[y][x]))
	}
}
