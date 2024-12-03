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
			want: 432,
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

func Test_safeReports2(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			name: "example",
			args: args{
				file: "example.txt",
			},
			want: 4,
		},
		{
			name: "input",
			args: args{
				file: "input.txt",
			},
			want: 488,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := safeReports2(tt.args.file); got != tt.want {
				t.Errorf("safeReports2() = %v, want %v", got, tt.want)
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
				t.Errorf("name : %s: allSameDirection() = %v, want %v", tt.name, got, tt.want)
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

func Test_allSameDirectionAndAdjacentWithDamper(t *testing.T) {
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
				values: []int{7, 6, 4, 2, 1},
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
			want: true,
		},
		{
			name: "One over 3",
			args: args{
				values: []int{24, 22, 23, 27, 28, 29, 32, 35},
			},
			want: false,
		},
		{
			name: "Safe with damper",
			args: args{
				values: []int{1, 3, 2, 4, 5},
			},
			want: true,
		},
		{
			name: "Safe with damper, remove last",
			args: args{
				values: []int{1, 3, 4, 5, 12},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allSameDirectionAndAdjacentWithDamper(tt.args.values); got != tt.want {
				t.Errorf("test: %s: allSameDirectionAndAdjacentWithDamper() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
