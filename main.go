package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func IsNumeric(s string) bool {
	myRune := []rune(s)
	for i := 0; i < len(myRune); i++ {
		if myRune[i] < '0' || myRune[i] > '9' {
			return false
		}
	}
	return true
}

func isDuplicat(str string) bool {
	seen := make(map[string]bool)

	parts := strings.Fields(str)

	for _, part := range parts {
		if seen[part] {
			return true
		}
		seen[part] = true
	}
	return false
}

func main() {
	if len(os.Args) == 2 {
		input := os.Args[1]
		numbers := strings.Fields(input)
		for _, num := range numbers {
			for _, char := range num {
				if !unicode.IsDigit(char) || isDuplicat(input){
					fmt.Println("Error")
					return
				}
			}
		}
	}
}
