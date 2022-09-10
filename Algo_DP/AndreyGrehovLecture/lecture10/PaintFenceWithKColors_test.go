package lecture10

import (
	"testing"
)

func Test_paintFenceWithKColors(t *testing.T) {
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
			name: "test case #1",
			args: args{
				n: 5,
				k: 3,
			},
			want: 30,
		},
		{
			name: "test case #2",
			args: args{
				n: 3,
				k: 2,
			},
			want: 0,
		},
		
	}	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paintFenceWithKColors(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("paintFenceWithKColors() = %v, want %v", got, tt.want)
			}

		})
	}
}