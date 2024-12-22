package parser

import (
	"reflect"
	"testing"
)

func TestParseToPostfix(t *testing.T) {
	tests := []struct {
		expression string
		expected   []string
	}{
		{"1+2", []string{"1", "2", "+"}},
		{"(1+2)*3", []string{"1", "2", "+", "3", "*"}},
		{"1+2*3", []string{"1", "2", "3", "*", "+"}},
		{"(1+2)*(3+4)", []string{"1", "2", "+", "3", "4", "+", "*"}},
	}
	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			result := ParseToPostfix(test.expression)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %v, expected %v", result, test.expected)
			}
		})
	}
}
func TestEvaluatePostfix(t *testing.T) {
	tests := []struct {
		postfix   []string
		expected  float64
		shouldErr bool
	}{
		{[]string{"1", "2", "+", "3", "*"}, 9, false},
		{[]string{"1", "2", "*", "3", "+"}, 5, false},
		{[]string{"1", "0", "/"}, 0, true},
		{[]string{"1", "2", "+"}, 3, false},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result, err := EvaluatePostfix(test.postfix)
			if test.shouldErr && err == nil {
				t.Errorf("expected error, got result %v", result)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("expected result %v, got error %v", test.expected, err)
			}
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
