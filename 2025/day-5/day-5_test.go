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

	lines := bytes.Split(bytes.TrimSpace(file), []byte("\n\n"))
	ranges := bytes.Split(lines[0], []byte("\n"))
	ingredientIDs := bytes.Split(lines[1], []byte("\n"))

	result := sumFreshIngredientIDsPart1(ranges, ingredientIDs)

	assert.Equal(t, result, 3)
}

func TestPart2(t *testing.T) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
		return
	}

	lines := bytes.Split(bytes.TrimSpace(file), []byte("\n\n"))
	ranges := bytes.Split(lines[0], []byte("\n"))

	result := sumFreshIngredientIDsPart2(ranges)

	assert.Equal(t, result, 14)
}
