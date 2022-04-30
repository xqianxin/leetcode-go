package sort

import (
	"reflect"
	"testing"
)

func Test_pancakeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{arr: []int{3, 2, 4, 1}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pancakeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pancakeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
