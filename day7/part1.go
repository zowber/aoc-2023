package day7

import (
	"bufio"
	"log"
	"os"
)

func part1() {

	log.Println("Part 1")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Println("Err reading input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//fmt.Println("answer", answer)
}
