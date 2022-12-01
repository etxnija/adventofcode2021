package day1

import (
	"testing"
)

func Test_findElfWithMostCalories(t *testing.T) {
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
			want: 24000,
		},
		{
			name: "problem",
			args: args{
				input: "input.txt",
			},
			want: 70613,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findElfWithMostCalories(tt.args.input); got != tt.want {
				t.Errorf("findElfWithMostCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeTotalElfWithMostCalories(t *testing.T) {
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
			want: 45000,
		},
		{
			name: "problem",
			args: args{
				input: "input.txt",
			},
			want: 70613,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeTotalElfWithMostCalories(tt.args.input); got != tt.want {
				t.Errorf("threeTotalElfWithMostCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}
