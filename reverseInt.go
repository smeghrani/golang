package main

import (
	"fmt"
	"strconv"
)

func reversedInt(takeAnInt int) int {

	convert := strconv.Itoa(takeAnInt)
	variableRunes := []rune(convert)
	lenOfVariableRunes := len(variableRunes)

	for i, j := 0, lenOfVariableRunes-1; i < j; i, j = i+1, j-1 {
		variableRunes[i], variableRunes[j] = variableRunes[j], variableRunes[i]
	}

	reversedInt, _ := strconv.Atoi(string(variableRunes))
	return reversedInt
}

func main() {
	originalInt := 123456789
	reversedInt := reversedInt(originalInt)
	fmt.Printf("Original: %d\nReversed: %d\n", originalInt, reversedInt)
}
