package day1

import (
	"log"
	"os"
)

func readInput() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Println("Err reading input file")
    }

    log.Println(file)
}
