package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(bytes.TrimSpace(file), []byte("\n"))
	fmt.Println(cephalopodSumPart1(lines))
}

func cephalopodSumPart1(lines [][]byte) int {
	listNums := make(map[int][]int)

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]

		nums := bytes.SplitSeq(line, []byte(" "))
		j := 0
		for char := range nums {
			if len(char) == 0 {
				continue
			}

			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}

			listNums[j] = append(listNums[j], num)
			j += 1
		}
	}

	total := 0
	i := 0
	for _, char := range lines[len(lines)-1] {
		if char == 32 {
			continue
		}

		agg := 1
		switch char {
		case 42:
			for _, num := range listNums[i] {
				agg *= num
			}

			total += agg
			i += 1
		case 43:
			for _, num := range listNums[i] {
				agg += num
			}

			total += agg - 1
			i += 1
		}
	}

	return total
}
