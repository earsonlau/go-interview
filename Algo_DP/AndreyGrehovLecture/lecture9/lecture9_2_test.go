package lecture9

import "testing"

func Test_uniquePathswithObstacles(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "small matrix",
			args: args{ grid: [][]int{
				{0,0},
			}},
			want: 1,
		},
		{
			name: "simple matrix",
			args: args{grid: [][]int{
				{0, 0, 0, 0},
				{0, 0, 1, 1},
				{0, 0, 0, 0},
			}},
			want: 3,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathswithObstacles(tt.args.grid); got != tt.want {
				t.Errorf("uniquePathswithObstacles() = %v, want %v", got, tt.want)
			}

		})
	}
}