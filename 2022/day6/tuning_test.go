package day6

import (
	"io"
	"strings"
	"testing"
)

func Test_startOfPacketMarker(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happuy 1",
			args: args{
				reader: strings.NewReader("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			},
			want: 5,
		},
		{
			name: "happuy 2",
			args: args{
				reader: strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg"),
			},
			want: 6,
		},
		{
			name: "happuy 3",
			args: args{
				reader: strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			},
			want: 10,
		},
		{
			name: "happuy 4",
			args: args{
				reader: strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startOfPacketMarker(tt.args.reader); got != tt.want {
				t.Errorf("startOfPacketMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMarker(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input",
			args: args{
				input: "input.txt",
			},
			want: 1766,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMarker(tt.args.input); got != tt.want {
				t.Errorf("findMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
