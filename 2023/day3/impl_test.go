package day3

import (
	"math"
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example",
			args: args{
				in: "example.txt",
			},
			want:    4361,
			wantErr: false,
		},
		{
			name: "example",
			args: args{
				in: "input.txt",
			},
			want:    4361,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vectorDist(t *testing.T) {
	type args struct {
		p1 []int
		p2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "adjacent",
			args: args{
				p1: []int{0, 2},
				p2: []int{1, 3},
			},
			want: math.Sqrt(2),
		},
		{
			name: "same",
			args: args{
				p1: []int{1, 3},
				p2: []int{1, 3},
			},
			want: 0,
		},
		{
			name: "one row away",
			args: args{
				p1: []int{0, 3},
				p2: []int{1, 3},
			},
			want: 1,
		},
		{
			name: "one col away",
			args: args{
				p1: []int{1, 2},
				p2: []int{1, 3},
			},
			want: 1,
		},
		{
			name: "two col away",
			args: args{
				p1: []int{1, 1},
				p2: []int{1, 3},
			},
			want: 2,
		},
		{
			name: "further away",
			args: args{
				p1: []int{1, 1},
				p2: []int{5, 7},
			},
			want: math.Sqrt(52),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vectorDist(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("vectorDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processRow(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]int
		wantErr bool
	}{
		{
			name: "one",
			args: args{
				row: ".......*............680.....*......876.........864..................259.................124.169....799............608..*.........98......951",
			},
			want: make(map[string][]int),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processRow(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("processRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve2(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example",
			args: args{
				in: "example.txt",
			},
			want:    467835,
			wantErr: false,
		},
		{
			name: "input",
			args: args{
				in: "input.txt",
			},
			want:    467835,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve2(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("solve2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solve2() = %v, want %v", got, tt.want)
			}
		})
	}
}
