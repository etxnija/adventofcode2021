package day4

import (
	"testing"
)

func Test_countContained(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				input: "example.txt",
			},
			want: 2,
		},
		{
			name: "input",
			args: args{
				input: "input.txt",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countContained(tt.args.input); got != tt.want {
				t.Errorf("countContained() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oneContained(t *testing.T) {
	type args struct {
		pair []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains",
			args: args{
				pair: []string{
					"2-8",
					"3-7",
				},
			},
			want: true,
		},
		{
			name: "contains second langer",
			args: args{
				pair: []string{
					"6-6",
					"4-6",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneContained(tt.args.pair); got != tt.want {
				t.Errorf("oneContained() = %v, want %v", got, tt.want)
			}
		})
	}
}
