package day3

import (
	"reflect"
	"testing"
)

func TestSumOfpriorities(t *testing.T) {
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
			want: 157,
		},
		{
			name: "input",
			args: args{
				input: "input.txt",
			},
			want: 7845,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfpriorities(tt.args.input); got != tt.want {
				t.Errorf("SumOfpriorities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCommon(t *testing.T) {
	type args struct {
		one string
		two string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "happy",
			args: args{
				one: "vJrwpWtwJgWr",
				two: "hcsFMMfFFhFp",
			},
			want: 'p',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCommon(tt.args.one, tt.args.two); got != tt.want {
				t.Errorf("findCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPrioritie(t *testing.T) {
	type args struct {
		common rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy",
			args: args{
				common: 'p',
			},
			want: 16,
		},
		{
			name: "capital",
			args: args{
				common: 'L',
			},
			want: 38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPrioritie(tt.args.common); got != tt.want {
				t.Errorf("findPrioritie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfBadges(t *testing.T) {
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
			want: 70,
		},
		{
			name: "input",
			args: args{
				input: "input.txt",
			},
			want: 2790,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfBadges(tt.args.input); got != tt.want {
				t.Errorf("SumOfBadges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPotential(t *testing.T) {
	type args struct {
		one string
		two string
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "find potential",
			args: args{
				one: "vJrwpW",
				two: "jqHrNJ",
			},
			want: []rune{'J', 'r'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPotential(tt.args.one, tt.args.two); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPotential() = %v, want %v", got, tt.want)
			}
		})
	}
}
