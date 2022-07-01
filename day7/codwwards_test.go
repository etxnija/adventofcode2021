package day7

import (
	"reflect"
	"testing"
)

func TestTwoOldestAges(t *testing.T) {
	type args struct {
		ages []int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{
			name: "happy",
			args: args{
				ages: []int{6, 5, 83, 5, 3, 18},
			},
			want: [2]int{18, 83},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoOldestAges(tt.args.ages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoOldestAges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArray(t *testing.T) {
	type args struct {
		array1 []string
		array2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "One",
			args: args{
				array1: []string{"live", "arp", "strong"},
				array2: []string{"lively", "alive", "harp", "sharp", "armstrong"},
			},
			want: []string{"arp", "live", "strong"},
		},
		{
			name: "two",
			args: args{
				array1: []string{"cod", "code", "wars", "ewar", "ar"},
				array2: []string{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOutlier(t *testing.T) {
	type args struct {
		integers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				integers: []int{2, 6, 8, -10, 3},
			},
			want: 3,
		},
		{
			name: "two",
			args: args{
				integers: []int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781},
			},
			want: 206847684,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOutlier(tt.args.integers); got != tt.want {
				t.Errorf("FindOutlier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHighAndLow(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				in: "1 2 3 4 5",
			},
			want: "5 1",
		},
		{
			name: "two",
			args: args{
				in: "8 3 -5 42 -1 0 0 -9 4 7 4 -4",
			},
			want: "42 -9",
		},
		{
			name: "three",
			args: args{
				in: "1 2 3",
			},
			want: "3 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HighAndLow(tt.args.in); got != tt.want {
				t.Errorf("HighAndLow() = %v, want %v", got, tt.want)
			}
		})
	}
}
