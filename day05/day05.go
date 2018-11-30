// day five
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	part := 2
	input := "abbhdwsy"
	incrementor := 0

	password := ""
	password2 := []rune{'_', '_', '_', '_', '_', '_', '_', '_'}

	if part == 1 {
		for len(password) < 8 {
			data := []byte(input + strconv.Itoa(incrementor))
			md5 := md5.Sum(data)
			md5Str := hex.EncodeToString(md5[:16])

			if md5Str[0:5] == "00000" {
				password += md5Str[5:6]
			}

			incrementor++
		}
	} else {
		charsFound := 0
		for charsFound < 8 {
			data := []byte(input + strconv.Itoa(incrementor))
			md5 := md5.Sum(data)
			md5Str := hex.EncodeToString(md5[:16])

			if md5Str[0:5] == "00000" {
				position, err := strconv.Atoi(md5Str[5:6])
				char := md5Str[6:7]

				if err == nil && position >= 0 && position < 8 {
					if password2[position] == '_' {
						password2[position] = []rune(char)[0]
						charsFound++

						fmt.Println(string(password2))
					}
				}

			}

			incrementor++
		}
	}

	if part == 1 {
		fmt.Println("password:", password)
	} else {
		fmt.Println("password:", string(password2))
	}
}
