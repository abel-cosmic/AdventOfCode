package main

import (
	"fmt"
	"regexp"
)

var wordToNumber = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func extractNumbersFromString(messyString string) ([]int, error) {
	wordPattern := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)
	matches := wordPattern.FindAllString(messyString, -1)
	var numericValues []int
	for _, word := range matches {
		numericValue, found := wordToNumber[word]
		if !found {
			return nil, fmt.Errorf("Unrecognized word: %s", word)
		}
		numericValues = append(numericValues, numericValue)
	}

	return numericValues, nil
}

func main() {
	messyString := "fivek4646fmj"
	result, err := extractNumbersFromString(messyString)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Extracted Numbers:", result)
	}
}
