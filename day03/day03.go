// day three
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	part := 1

	input, err := ioutil.ReadFile("./day03/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	triangles := strings.Split(string(input), "\n")

	possibleTriangles, impossibleTriangles := 0, 0

	for i := 0; i < len(triangles); i++ {
		var numbersAsStrings = strings.Split(strings.TrimSpace(triangles[i]), " ")
		var numbers []int
		possible := false

		fmt.Println("============")

		if part == 1 {
			for _, j := range numbersAsStrings {
				if j != "" {
					k, _ := strconv.Atoi(j)
					numbers = append(numbers, k)
				}
			}
		} else {

		}

		fmt.Println(numbers)

		if numbers[0]+numbers[1] > numbers[2] &&
			numbers[0]+numbers[2] > numbers[1] &&
			numbers[1]+numbers[2] > numbers[0] {
			possible = true
		}

		if possible {
			possibleTriangles++
		} else {
			impossibleTriangles++
		}
	}

	fmt.Println(possibleTriangles, "possible triangles")
	fmt.Println(impossibleTriangles, "impossible triangles")

}
