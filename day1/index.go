package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// learning go on the way
/**
 * @Description: main function
 * @param
 * @return
**/
func main() {
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
		for i := range characters {
			char := characters[i]
			fmt.Println(char)
		}
		fmt.Println(characters)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
