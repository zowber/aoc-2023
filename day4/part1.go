package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day1() {
    
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

    // Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
    // Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19

    var answer int

    for i, line := range lines {
        fmt.Println("Game", i + 1)
		firstSplit := strings.Split(line, ":")
        dataStr := strings.Join(firstSplit[1:], "")
        
        woDblSpc := strings.Split(strings.ReplaceAll(dataStr, "  ", " "), "[1:]")
        woDblSpcStr := strings.Join(woDblSpc, "") 

        numSplit := strings.Split(woDblSpcStr[1:], " | ")
        
        winningNums := strings.Split(numSplit[0], " ")
        myNums :=  strings.Split(numSplit[1], " ")

        fmt.Println("winning", winningNums)
        fmt.Println("have", myNums)

        var matchingNums []int

        for _, myNum := range myNums {
            for _, winningNum := range winningNums {
                if myNum == winningNum {
                    matchingNum, _ := strconv.Atoi(myNum)
                    matchingNums = append(matchingNums, matchingNum)
                }
            }
        }
        fmt.Println("matching", matchingNums)

        matchingNumsCount := len(matchingNums)
        var points int

        switch matchingNumsCount {
        case 0:
            fmt.Println("no matches")
        case 1:
            fmt.Println("one match")
            points = 1
        default:
            points = 1
            for i := 0; i < (matchingNumsCount - 1); i++ {
                points = points * 2
            }
        }

        fmt.Println("game points", points)

        answer += points
    }

    fmt.Println("answer", answer)
}
