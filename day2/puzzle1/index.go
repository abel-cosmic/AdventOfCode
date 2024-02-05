package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../day2pzzuleinput.txt")
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
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	var gamesData [][][]string

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			gameValues := strings.TrimSpace(parts[1])
			colorPairs := strings.Split(gameValues, ";")
			var colorCounts [][]string
			for _, pair := range colorPairs {
				pair = strings.TrimSpace(pair)
				if len(pair) > 0 {
					colorAndCount := strings.SplitN(pair, " ", 2)
					if len(colorAndCount) == 2 {
						color := strings.TrimSpace(colorAndCount[1])
						count, err := strconv.Atoi(strings.TrimSpace(colorAndCount[0]))
						if err == nil {
							colorCounts = append(colorCounts, []string{color, strconv.Itoa(count)})
						}
					}
				}
			}
			gamesData = append(gamesData, colorCounts)
		}
	}

	// Print the result
	for i, data := range gamesData {
		fmt.Printf("%d: [", i+1)
		for j, pair := range data {
			fmt.Printf("[%s:%s]", pair[0], pair[1])
			if j < len(data)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
	}
}
