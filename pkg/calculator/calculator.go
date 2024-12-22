package calculator

import (
	"github.com/Darya-Tolmeneva/Sprint1_Yandex_Lyceum/pkg/checker"
	"github.com/Darya-Tolmeneva/Sprint1_Yandex_Lyceum/pkg/parser"
	"github.com/Darya-Tolmeneva/Sprint1_Yandex_Lyceum/pkg/validator"
)

// Calc принимает строковое математическое выражение, проверяет его корректность и возвращает результат вычисления.
// В случае некорректного выражения возвращает ошибку.
func Calc(expression string) (float64, error) {
	ch := checker.ExpressionChecker{}
	if err := validator.ValidateExpression(ch, expression); err != nil {
		return 0, err
	}

	postfix := parser.ParseToPostfix(expression)
	result, err := parser.EvaluatePostfix(postfix)

	return result, err
}
