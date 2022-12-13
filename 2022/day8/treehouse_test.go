package day8

import (
	"testing"
)

func Test_visableTrees(t *testing.T) {
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
			want: 21,
		},
		{
			name: "input",
			args: args{
				input: "input.txt",
			},
			want: 1825,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := visableTrees(tt.args.input); got != tt.want {
				t.Errorf("visableTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anyTaller(t *testing.T) {
	type args struct {
		lineOfSight []int
		tree        int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "One taller",
			args: args{
				lineOfSight: []int{2, 3, 6, 3, 4},
				tree:        5,
			},
			want: true,
		},
		{
			name: "none taller",
			args: args{
				lineOfSight: []int{2, 3, 4, 3, 4},
				tree:        5,
			},
			want: false,
		},
		{
			name: "one same",
			args: args{
				lineOfSight: []int{2, 3, 4, 5, 4},
				tree:        5,
			},
			want: true,
		},
		{
			name: "two same",
			args: args{
				lineOfSight: []int{2, 5, 4, 5, 4},
				tree:        5,
			},
			want: true,
		},
		{
			name: "first taller",
			args: args{
				lineOfSight: []int{6, 3, 4, 3, 4},
				tree:        5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anyTaller(tt.args.lineOfSight, tt.args.tree); got != tt.want {
				t.Errorf("anyTaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
