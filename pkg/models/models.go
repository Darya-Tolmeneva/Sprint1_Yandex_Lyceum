package models

import "errors"

type Request struct {
	Expression string `json:"expression"`
}
type Response struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

var (
	ErrInvalidSymbols         = errors.New("invalid symbols")
	ErrInvalidParenthesis     = errors.New("invalid parenthesis")
	ErrInvalidOperations      = errors.New("invalid operations")
	ErrNotEnoughOperands      = errors.New("there are not enough operands in the expression")
	ErrDivisionByZero         = errors.New("division by zero")
	ErrExpressionNotEvaluated = errors.New("the expression is not evaluated")
)
