package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2() {
	log.Println("Part 2")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Println("Err reading input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	wordNumbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var newLine []int
	for _, line := range lines {
		for idx, wordNumber := range wordNumbers {
			if strings.HasPrefix(line, wordNumber) {
				newLine = append(newLine, idx+1)
			}
		}
		for _, number := range numbers {
			if string(line[0]) == number {
                foundNumber, _ := strconv.Atoi(number)
				newLine = append(newLine, foundNumber)
			}
		}
        log.Println(newLine)
	}

    var answer int
	log.Println("Answer:", answer)
}
