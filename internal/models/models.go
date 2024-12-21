package models

import "errors"

type Request struct {
	Expression string `json:"expression"`
}
type Response struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

var ErrInvalidSymbols = errors.New("invalid symbols")
var ErrInvalidParenthesis = errors.New("invalid parenthesis")
var ErrInvalidOperations = errors.New("invalid operations")
