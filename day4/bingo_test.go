package day4

import (
	"reflect"
	"testing"
)

func Test_playBingo(t *testing.T) {
	type args struct {
		boards  []bingoboard
		numgers []string
	}
	bb, in := readBoards("example.txt")
	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "Happy",
		args: args{
			boards:  bb,
			numgers: in,
		},
		want: 4512,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := caluculateScore(tt.args.boards, tt.args.numgers)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("playBingo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateLast(t *testing.T) {
	type args struct {
		boards  []bingoboard
		numbers []string
	}

	bb, in := readBoards("input.txt")

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy",
			args: args{
				boards:  bb,
				numbers: in,
			},
			want: 17408,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateLast(tt.args.boards, tt.args.numbers); got != tt.want {
				t.Errorf("calculateLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
