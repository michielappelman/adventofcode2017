package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	permute "github.com/cweill/Permute-Golang"
)

func StarOne(phrases []string) int {
	var count int
	for _, phrase := range phrases {
		if validPassphrase(phrase) {
			count++
		}
	}
	return count
}

func validPassphrase(phrase string) bool {
	words := strings.Fields(phrase)
	seen := make(map[string]bool)
	for _, word := range words {
		if seen[word] {
			return false
		}
		seen[word] = true
	}
	return true
}

func StarTwo(phrases []string) int {
	var count int
	for _, phrase := range phrases {
		if validSecurePassphrase(phrase) {
			count++
		}
	}
	return count
}

func validSecurePassphrase(phrase string) bool {
	words := strings.Fields(phrase)
	seen := make(map[string]bool)
	for _, word := range words {
		for _, permutation := range permute.LexicographicPermutations(word) {
			if seen[permutation] {
				return false
			}
			seen[permutation] = true
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var passphrases []string
	for scanner.Scan() {
		passphrases = append(passphrases, scanner.Text())
	}
	fmt.Println("1:", StarOne(passphrases))
	fmt.Println("2:", StarTwo(passphrases))
}
