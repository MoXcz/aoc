package main

import (
	"testing"

	"github.com/MoXcz/aoc/2025/internal/assert"
)

func TestPart1(t *testing.T) {
	expected := 3
	actual := runDialPart1(50, 0, "test.txt")

	assert.Equal(t, actual, expected)
}

func TestPart2(t *testing.T) {
	expected := 6
	actual := runDial(50, 0, "test.txt")

	assert.Equal(t, actual, expected)
}
