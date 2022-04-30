package dynamicpro

import "testing"

func TestNthSuperUglyNumber(t *testing.T) {
	type args struct {
		n      int
		primes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{12, []int{2, 7, 13, 19}},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NthSuperUglyNumber(tt.args.n, tt.args.primes); got != tt.want {
				t.Errorf("NthSuperUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
