package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func part1() {
    log.Println("Part 1")

    file, err := os.Open("input.txt")
    if err != nil {
        log.Println("Err reading input file")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    numbers := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

    var filteredLines [][]string
    for scanner.Scan() {
        line := scanner.Text()
        var filteredLine []string
        for _, rune := range line {
            for _, number := range numbers {
                if string(rune) == number {
                    filteredLine = append(filteredLine, number)
                }
            }
        }
        filteredLines = append(filteredLines, filteredLine)
    }

    var finalLines []string
    for _, line := range filteredLines {
        finalLine := line[0] + (line[len(line)-1])
        finalLines = append(finalLines, finalLine)
    } 

    var answer int 
    for _, line := range finalLines {
        lineAsInt, _ := strconv.Atoi(line)
        answer += lineAsInt
    }
        
    log.Println("Answer:", answer)
}
