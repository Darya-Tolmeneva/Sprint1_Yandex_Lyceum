package validator

import (
	"github.com/Darya-Tolmeneva/Sprint1_Yandex_Lyceum/pkg/checker"
	"github.com/Darya-Tolmeneva/Sprint1_Yandex_Lyceum/pkg/models"
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
