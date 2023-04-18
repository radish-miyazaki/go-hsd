package hsd

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringDistance(t *testing.T) {
	tests := []struct {
		name     string
		lhs, rhs string
		want     int
	}{
		{name: "lhs is longer than rhs", lhs: "abc", rhs: "ab", want: -1},
		{name: "lhs is shorter than rhs", lhs: "ab", rhs: "abc", want: -1},
		{name: "one different character", lhs: "abc", rhs: "abf", want: 1},
		{name: "two different characters", lhs: "abc", rhs: "aef", want: 2},
		{name: "three different characters", lhs: "abc", rhs: "def", want: 3},
		{name: "multibyte characters", lhs: "あいう", rhs: "あうえ", want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringDistance(tt.lhs, tt.rhs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDistance(%q, %q) = %d, want %d", tt.lhs, tt.rhs, got, tt.want)
			}
		})
	}
}

func ExampleStringDistance() {
	lhs := "abc"
	rhs := "abf"
	dist := StringDistance(lhs, rhs)

	fmt.Println(dist)
	// Output: 1
}
