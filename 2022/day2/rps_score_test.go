package day2

import (
	"testing"
)

func TestCalculateRPSScore(t *testing.T) {
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
			want: 15,
		},
		{
			name: "input",
			args: args{
				input: "input.txt",
			},
			want: 12794,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateRPSScore(tt.args.input); got != tt.want {
				t.Errorf("CalculateRPSScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRPSScoreStratergy2(t *testing.T) {
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
			want: 12,
		},
		{
			name: "input",
			args: args{
				input: "input.txt",
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateRPSSscoreStratergy2(tt.args.input); got != tt.want {
				t.Errorf("CalculateRPSScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
