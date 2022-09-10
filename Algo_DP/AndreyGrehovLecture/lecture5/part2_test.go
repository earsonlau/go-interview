package lecture5

import "testing"

func Test_climbStairs3steps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple case #1",
			args: args{
				n: 3,
			},
			want: 4,
		},
		{
			name: "simple case #2",
			args: args{
				n: 5,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs3steps(tt.args.n); got != tt.want {
				t.Errorf("climbStairs3steps() = %v, want %v", got, tt.want)
			}

		})
	}
}