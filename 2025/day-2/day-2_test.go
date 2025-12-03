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
		t.Fatal(err)
	}

	expected := 1227775554
	ranges := bytes.Split(bytes.TrimSpace(file), []byte(","))
	actual := repeatedSumPart1(ranges)

	assert.Equal(t, actual, expected)
}

func TestPart2(t *testing.T) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expected := 4174379265
	ranges := bytes.Split(bytes.TrimSpace(file), []byte(","))
	actual := repeatedSumPart2(ranges)

	assert.Equal(t, actual, expected)
}
