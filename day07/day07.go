// day seven
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type m map[rune]int

var signalCount []m

func main() {
	input, err := ioutil.ReadFile("./day07/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	ips := strings.Split(string(input), "\n")
	validIps, ssl := 0, 0

	for _, ip := range ips {
		sections := strings.FieldsFunc(ip, split)
		var insideBrackets, outsideBrackets []string

		for i := 0; i < len(sections); i++ {
			if i%2 == 1 {
				insideBrackets = append(insideBrackets, sections[i])
			} else {
				outsideBrackets = append(outsideBrackets, sections[i])
			}
		}

		if isValid(insideBrackets, outsideBrackets) {
			validIps++
		}

		if supportsSsl(insideBrackets, outsideBrackets) {
			ssl++
		}
	}

	fmt.Println(validIps, "valid IP addresses.")
	fmt.Println(ssl, "support SSL.")
}

func split(r rune) bool {
	return r == '[' || r == ']'
}

func isValid(insideBrackets []string, outsideBrackets []string) bool {
	hasAbbaInBrackets := false
	hasAbbaOutBrackets := false

	for _, section := range insideBrackets {
		if hasAbba(section) {
			hasAbbaInBrackets = true
		}
	}

	for _, section := range outsideBrackets {
		if hasAbba(section) {
			hasAbbaOutBrackets = true
		}
	}

	return hasAbbaOutBrackets && !hasAbbaInBrackets
}

func supportsSsl(insideBrackets []string, outsideBrackets []string) bool {
	for _, insideSection := range insideBrackets {
		hasAbaBool, babs := hasAba(insideSection)
		if hasAbaBool {
			for _, outsideSection := range outsideBrackets {
				for _, bab := range babs {
					if strings.Contains(outsideSection, bab) {
						return true
					}
				}
			}
		}
	}

	return false
}

func hasAbba(ipSection string) bool {
	splitString := []rune(ipSection)
	checker := []rune{'-', '-', '-', '-'}

	for _, character := range splitString {
		checker[0] = checker[1]
		checker[1] = checker[2]
		checker[2] = checker[3]
		checker[3] = character

		if checker[0] == checker[3] && checker[1] == checker[2] && checker[0] != checker[1] {
			return true
		}
	}

	return false
}

func hasAba(ipSection string) (bool, []string) {
	splitString := []rune(ipSection)
	checker := []rune{'-', '-', '-'}

	var babs []string

	for _, character := range splitString {
		checker[0] = checker[1]
		checker[1] = checker[2]
		checker[2] = character

		if checker[0] == checker[2] && checker[0] != checker[1] {
			babs = append(babs, string([]rune{checker[1], checker[0], checker[1]}))
		}
	}

	if len(babs) > 0 {
		return true, babs
	}

	return false, babs
}
