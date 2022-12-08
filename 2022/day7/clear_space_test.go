package day7

import (
	"reflect"
	"testing"
)

func Test_flatTree(t *testing.T) {
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
			want: 95437,
		},
		{
			name: "input",
			args: args{
				input: "input.txt",
			},
			want: 1644735,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatTree(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_locationPath(t *testing.T) {
	type args struct {
		l []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "root",
			args: args{
				l: []string{"/"},
			},
			want: "root/",
		},
		{
			name: "one level",
			args: args{
				l: []string{"/", "a"},
			},
			want: "root/a/",
		},
		{
			name: "two level",
			args: args{
				l: []string{"/", "a", "b"},
			},
			want: "root/a/b/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := locationPath(tt.args.l); got != tt.want {
				t.Errorf("locationPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileToRemove(t *testing.T) {
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
			want: 24933642,
		},
		{
			name: "input",
			args: args{
				input: "input.txt",
			},
			want: 1644735,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileToRemove(tt.args.input); got != tt.want {
				t.Errorf("fileToRemove() = %v, want %v", got, tt.want)
			}
		})
	}
}
