package main

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/MoXcz/aoc/2025/internal/assert"
)

func TestPart1(t *testing.T) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(bytes.TrimSpace(file), []byte("\n"))
	actual := cephalopodSumPart1(lines)
	expected := 4277556

	assert.Equal(t, actual, expected)
}
