// day six
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type m map[rune]int

var signalCount []m

func main() {
	part := 2
	input, err := ioutil.ReadFile("./day06/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	signals := strings.Split(string(input), "\n")

	//run once only
	for i := 0; i < len(signals[0]); i++ {
		signalCount = append(signalCount, make(m))
	}

	for _, signal := range signals {
		for pos, char := range []rune(signal) {
			signalCount[pos][char] = signalCount[pos][char] + 1
		}
	}

	for _, charOccurences := range signalCount {
		var possibilities []rune
		for k := range charOccurences {
			possibilities = append(possibilities, k)
		}

		sort.SliceStable(possibilities, func(i, j int) bool {
			if part == 1 {
				return charOccurences[possibilities[i]] > charOccurences[possibilities[j]]
			}

			return charOccurences[possibilities[i]] < charOccurences[possibilities[j]]
		})

		fmt.Println(string(possibilities[0]))
	}
}
