package parser

import (
	"Sprint1/internal/checker"
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func priority(operator rune) int {
	switch operator {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func operations(a, b float64, op rune) float64 {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}

// ParseToPostfix преобразует инфиксное выражение в постфиксное (обратную польскую нотацию).
// Такой метод упрощает вычисление выражения, так как в постфиксной нотации операции выполняются
// сразу над последними двумя операндами в порядке приоритета.
// Возвращает массив строк, представляющий постфиксное выражение.
func ParseToPostfix(expression string) []string {
	var numbers []string
	var operators []rune
	var currentNum strings.Builder

	for _, char := range expression {
		if unicode.IsDigit(char) || char == '.' {
			currentNum.WriteRune(char)
		} else {
			if currentNum.Len() > 0 {
				numbers = append(numbers, currentNum.String())
				currentNum.Reset()
			}
			if char == '(' {
				operators = append(operators, '(')
			} else if char == ')' {
				for len(operators) > 0 && operators[len(operators)-1] != '(' {
					numbers = append(numbers, string(operators[len(operators)-1]))
					operators = operators[:len(operators)-1]
				}
				if len(operators) > 0 && operators[len(operators)-1] == '(' {
					operators = operators[:len(operators)-1]
				}
			} else if checker.IsOperator(char) {
				op := char
				for len(operators) > 0 && priority(operators[len(operators)-1]) >= priority(op) {
					numbers = append(numbers, string(operators[len(operators)-1]))
					operators = operators[:len(operators)-1]
				}
				operators = append(operators, op)
			}
		}
	}
	if currentNum.Len() > 0 {
		numbers = append(numbers, currentNum.String())
	}

	for len(operators) > 0 {
		numbers = append(numbers, string(operators[len(operators)-1]))
		operators = operators[:len(operators)-1]
	}

	return numbers
}

// EvaluatePostfix вычисляет результат постфиксного выражения.
// Принимает массив строк, представляющих постфиксное выражение, и возвращает вычисленный результат.
// В случае деления на ноль возвращает ошибку.
func EvaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else if len(token) == 1 && checker.IsOperator(rune(token[0])) {
			if len(stack) < 2 {
				return 0, errors.New("недостаточно операндов в выражении")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			if rune(token[0]) == '/' && b == 0 {
				return 0, errors.New("деление на ноль")
			}

			result := operations(a, b, rune(token[0]))
			stack = append(stack, result)
		}
	}

	if len(stack) == 0 {
		return 0, errors.New("выражение не вычислено")
	}
	return stack[0], nil
}
