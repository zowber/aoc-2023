package day2

import (
	"bufio"
	"log"
	"slices"
	"strconv"
	"strings"
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

	//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	//Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	//Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	//Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	//Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

	type Result struct {
		Count int
		Color string
	}
	type Set struct {
		Results []Result
	}
	type Game struct {
		Id   int
		Sets []Set
	}

	var games []Game
	for i, line := range lines {

		rawGame := strings.Split(line, ": ")
        gameString := strings.Join(rawGame[1:], "")
		rawSets := strings.Split(gameString, "; ")

		var sets []Set
		for _, rawSet := range rawSets {
			rawResults := strings.Split(rawSet, ", ")

			var results []Result
			for _, rawResult := range rawResults {
				tmp := strings.Split(rawResult, " ")

				count, _ := strconv.Atoi(tmp[0])
				color := tmp[1]

				results = append(results, Result{Count: count, Color: color})
			}
			set := Set{Results: results}
			sets = append(sets, set)
		}
		game := Game{Id: i + 1, Sets: sets}
		games = append(games, game)
	}

	redTotal := 12
	greenTotal := 13
	blueTotal := 14

	var impossibleGames []int

	for _, g := range games {
		for _, s := range g.Sets {
			for _, r := range s.Results {
				if r.Color == "red" && r.Count > redTotal {
					impossibleGames = append(impossibleGames, g.Id)
				}
				if r.Color == "green" && r.Count > greenTotal {
					impossibleGames = append(impossibleGames, g.Id)
				}
				if r.Color == "blue" && r.Count > blueTotal {
					impossibleGames = append(impossibleGames, g.Id)
				}
			}
		}
	}

	compacted := slices.Compact(impossibleGames)

    var reference []int
    for i := 0; i < len(games); i++ {
        reference = append(reference, i+1)
    }

    for _, impossible := range compacted {
            reference[impossible-1] = 0 
    }

	answer := 0
	for _, val := range reference {
		answer += val
	}

	log.Println("Answer:", answer)
}
