package tests

import (
	"Sprint1/internal/checker"
	"testing"
)

func TestCheckParanthesis(t *testing.T) {
	ch := checker.ExpressionChecker{}
	tests := []struct {
		expression string
		expected   bool
	}{
		{"(1+2)", true},
		{"((1+2)", false},
		{"(1+2))", false},
		{"(1+2)*(3+4)", true},
	}
	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			result := ch.CheckParenthesis(test.expression)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestCheckOperations(t *testing.T) {
	ch := checker.ExpressionChecker{}
	tests := []struct {
		expression string
		expected   bool
	}{
		{"1+2", true},
		{"1++2", false},
		{"1+*2", false},
		{"1+2*", false},
		{"1+2-3", true},
	}
	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			result := ch.CheckOperations(test.expression)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
func TestCheckSymbols(t *testing.T) {
	ch := checker.ExpressionChecker{}
	tests := []struct {
		expression string
		expected   bool
	}{
		{"1+2", true},
		{"1+a", false},
		{"1+2*", true},
		{"1+2#", false},
	}
	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			result := ch.CheckSymbols(test.expression)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
