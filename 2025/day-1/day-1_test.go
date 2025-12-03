package main

import "testing"

func TestPart1(t *testing.T) {
	want := 50
	clicks := runDialPart1(50, 0, "test.txt")
	if clicks != 3 {
		t.Errorf("got = %d, want = %d", clicks, want)
	}
}

func TestPart2(t *testing.T) {
	want := 50
	clicks := runDial(50, 0, "test.txt")
	if clicks != 6 {
		t.Errorf("got = %d, want = %d", clicks, want)
	}
}
