package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	clicks := runDial(50, 0, "input.txt")
	fmt.Println("Count:", clicks)
}

func turnDial(dial, turn, clicks int, side string) (int, int) {
	for range turn {
		if side == "L" {
			dial--
			if dial < 0 {
				dial = 99
			}
		} else if side == "R" {
			dial++
			if dial > 99 {
				dial = 0
			}
		}

		if dial == 0 {
			clicks++
		}
	}

	return dial, clicks
}

func turnDialPart1(dial, turn, clicks int, side string) (int, int) {
	if side == "L" {
		dial = dial - turn
	} else if side == "R" {
		dial = dial + turn
	}

	for dial > 99 || dial < 0 {
		if dial < 0 {
			dial = 100 + dial
		}
		if dial > 99 {
			dial = dial - 100
		}
	}

	if dial == 0 {
		clicks += 1
	}
	return dial, clicks
}

func runDial(dial, clicks int, filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	turns := bytes.SplitSeq(bytes.Trim(file, "\n"), []byte("\n"))
	for t := range turns {
		side := string(t[0])
		turn, err := strconv.Atoi(string(t[1:]))
		if err != nil {
			log.Fatal(err)
		}
		dial, clicks = turnDial(dial, turn, clicks, side)
	}

	return clicks
}

func runDialPart1(dial, clicks int, filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	turns := bytes.SplitSeq(bytes.Trim(file, "\n"), []byte("\n"))
	for t := range turns {
		side := string(t[0])
		turn, err := strconv.Atoi(string(t[1:]))
		if err != nil {
			log.Fatal(err)
		}
		dial, clicks = turnDialPart1(dial, turn, clicks, side)
	}

	return clicks
}
