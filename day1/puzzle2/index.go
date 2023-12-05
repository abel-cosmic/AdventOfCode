package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var results []int
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

func separateNumbers(characters []string) []int {
	var mixed []int
	var currentString string

	for _, char := range characters {
		if unicode.IsDigit(rune(char[0])) {
			num, err := strconv.Atoi(char)
			if err == nil {
				if currentString != "" {
					extractedNumbers, err := extractNumbersFromString(currentString)
					if err == nil {
						mixed = append(mixed, extractedNumbers...)
					} else {
						fmt.Println("Error extracting numbers:", err)
					}
					currentString = ""
				}
				mixed = append(mixed, num)
			}
		} else {
			currentString += char
		}
	}
	if currentString != "" {
		extractedNumbers, err := extractNumbersFromString(currentString)
		if err == nil {
			mixed = append(mixed, extractedNumbers...)
		} else {
			fmt.Println("Error extracting numbers:", err)
		}
	}

	return mixed
}

func main() {
	file, err := os.Open("../day1puzzle1input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		characters := strings.Split(line, "")
		mixed := separateNumbers(characters)
		if len(mixed) >= 2 {
			firstDigit := mixed[0]
			lastDigit := mixed[len(mixed)-1]
			twoDigitNumber := firstDigit*10 + lastDigit
			results = append(results, twoDigitNumber)
		} else if len(mixed) == 1 {
			firstDigit := mixed[0]
			twoDigitNumber := firstDigit*10 + firstDigit
			results = append(results, twoDigitNumber)
		}
	}
	totalSum := 0
	for _, result := range results {
		totalSum += result
	}
	fmt.Println("Total Sum:", totalSum)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
