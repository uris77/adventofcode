package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestExpandSections(t *testing.T) {
	var tcs = []struct {
		name  string
		input string
		want  []int
	}{
		{
			name:  "example 1",
			input: "2-4",
			want:  []int{2, 3, 4},
		},
		{
			name:  "example 2",
			input: "6-8",
			want:  []int{6, 7, 8},
		},
		{
			name:  "example 3",
			input: "6-6",
			want:  []int{6},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandSections(tt.input)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestIsContainedIn(t *testing.T) {
	var tcs = []struct {
		name   string
		first  []int
		second []int
		want   bool
	}{
		{
			name:   "example 1",
			first:  []int{2, 3, 4},
			second: []int{6, 7, 8},
			want:   false,
		},
		{
			name:   "example 2",
			first:  []int{2, 3, 4, 5, 6, 7, 8},
			second: []int{6, 7},
			want:   true,
		},
		{
			name:   "example 3",
			first:  []int{6, 7},
			second: []int{2, 3, 4, 5, 6, 7, 8},
			want:   false,
		},
		{
			name:   "example 4",
			first:  []int{2, 3, 4},
			second: []int{4},
			want:   true,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {

			got := IsContainedIn(tt.first, tt.second)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestOverlaps(t *testing.T) {
	var tcs = []struct {
		name   string
		first  []int
		second []int
		want   bool
	}{
		{
			name:   "example 1",
			first:  []int{5, 6, 7},
			second: []int{7, 8, 9},
			want:   true,
		},
		{
			name:   "example 2",
			first:  []int{2, 3, 4, 5, 6, 7, 8},
			second: []int{3, 4, 5, 6, 7},
			want:   true,
		},
		{
			name:   "example 3",
			first:  []int{6},
			second: []int{4, 5, 6},
			want:   true,
		},
		{
			name:   "example 4",
			first:  []int{4, 5, 6},
			second: []int{6},
			want:   true,
		},
		{
			name:   "example 5",
			first:  []int{5, 6, 7},
			second: []int{9, 10, 11},
			want:   false,
		},
		{
			name:   "example 6",
			first:  []int{9, 10, 11},
			second: []int{5, 6, 7},
			want:   false,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got := Overlaps(tt.first, tt.second)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
