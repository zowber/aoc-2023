package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Map struct {
	Name     string
	NumLines []string
}

func parseMap(lines []string) Map {

	var m Map

	for _, line := range lines {
		if strings.Contains(line, "map") {
			name := strings.TrimSuffix(line, " map:")
			m.Name = name
		} else if line != "" {
			m.NumLines = append(m.NumLines, line)
		} else {
			return m
		}
	}

	return m
}

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

	seeds := strings.Split(lines[0][6:], " ")
	fmt.Println(seeds)

	var mapBlocks [][]string
	var block []string

    for i, line := range lines[2:] {
        if i == len(lines[2:]) - 1 {
            block = append(block, line)
            mapBlocks = append(mapBlocks, block)
        } else if line != "" && i != len(lines[2:]) - 1 {
            block = append(block, line)
        } else {
            mapBlocks = append(mapBlocks, block)
            block = nil
        }
	}

	var maps []Map

	for _, block := range mapBlocks {
        m := parseMap(block)
        maps = append(maps, m)
        fmt.Println("map", m)
	}



	//fmt.Println("answer", answer)
}
