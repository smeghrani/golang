package main

import "fmt"

func reverseString(takeAString string) string {

	variableRune := []rune(takeAString)

	checkLenOfVariableRunes := len(variableRune)

	for i, j := 0, checkLenOfVariableRunes-1; i < j; i, j = i+1, j-1 {
		variableRune[i], variableRune[j] = variableRune[j], variableRune[i]
	}

	reversedString := string(variableRune)

	return reversedString
}

func m2() {

	originalString := "Mumbai"

	reversedString := reverseString(originalString)

	fmt.Printf("Original: %s\nReversed: %s\n", originalString, reversedString)
}
