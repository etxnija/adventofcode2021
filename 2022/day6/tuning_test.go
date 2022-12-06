package day6

import (
	"io"
	"strings"
	"testing"
)

func Test_startOfPacketMarker(t *testing.T) {
	type args struct {
		reader io.Reader
		length int
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
				length: 4,
			},
			want: 5,
		},
		{
			name: "happuy 2",
			args: args{
				reader: strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg"),
				length: 4,
			},
			want: 6,
		},
		{
			name: "happuy 3",
			args: args{
				reader: strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
				length: 4,
			},
			want: 10,
		},
		{
			name: "happuy 4",
			args: args{
				reader: strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
				length: 4,
			},
			want: 11,
		},
		{
			name: "happy 14 1",
			args: args{
				reader: strings.NewReader("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
				length: 14,
			},
			want: 19,
		},
		{
			name: "happy 14 2",
			args: args{
				reader: strings.NewReader("bvwbjplbgvbhsrlpgdmjqwftvncz"),
				length: 14,
			},
			want: 23,
		},
		{
			name: "happy 14 3",
			args: args{
				reader: strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg"),
				length: 14,
			},
			want: 23,
		},
		{
			name: "happy 14 3",
			args: args{
				reader: strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
				length: 14,
			},
			want: 29,
		},
		{
			name: "happy 14 3",
			args: args{
				reader: strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
				length: 14,
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startOfPacketMarker(tt.args.reader, tt.args.length); got != tt.want {
				t.Errorf("startOfPacketMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMarker(t *testing.T) {
	type args struct {
		input  string
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input",
			args: args{
				input:  "input.txt",
				length: 4,
			},
			want: 1766,
		},
		{
			name: "input",
			args: args{
				input:  "input.txt",
				length: 14,
			},
			want: 2383,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMarker(tt.args.input, tt.args.length); got != tt.want {
				t.Errorf("findMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
