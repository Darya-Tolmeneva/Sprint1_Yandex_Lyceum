package validator

import (
	"Sprint1/internal/checker"
	"Sprint1/internal/models"
)

func ValidateExpression(checker checker.Checker, expression string) error {
	if !checker.CheckParenthesis(expression) {
		return models.ErrInvalidParenthesis
	}
	if !checker.CheckOperations(expression) {
		return models.ErrInvalidOperations
	}
	if !checker.CheckSymbols(expression) {
		return models.ErrInvalidSymbols
	}
	return nil
}
