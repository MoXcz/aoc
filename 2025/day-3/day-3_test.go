package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/MoXcz/aoc/2025/internal/assert"
)

func TestPart1(t *testing.T) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
		return
	}

	jolts := bytes.Split(bytes.TrimSpace(file), []byte("\n"))
	result := joltageSumPart1(jolts)

	assert.Equal(t, result, 357)
}

func TestPart2(t *testing.T) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
		return
	}

	jolts := bytes.Split(bytes.TrimSpace(file), []byte("\n"))
	result := joltageSumPart2(jolts)

	assert.Equal(t, result, 3121910778619)
}
