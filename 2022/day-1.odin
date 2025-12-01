package main

import "core:fmt"
import "core:os"
import "core:strconv"
import "core:strings"

main :: proc() {
	input: string
	data, err := os.read_entire_file_or_err("./data/day-1.txt")
	if err != nil {
		fmt.printf("Error reading file: %v\n", err)
		return
	}

	total: int

	split_data := strings.split(string(data), "\n\n")
	for v, i in split_data {
		d := strings.split(v, "\n")
		sum := 0
		for elem, idx in d {
			elem_int, ok := strconv.parse_int(elem)
			if ok {
				sum += elem_int
			}
		}
		if total < sum {
			total = sum
		}
	}

	fmt.println(total)
}

