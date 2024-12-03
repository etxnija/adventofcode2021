package day2

import (
	"testing"
)

func Test_safeReports(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "example",
			args: args{
				file: "example.txt",
			},
			want: 2,
		},
		{
			name: "input",
			args: args{
				file: "input.txt",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := safeReports(tt.args.file); got != tt.want {
				t.Errorf("safeReports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allSameDirection(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "all increasing",
			args: args{
				values: []int{7, 6, 4, 2, 1},
			},
			want: true,
		},
		{
			name: "all decreasing",
			args: args{
				values: []int{7, 6, 4, 2, 2, 1},
			},
			want: true,
		},
		{
			name: "both directions",
			args: args{
				values: []int{24, 22, 23, 27, 28, 29, 32, 35},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allSameDirection(tt.args.values); got != tt.want {
				t.Errorf("allSameDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_withinAdjacent(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "all within",
			args: args{
				values: []int{7, 6, 4, 2, 1},
			},
			want: true,
		},
		{
			name: "Two same",
			args: args{
				values: []int{7, 6, 4, 2, 2, 1},
			},
			want: false,
		},
		{
			name: "One over 3",
			args: args{
				values: []int{24, 22, 23, 27, 28, 29, 32, 35},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := withinAdjacent(tt.args.values); got != tt.want {
				t.Errorf("withinAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}
