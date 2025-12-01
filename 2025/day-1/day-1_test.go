package main

import "testing"

func Test_runDial(t *testing.T) {
	tests := []struct {
		name     string
		dial     int
		filename string
		want     int
	}{
		{
			name:     "Example input",
			dial:     50,
			filename: "test.txt",
			want:     6,
		},
		{
			name:     "Test right count of 0",
			dial:     50,
			filename: "test_2.txt",
			want:     10,
		},
		{
			name:     "Negative count",
			dial:     50,
			filename: "test_3.txt",
			want:     3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clicks := runDial(tt.dial, 0, tt.filename)
			if tt.want != clicks {
				t.Errorf("got = %d, want = %d", clicks, tt.want)
			}
		})
	}
}

func Test_runDialPart1(t *testing.T) {
	tests := []struct {
		name     string
		dial     int
		filename string
		want     int
	}{
		{
			name:     "Example input",
			dial:     50,
			filename: "test.txt",
			want:     3,
		},
		{
			name:     "Test right count of 0",
			dial:     50,
			filename: "test_2.txt",
			want:     0,
		},
		{
			name:     "Negative count",
			dial:     50,
			filename: "test_3.txt",
			want:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clicks := runDialPart1(tt.dial, 0, tt.filename)
			if tt.want != clicks {
				t.Errorf("got = %d, want = %d", clicks, tt.want)
			}
		})
	}
}
