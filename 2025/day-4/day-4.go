package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return
	}

	lines := bytes.Split(bytes.TrimSpace(file), []byte("\n"))
	fmt.Println(sumRollsOfPaperPart2(lines))
}

func sumRollsOfPaperPart1(lines [][]byte) int {
	total := 0

	for i, line := range lines {
		// fmt.Println("Current line:", string(line))
		// fmt.Println()
		// '@' -> 64 in ascii

		// ..@@.@@@@.  <- top
		// @@@.@.@.@@  <- HERE
		// @@@@@.@.@@  <- bottom
		for charIdx, char := range line {
			counter := 0
			if char == 46 {
				continue // '.'
			}
			//  @
			//  .
			//  @
			if i > 0 {
				if lines[i-1][charIdx] == 64 {
					counter += 1
				}

			}
			if i < len(lines)-1 {
				if lines[i+1][charIdx] == 64 {
					counter += 1
				}
			}

			// @
			// @.
			// @
			// it adds up these with '.' being 'char'
			if charIdx > 0 { // check to the left
				if line[charIdx-1] == 64 {
					counter += 1
				}

				if i > 0 { // avoid adding 'top'
					if lines[i-1][charIdx-1] == 64 {
						counter += 1
					}
				}

				if i < len(lines)-1 {
					if lines[i+1][charIdx-1] == 64 {
						counter += 1
					}
				}

			}

			//  @
			// .@
			//  @
			// it adds up these with '.' being 'char'
			if charIdx < len(line)-1 {
				if line[charIdx+1] == 64 {
					counter += 1
				}

				if i > 0 {
					if lines[i-1][charIdx+1] == 64 {
						counter += 1
					}
				}
				if i < len(lines)-1 {
					if lines[i+1][charIdx+1] == 64 {
						counter += 1
					}
				}

			}

			if counter < 4 {
				total += 1
			}
		}
	}
	return total
}

func sumRollsOfPaperPart2(lines [][]byte) int {
	total := 0

	for i := range lines {
		line := lines[i]
		// '@' -> 64 in ascii
		// '.' -> 46 in ascii

		for charIdx, char := range line {
			counter := 0
			if char == 46 {
				continue // '.'
			}
			//  @
			//  .
			//  @
			if i > 0 {
				if lines[i-1][charIdx] == 64 {
					counter += 1
				}

			}
			if i < len(lines)-1 {
				if lines[i+1][charIdx] == 64 {
					counter += 1
				}
			}

			// @
			// @.
			// @
			if charIdx > 0 { // check to the left
				if line[charIdx-1] == 64 {
					counter += 1
				}

				if i > 0 { // avoid adding 'top'
					if lines[i-1][charIdx-1] == 64 {
						counter += 1
					}
				}

				if i < len(lines)-1 {
					if lines[i+1][charIdx-1] == 64 {
						counter += 1
					}
				}

			}

			//  @
			// .@
			//  @
			if charIdx < len(line)-1 {
				if line[charIdx+1] == 64 {
					counter += 1
				}

				if i > 0 {
					if lines[i-1][charIdx+1] == 64 {
						counter += 1
					}
				}
				if i < len(lines)-1 {
					if lines[i+1][charIdx+1] == 64 {
						counter += 1
					}
				}

			}

			if counter < 4 {
				total += 1
				lines[i][charIdx] = 46
				if i > 0 {
					// long live recursion!
					total += sumRollsOfPaperPart2(lines)
					// This basically re-does all the lines after every character that
					// was deleted, essentially adding any new deleted rolls of paper.
					// This can probably be optimized by just recurisvely going through
					// the previous lines, instead of spawning a new function call
					// with each deleted character
				}
			}
		}
	}
	return total
}
