package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var results []int

func separateNumbers(characters []string) ([]int, []string) {
	var numbers []int
	var nonNumbers []string

	for _, char := range characters {
		if unicode.IsDigit(rune(char[0])) {
			num, err := strconv.Atoi(char)
			if err == nil {
				numbers = append(numbers, num)
			}
		} else {
			nonNumbers = append(nonNumbers, char)
		}
	}
	return numbers, nonNumbers
}

func main() {
	var alphabetNumbers = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	m := make(map[string]int)

	for i, word := range alphabetNumbers {
		m[word] = i + 1
	}

	file, err := os.Open("day1puzzle1input.txt")
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
		numbers, nonNumbers := separateNumbers(characters)
		fmt.Println("Numbers:", numbers)
		fmt.Println("Non Numbers:", nonNumbers)

		if len(numbers) >= 2 {
			firstDigit := numbers[0]
			lastDigit := numbers[len(numbers)-1]
			twoDigitNumber := firstDigit*10 + lastDigit
			results = append(results, twoDigitNumber)
		} else if len(numbers) == 1 {
			firstDigit := numbers[0]
			lastDigit := numbers[0]
			twoDigitNumber := firstDigit*10 + lastDigit
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
