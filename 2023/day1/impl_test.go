package day1

import "testing"

func Test_calibrate(t *testing.T) {
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
			want:    142,
			wantErr: false,
		},
		{
			name: "input",
			args: args{
				in: "input.txt",
			},
			want:    55358,
			wantErr: false,
		},
		{
			name: "example2",
			args: args{
				in: "example2.txt",
			},
			want:    281,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calibrate(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("calibrate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calibrate() = %v, want %v", got, tt.want)
			}
		})
	}
}
