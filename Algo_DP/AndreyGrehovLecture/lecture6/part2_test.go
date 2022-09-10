package lecture6

import "testing"

func Test_climbStairsksteps_so_sr(t *testing.T) {
	type args struct {
		n int
		k int
		stairs []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple case #1",
			args: args{
				n: 7,
				k: 3,
				stairs: []bool{false, true, false, true, true, false, false},
			},
			want: 2,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsksteps_so_sr(tt.args.n,tt.args.k,tt.args.stairs); got != tt.want {
				t.Errorf("climbStairsksteps_so_sr() = %v, want %v", got, tt.want)
			}

		})
	}
}