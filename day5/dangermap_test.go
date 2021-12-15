package day5

import (
	"reflect"
	"testing"
)

func Test_danagerMap(t *testing.T) {
	type args struct {
		wents []line
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy",
			args: args{
				wents: readWents("example.txt"),
			},
			want: 5,
		},
		{
			name: "Problem",
			args: args{
				wents: readWents("input.txt"),
			},
			want: 5,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := danagerMap(tt.args.wents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("danagerMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_line_points(t *testing.T) {
	type fields struct {
		start point
		end   point
	}
	tests := []struct {
		name   string
		fields fields
		want   []point
	}{
		{
			name: "horizontal_left",
			fields: fields{
				start: point{4, 4},
				end:   point{0, 4},
			},
			want: []point{{4, 4}, {3, 4}, {2, 4}, {1, 4}, {0, 4}},
		},
		{
			name: "horizontal_right",
			fields: fields{
				start: point{0, 4},
				end:   point{4, 4},
			},
			want: []point{{0, 4}, {1, 4}, {2, 4}, {3, 4}, {4, 4}},
		},
		{
			name: "Vertical_down",
			fields: fields{
				start: point{4, 4},
				end:   point{4, 0},
			},
			want: []point{{4, 4}, {4, 3}, {4, 2}, {4, 1}, {4, 0}},
		},
		{
			name: "Vertical_up",
			fields: fields{
				start: point{4, 0},
				end:   point{4, 4},
			},
			want: []point{{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}},
		},
		{
			name: "Down_left",
			fields: fields{
				start: point{4, 4},
				end:   point{0, 0},
			},
			want: []point{{4, 4}, {3, 3}, {2, 2}, {1, 1}, {0, 0}},
		},
		{
			name: "Down_right",
			fields: fields{
				start: point{4, 4},
				end:   point{8, 0},
			},
			want: []point{{4, 4}, {5, 3}, {6, 2}, {7, 1}, {8, 0}},
		},
		{
			name: "Up_left",
			fields: fields{
				start: point{4, 4},
				end:   point{0, 8},
			},
			want: []point{{4, 4}, {3, 5}, {2, 6}, {1, 7}, {0, 8}},
		},
		{
			name: "Up_Right",
			fields: fields{
				start: point{4, 4},
				end:   point{8, 8},
			},
			want: []point{{4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := line{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := l.points(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("line.points() = %v, want %v", got, tt.want)
			}
		})
	}
}
