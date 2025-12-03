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

	jolts := bytes.Split(bytes.TrimSpace(file), []byte("\n"))
	fmt.Println(joltageSumPart2(jolts))
}

func joltageSumPart1(jolts [][]byte) int {
	var nums []int

	for _, jolt := range jolts {
		first := 0
		secondSeq := 0
		for i := 0; i < len(jolt)-1; i++ {
			num, err := strconv.Atoi(string(jolt[i]))
			if err != nil {
				log.Fatalf("Invalid number %d in sequence %s", num, jolt)
			}

			if first < num {
				first = num
				secondSeq = i
			}
		}

		second := 0
		for j := secondSeq + 1; j < len(jolt); j++ {
			num, err := strconv.Atoi(string(jolt[j]))
			if err != nil {
				log.Fatalf("Invalid number %d in sequence %s", num, jolt)
			}

			if second < num {
				second = num
			}
		}

		maxJolt, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))
		if err != nil {
			log.Fatalf("Invalid number conversion of %d and %d", first, second)
		}
		nums = append(nums, maxJolt)
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func joltageSumPart2(jolts [][]byte) int {
	var nums []int

	for _, jolt := range jolts {
		numsToSum := make([]int, 12)
		// fmt.Println("Jolt:", string(jolt))

		last := 0
		for k := 11; k >= 0; k-- {
			for i := last; i < len(jolt)-k; i++ {
				num, err := strconv.Atoi(string(jolt[i]))
				if err != nil {
					log.Fatalf("Invalid number %d in sequence %s", num, jolt)
				}

				if numsToSum[12-(k+1)] < num {
					numsToSum[12-(k+1)] = num
					last = i + 1
				}
				// fmt.Println("K:", k)
				// fmt.Println("Last:", last)
				// fmt.Println("Remain:", string(jolt[0:len(jolt)-k]))
				// fmt.Println(numsToSum)
			}
		}

		var acc string
		for _, num := range numsToSum {
			acc += strconv.Itoa(num)
		}
		// fmt.Println("Acc:", acc)

		maxJolt, err := strconv.Atoi(acc)
		if err != nil {
			log.Fatalf("Invalid number conversion of %s", acc)
		}
		nums = append(nums, maxJolt)
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
