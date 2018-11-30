// day three
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	part := 2

	input, err := ioutil.ReadFile("./day03/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	triangles := strings.Split(string(input), "\n")

	if part == 2 {
		var trianglesPt2 []string

		for i := 0; i < len(triangles); i += 3 {
			var sides1 = strings.Split(strings.TrimSpace(triangles[i]), " ")
			var sides2 = strings.Split(strings.TrimSpace(triangles[i+1]), " ")
			var sides3 = strings.Split(strings.TrimSpace(triangles[i+2]), " ")

			var sides1Filtered, sides2Filtered, sides3Filtered []string

			for j := 0; j < len(sides1); j++ {
				if sides1[j] != "" {
					sides1Filtered = append(sides1Filtered, sides1[j])
				}
			}

			for j := 0; j < len(sides2); j++ {
				if sides2[j] != "" {
					sides2Filtered = append(sides2Filtered, sides2[j])
				}
			}

			for j := 0; j < len(sides3); j++ {
				if sides3[j] != "" {
					sides3Filtered = append(sides3Filtered, sides3[j])
				}
			}

			trianglesPt2 = append(trianglesPt2, sides1Filtered[0]+" "+sides2Filtered[0]+" "+sides3Filtered[0])
			trianglesPt2 = append(trianglesPt2, sides1Filtered[1]+" "+sides2Filtered[1]+" "+sides3Filtered[1])
			trianglesPt2 = append(trianglesPt2, sides1Filtered[2]+" "+sides2Filtered[2]+" "+sides3Filtered[2])
		}

		triangles = trianglesPt2
	}

	possibleTriangles, impossibleTriangles := 0, 0

	for i := 0; i < len(triangles); i++ {
		var numbersAsStrings = strings.Split(strings.TrimSpace(triangles[i]), " ")
		var numbers []int
		possible := false

		for _, j := range numbersAsStrings {
			if j != "" {
				k, _ := strconv.Atoi(j)
				numbers = append(numbers, k)
			}
		}

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
