// day four
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./day04/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	rooms := strings.Split(string(input), "\n")

	sum := 0
	numRealRooms := 0

	for i := 0; i < len(rooms); i++ {
		name, id, checksum := decodeLine(rooms[i])

		if processRoom(name, checksum) {
			numRealRooms++
			sum += id

			decrypted := decryptRoom(name, id)

			if strings.Contains(decrypted, "pole") {
				fmt.Println(decrypted, id)
			}
		}
	}

	fmt.Println(numRealRooms, "real rooms out of", len(rooms), "total")
	fmt.Println(sum, "total id of real rooms")
}

func decodeLine(line string) (string, int, string) {
	splitLine := strings.Split(line, "-")
	last := len(splitLine) - 1

	idAndChecksum := strings.Split(splitLine[last], "[")

	name := strings.TrimSpace(strings.Join(splitLine[0:last], "-"))
	id, _ := strconv.Atoi(idAndChecksum[0])
	checksum := strings.TrimSpace(strings.TrimRight(idAndChecksum[1], "]"))

	return name, id, checksum
}

func processRoom(name string, checksum string) bool {
	var store = make(map[string]int)
	var order []string

	for _, char := range name {
		charS := string(char)
		if char != '-' {
			if !contains(order, charS) {
				order = append(order, charS)
			}
			store[charS] = store[charS] + 1
		}
	}

	sort.SliceStable(order, func(i, j int) bool {
		if store[order[i]] > store[order[j]] {
			return true
		} else if store[order[i]] < store[order[j]] {
			return false
		}

		return []rune(order[i])[0] < []rune(order[j])[0]
	})

	checksumCalculated := strings.Join(order[0:5], "")

	return checksumCalculated == checksum
}

func contains(slice []string, item string) bool {
	for _, itemInSlice := range slice {
		if itemInSlice == item {
			return true
		}
	}

	return false
}

func decryptRoom(name string, id int) string {
	iterations := id % 26 // num of letters in alphabet
	nameSplit := []rune(name)

	for i := 0; i < iterations; i++ {
		var newNameSplit []rune
		for _, char := range nameSplit {
			var newChar rune
			if string(char) == "-" || string(char) == " " {
				newChar = []rune(" ")[0]
			} else if string(char) == "z" {
				newChar = []rune("a")[0]
			} else {
				newChar = char + 1
			}

			newNameSplit = append(newNameSplit, newChar)
		}

		nameSplit = newNameSplit
	}

	return string(nameSplit)
}
