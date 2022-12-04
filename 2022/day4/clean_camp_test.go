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
			want: 556,
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

func Test_opverlap(t *testing.T) {
	type args struct {
		pair []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one overlap right",
			args: args{
				pair: []string{
					"5-7",
					"7-9",
				},
			},
			want: true,
		},
		{
			name: "No overlap",
			args: args{
				pair: []string{
					"2-3",
					"4-5",
				},
			},
			want: false,
		},
		{
			name: "overlap left",
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
			if got := opverlap(tt.args.pair); got != tt.want {
				t.Errorf("opverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countOverlap(t *testing.T) {
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
			want: 4,
		},
		{
			name: "input",
			args: args{
				input: "input.txt",
			},
			want: 876,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOverlap(tt.args.input); got != tt.want {
				t.Errorf("countOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
