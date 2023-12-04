package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func extractNumericSequence(input string) []int {
	var seq []int

	for i, char := range input {
		if unicode.IsDigit(char) {
			currPos := i
			seq = append(seq, currPos)
		}
	}

	return seq
}

func extractSymbolPositions(input string) []int {
	var positions []int
	var symbols = []string{"*", "#", "+", "$"}

	for i, char := range input {
		for _, symbol := range symbols {
			if string(char) == symbol {
				positions = append(positions, i)
			}
		}
	}

	return positions
}

func part1() {

	log.Println("Part 1")

	file, err := os.Open("part1_example.txt")
	if err != nil {
		log.Println("Err reading input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		//fmt.Println(line)
		// consider the boundaries when looking for symbols
		// when on first col don't look left
		// when on last col don't look right

		// when on row 0 don't look above
		var prevLine string
		if i > 0 {
			prevLine = lines[i-1]
		}
		// when on last row don't look above
		var nextLine string
		if i < len(lines)-1 {
			nextLine = lines[i+1]
		}

		// get the indexes of number sequences
		digitSeq := extractNumericSequence(line)

		// get the symbols for prev, curr, and next lines
		if prevLine != "" {
			prevLineSymbols := extractSymbolPositions(prevLine)
			fmt.Println("prevSymbols:", prevLineSymbols)
		}

		lineSymbols := extractSymbolPositions(line)
		fmt.Println("lineSymbols", lineSymbols)

        var nextLineSymbols []int
		if nextLine != "" {
			nextLineSymbols = extractSymbolPositions(nextLine)
			fmt.Println("nextSymbols", nextLineSymbols)
		}
        // this doesn't work
		for _, symPos := range lineSymbols {
			for _, digitPos := range digitSeq {
                fmt.Println("symPos", symPos, "digitPos", digitPos)
                if symPos == (digitPos - 1) || symPos == digitPos || symPos == (digitPos + 1){
                    log.Println("Found symbol on line")
                }  
            }
		}
	}
}

//   0123456789
// 0|467..114..
// 1|...*......
// 2|..35..633.
// 3|......#...
// 4|617*......
// 5|.....+.58.
// 6|..592.....
// 7|......755.
// 8|...$.*....
// 9|.664.598..

