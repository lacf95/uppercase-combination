package uppercase_combination

import (
	"reflect"
	"testing"
)

func TestUppercaseCombination(t *testing.T) {
	type test struct {
		input string
		want []string
	}

	tests := []test{
		{"a", []string{"a", "A"}},
		{"ab", []string{"ab", "aB", "Ab", "AB"}},
		{"abc", []string{"abc", "abC", "aBc", "aBC", "Abc", "AbC", "ABc", "ABC"}},
	}

	for _, tc := range tests {
		got := UppercaseCombination(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func BenchmarkUppercaseCombination(b *testing.B) {
	type test struct {
		input string
	}

	t := test{"abcdefghijklmnopqrstu"}

	UppercaseCombination(t.input)
}
