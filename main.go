package main

import (
	"fmt"
	"strings"
)

func clearInput(text string) []string {
	tokens := strings.Fields(text)
	for idx := range tokens {
		tokens[idx] = strings.ToLower(tokens[idx])
	}
	return tokens
}

func main() {
	fmt.Println("Hello, World!")
}
