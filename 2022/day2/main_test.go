package main

import "testing"

func TestPlayHand(t *testing.T) {
	var tcs = []struct {
		name string
		opp  string
		you  string
		want int
	}{
		{
			name: "scissors vs scissors should be a draw",
			opp:  "C",
			you:  "Z",
			want: 3,
		},
		{
			name: "scissors vs rock should be a win",
			opp:  "C",
			you:  "X",
			want: 6,
		},
		{
			name: "scissors vs paper should be a loss",
			opp:  "C",
			you:  "Y",
			want: 0,
		},
		{
			name: "paper vs scissors should be a win",
			opp:  "B",
			you:  "Z",
			want: 6,
		},
		{
			name: "paper vs rock should be a loss",
			opp:  "B",
			you:  "X",
			want: 0,
		},
		{
			name: "paper vs paper should be a draw",
			opp:  "B",
			you:  "Y",
			want: 3,
		},
		{
			name: "rock vs rock should be a draw",
			opp:  "A",
			you:  "X",
			want: 3,
		},
		{
			name: "rock vs paper should be a win",
			opp:  "A",
			you:  "Y",
			want: 6,
		},
		{
			name: "rock vs scissors should be a loss",
			opp:  "A",
			you:  "Z",
			want: 0,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got := PlayHand(tt.opp, tt.you)
			if got != tt.want {
				t.Errorf("want: %d, got: %d", tt.want, got)
			}
		})
	}
}
