package lecture6

import "testing"

func Test_climbStairsksteps_so(t *testing.T) {
	type args struct {
		n int
		k int
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
				k: 2,
			},
			want: 3,
		},
		{
			name: "simple case #2",
			args: args{
				n: 5,
				k: 2,
			},
			want: 8,
		},
		{
			name: "simple case #3",
			args: args{
				n: 3,
				k: 3,
			},
			want: 4,
		},
		{
			name: "simple case #4",
			args: args{
				n: 5,
				k: 3,
			},
			want: 13,
		},
		{
			name: "n bignum case #1",
			args: args{
				n: 1000000,
				k: 2,
			},
			want: 2756670985995446685,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsksteps_so(tt.args.n,tt.args.k); got != tt.want {
				t.Errorf("climbStairsksteps_so() = %v, want %v", got, tt.want)
			}

		})
	}
}