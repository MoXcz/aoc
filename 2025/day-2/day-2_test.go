package main

import (
	"bytes"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 1227775554
	ranges := bytes.Split(bytes.TrimSpace(file), []byte(","))
	sum := repeatedSumPart1(ranges)
	if want != sum {
		t.Errorf("got = %d, want = %d", sum, want)
	}
}

func TestPart2(t *testing.T) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 4174379265
	ranges := bytes.Split(bytes.TrimSpace(file), []byte(","))
	sum := repeatedSumPart2(ranges)
	if want != sum {
		t.Errorf("got = %d, want = %d", sum, want)
	}
}
