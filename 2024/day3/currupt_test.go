package day3

import (
	"reflect"
	"testing"
)

func Test_calculate(t *testing.T) {
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
			want: 161,
		},
		{
			name: "input",
			args: args{
				file: "input.txt",
			},
			want: 167090022,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.file); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
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
		want    [][]int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "basic",
			args: args{
				row: "mul(2,2)",
			},
			want:    [][]int{{2, 2}},
			wantErr: false,
		},
		{
			name: "basic with junk after",
			args: args{
				row: "mul(2,2)_sdf fdf",
			},
			want:    [][]int{{2, 2}},
			wantErr: false,
		},
		{
			name: "basic with junk before and after",
			args: args{
				row: "(((mul(2,2)_sdf fdf",
			},
			want:    [][]int{{2, 2}},
			wantErr: false,
		},
		{
			name: "three matches",
			args: args{
				row: "(((mul(2,2)_sdf fdfmul(3,4)mul(445,5)",
			},
			want:    [][]int{{2, 2}, {3, 3}, {445, 445}},
			wantErr: false,
		},
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

func Test_calculate2(t *testing.T) {
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
			name: "example2",
			args: args{
				file: "example2.txt",
			},
			want: 96,
		},
		{
			name: "input",
			args: args{
				file: "input.txt",
			},
			want: 89823704,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate2(tt.args.file); got != tt.want {
				t.Errorf("calculate2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterRow(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "no do at start",
			args: args{
				row: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			want: "xmul(2,4)&mul[3,7]!^?mul(8,5))",
		},
		{
			name: "no dont at start",
			args: args{
				row: "xmul(2,4)&mul[3,7]!^don't_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			want: "xmul(2,4)&mul[3,7]!^don't_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
		},
		{
			name: "no dont",
			args: args{
				row: "xmul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			want: "xmul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
		},
		{
			name: "one dont first,empty return",
			args: args{
				row: "don't()xmul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64](mul(11,8)undo?mul(8,5))",
			},
			want: "",
		},
		{
			name: "one dont last",
			args: args{
				row: "xmul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64](mul(11,8)undo?mul(8,5))don't()",
			},
			want: "xmul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64](mul(11,8)undo?mul(8,5))",
		},
		{
			name: "one dont middle, no do after",
			args: args{
				row: "xmul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64]_don't()_(mul(11,8)undo?mul(8,5))",
			},
			want: "xmul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64]_",
		},
		{
			name: "two don'ts",
			args: args{
				row: "xmul(2,4)&mul[3,7]!don't()^_mul(5,5)+mul(32,64]_don't()_(mul(11,8)undo?mul(8,5))",
			},
			want: "xmul(2,4)&mul[3,7]!",
		},
		{
			name: "two don'ts",
			args: args{
				row: "xmul(2,4)&mul[3,7]!don't()^_mul(5,5)+mul(32,64]_do()_(mul(11,8)mul(630,973)#+how()who()+do()undo?mul(8,5))don't()_(mul(11,8)undo?mul(8,5))undo()?mul(8,5))",
			},
			want: "xmul(2,4)&mul[3,7]!_(mul(11,8)mul(630,973)#+how()who()+do()undo?mul(8,5))?mul(8,5))",
		},
		{
			name: "who()",
			args: args{
				row: ";@who()?mul(607,783)<+don't()what()mul(78,875)@:when()~mul(665,145)[)][>@<{mul(469,481)why()select(44,921)>who()>:@[}mul(762,260)%mul(155,860)[mul(886,453))m",
			},
			want: ";@who()?mul(607,783)<+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterRow(tt.args.row); got != tt.want {
				t.Errorf("%s:  filterRow() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
