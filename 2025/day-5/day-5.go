package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return
	}

	lines := bytes.Split(bytes.TrimSpace(file), []byte("\n\n"))

	ranges := bytes.Split(lines[0], []byte("\n"))
	// ingredientIDs := bytes.Split(lines[1], []byte("\n"))

	fmt.Println(sumFreshIngredientIDsPart2(ranges))
}

func sumFreshIngredientIDsPart1(ranges, ingredientIDs [][]byte) int {
	ids := make([]int, len(ingredientIDs))
	for _, id := range ingredientIDs {
		val, err := strconv.Atoi(string(id))
		if err != nil {
			log.Fatal(err)
		}
		ids = append(ids, val)
	}

	nums := make(map[int]bool)

	for _, r := range ranges {
		values := bytes.Split(r, []byte("-"))
		lower, err := strconv.Atoi(string(values[0]))
		if err != nil {
			log.Fatal(err)
		}
		upper, err := strconv.Atoi(string(values[1]))
		if err != nil {
			log.Fatal(err)
		}

		for _, id := range ids {
			if id >= lower && id <= upper {
				nums[id] = true
			}
		}
	}

	return len(nums)
}

// extremely slow, a better solution would include using some kind of 'merge'
// operation that would rm the need to go through whole ranges by identifying
// 'lower-upper' ranges that have been gone through
// 3-5 (5-3 -> 2+1 -> 3)
// 12-36 (36-12 -> 24+1 -> 25)
// (too lazy to do it)
func sumFreshIngredientIDsPart2(ranges [][]byte) int {
	nums := make(map[int]bool)

	for _, r := range ranges {
		values := bytes.Split(r, []byte("-"))
		lower, err := strconv.Atoi(string(values[0]))
		if err != nil {
			log.Fatal(err)
		}
		upper, err := strconv.Atoi(string(values[1]))
		if err != nil {
			log.Fatal(err)
		}

		for i := lower; i <= upper; i++ {
			nums[i] = true
		}
	}

	return len(nums)
}
