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

	ranges := bytes.Split(bytes.TrimSpace(file), []byte(","))
	fmt.Println(repeatedSumPart2(ranges))
}

func repeatedSumPart2(ranges [][]byte) int {
	nums := make(map[int]bool)

	for _, r := range ranges {
		values := bytes.Split(r, []byte("-"))
		lower, err := strconv.Atoi(string(values[0]))
		if err != nil {
			log.Fatalf("Invalid number")
		}
		upper, err := strconv.Atoi(string(values[1]))
		if err != nil {
			log.Fatalf("Invalid number")
		}

		for i := lower; i <= upper; i++ {
			str := strconv.Itoa(i)
			n := len(str)

			for l := 1; l <= n/2; l++ {
				if n%l != 0 {
					continue
				}
				pattern := str[:l]

				ok := true
				for pos := 0; pos < n; pos += l {
					if str[pos:pos+l] != pattern {
						ok = false
						break
					}
				}

				if ok {
					nums[i] = true
					break
				}
			}
		}
	}

	sum := 0
	for value := range nums {
		sum += value
	}
	return sum
}

func repeatedSumPart1(ranges [][]byte) int {
	nums := make(map[int]bool)
	for _, r := range ranges {
		values := bytes.Split(r, []byte("-"))
		lower, err := strconv.Atoi(string(values[0]))
		if err != nil {
			log.Fatalf("Invalid number")
		}
		upper, err := strconv.Atoi(string(values[1]))
		if err != nil {
			log.Fatalf("Invalid number")
		}

		for i := lower; i <= upper; i++ {
			str := strconv.Itoa(i)
			if len(str)%2 != 0 {
				continue
			}

			l := len(str) / 2
			if str[0:l] == str[l:] {
				nums[i] = true
			}
		}
	}
	sum := 0
	for value := range nums {
		sum += value
	}
	return sum
}
