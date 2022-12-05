package day5

import (
	"reflect"
	"testing"
)

func Test_rearageStacks(t *testing.T) {
	type args struct {
		stacks string
		moves  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{
				stacks: "example_stacks.txt",
				moves:  "example_moves.txt",
			},
			want: "CMZ",
		},
		{
			name: "input",
			args: args{
				stacks: "input_stacks.txt",
				moves:  "input_moves.txt",
			},
			want: "SHQWSRBDL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearageStacks(tt.args.stacks, tt.args.moves); got != tt.want {
				t.Errorf("rearageStacks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadStacks(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{
			name: "happy load",
			args: args{
				s: "example_stacks.txt",
			},
			want: [][]rune{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadStacks(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadStacks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadMoves(t *testing.T) {
	type args struct {
		m string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "happy moves",
			args: args{
				m: "example_moves.txt",
			},
			want: [][]int{{1, 2, 1}, {3, 1, 3}, {2, 2, 1}, {1, 1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadMoves(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rearageStacks9001(t *testing.T) {
	type args struct {
		s string
		m string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{
				s: "example_stacks.txt",
				m: "example_moves.txt",
			},
			want: "MCD",
		},
		{
			name: "input",
			args: args{
				s: "input_stacks.txt",
				m: "input_moves.txt",
			},
			want: "JHJH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearageStacks9001(tt.args.s, tt.args.m); got != tt.want {
				t.Errorf("rearageStacks9001() = %v, want %v", got, tt.want)
			}
		})
	}
}
