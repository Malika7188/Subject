package main

import (
	"log"
	"reflect"
)

func CanJump(num []uint) bool {
	if len(num) == 0 {
		return false
	}
	if len(num) == 1 {
		return true
	}
	i := 0

	for i < len(num)-1 {
		steps := int(num[i])
		if steps == 0 {
			return false
		}
		i += steps
	}
	return i == len(num)-1
}
func main() {
	tests := []struct {
		args []uint
		want bool
	}{
		{args: []uint{2, 3, 1, 1, 4}, want: true},
		{args: []uint{1, 1, 1, 1, 0}, want: true},
		{args: []uint{5, 4, 3, 2, 1, 0}, want: true},
		{args: []uint{0}, want: true},
		{args: []uint{5}, want: true},
		{args: []uint{}, want: false},
		{args: []uint{1, 2, 3, 0, 2}, want: false},
		{args: []uint{3, 2, 1, 0, 4}, want: false},
		{args: []uint{0, 0, 0, 0, 0}, want: false},
		{args: []uint{1, 2, 3}, want: false},
		{args: []uint{1, 2, 3, 0, 1}, want: false},
		{args: []uint{1, 0, 0, 0, 0}, want: false},
		{args: []uint{1, 0, 1, 0, 1}, want: false},
		{args: []uint{10, 20, 30, 40, 0}, want: false},
	}

	for _, tc := range tests {
		got := CanJump(tc.args)
		if !reflect.DeepEqual(got, tc.want) {
			log.Fatalf("%s(%+v) == %+v instead of %+v\n",
				"CanJump",
				tc.args,
				got,
				tc.want,
			)
		}
	}
}
