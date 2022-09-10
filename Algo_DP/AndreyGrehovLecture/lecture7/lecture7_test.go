package lecture7

import "testing"

func Test_climbPaidStairs(t *testing.T) {
	type args struct {
		n int
		paid []int
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
				paid: []int{0,3,2,4},
			},
			want: 6,
		},
		{
			name: "simple case #2",
			args: args{
				n: 4,
				paid: []int{0,2,3,4,5},
			},
			want: 8,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbPaidStairs(tt.args.n,tt.args.paid); got != tt.want {
				t.Errorf("climbPaidStairs() = %v, want %v", got, tt.want)
			}

		})
	}
}