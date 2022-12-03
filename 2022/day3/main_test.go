package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestHalveRunes(t *testing.T) {
	var tcs = []struct {
		name  string
		input []rune
		want  [][]rune
	}{
		{
			name: "example 1",
			input: []rune{'F', 'z', 'Q', 'r', 'h', 'Q', 'p', 'J', 't', 'J', 'M', 'F', 'z', 'l', 'p', 'p', 'l', 'r',
				'T', 'W', 'j', 'T', 'n', 'T', 'T', 'r', 'j', 'V', 's', 'V', 'v', 'v', 'T', 'n', 'T', 's'},
			want: [][]rune{
				{'F', 'z', 'Q', 'r', 'h', 'Q', 'p', 'J', 't', 'J', 'M', 'F', 'z', 'l', 'p', 'p', 'l', 'r'},
				{'T', 'W', 'j', 'T', 'n', 'T', 'T', 'r', 'j', 'V', 's', 'V', 'v', 'v', 'T', 'n', 'T', 's'},
			},
		},
		{
			name: "example 2",
			input: []rune{'F', 'z', 'Q', 'r', 'h', 'Q', 'p', 'J', 't', 'J', 'M', 'F', 'z', 'l', 'p', 'p',
				'T', 'W', 'j', 'T', 'n', 'T', 'T', 'r', 'j', 'V', 's', 'V', 'v', 'v', 'T', 'n'},
			want: [][]rune{
				{'F', 'z', 'Q', 'r', 'h', 'Q', 'p', 'J', 't', 'J', 'M', 'F', 'z', 'l', 'p', 'p'},
				{'T', 'W', 'j', 'T', 'n', 'T', 'T', 'r', 'j', 'V', 's', 'V', 'v', 'v', 'T', 'n'},
			},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got := HalveRunes(tt.input)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("got: %c; want: %c", got, tt.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	var tcs = []struct {
		name  string
		input [][]rune
		want  []rune
	}{
		{
			name: "one duplicate",
			input: [][]rune{
				{'c', 'w', 'T', 'f', 'L', 'w', 'B', 'V', 'w', 'C', 'W', 'b', 'L', 'c', 'V', 'T', 'V', 'V', 'v'},
				{'r', 'd', 'n', 'd', 'G', 'j', 'M', 'H', 'r', 'n', 'G', 'J', 't', 'n', 't', 't', 'd', 'M', 'C'}},
			want: []rune{'C'},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersection(tt.input)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("got: %c; want: %c", got, tt.want)
			}
		})
	}
}
