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

	lines := bytes.Split(bytes.TrimSpace(file), []byte("\n"))
	actual := sumRollsOfPaperPart1(lines)

	assert.Equal(t, actual, 13)
}

func TestPart2(t *testing.T) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
		return
	}

	lines := bytes.Split(bytes.TrimSpace(file), []byte("\n"))
	result := sumRollsOfPaperPart2(lines)

	assert.Equal(t, result, 43)
}
