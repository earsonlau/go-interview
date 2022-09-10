package lecture8

import 
("testing"
"reflect")

func Test_climbPaidStairs_best_path_v2(t *testing.T) {
	type args struct {
		n int
		paid []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "simple case #1",
			args: args{
				n: 3,
				paid: []int{0,3,2,4},
			},
			want: []int{0,2,3},
		},
		{
			name: "simple case #2",
			args: args{
				n: 4,
				paid: []int{0,2,3,4,5},
			},
			want: []int{0,2,4},
		},
		{
			// 0   3   2   4   6   1   1   5   3
			// |---|---|---|---|---|---|---|---|
			// 0   1   2   3   4   5   6   7   8
			//
			// 0 -> 2 -> 3 -> 5 -> 6 -> 8 = 11
			name: "complex case",
			args: args{
				n: 8,
				paid: []int{0, 3, 2, 4, 6, 1, 1, 5, 3},
			},
			want: []int{0, 2, 3, 5, 6, 8},
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbPaidStairs_best_path_v2(tt.args.n,tt.args.paid); !reflect.DeepEqual(got,tt.want) {
				t.Errorf("climbPaidStairs_best_path_v2() = %v, want %v", got, tt.want)
			}

		})
	}
}