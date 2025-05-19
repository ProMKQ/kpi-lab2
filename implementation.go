package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

type stack []string

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func (s *stack) Pop() string {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

// PostfixToInfix converts an expression from Reverse Polish (postfix) notation to infix notation
func PostfixToInfix(input string) (string, error) {
	tokens := strings.Fields(input)
	s := make(stack, 0)

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
			if len(s) < 2 {
				return "", fmt.Errorf("expected at least 2 tokens, got %d", len(s))
			}
			right := s.Pop()
			left := s.Pop()
			s.Push("(" + right + " " + token + " " + left + ")")
		} else if _, err := strconv.Atoi(token); err == nil {
			s.Push(token)
		} else {
			return "", fmt.Errorf("invalid token: '%s'", token)
		}
	}

	if len(s) != 1 {
		return "", fmt.Errorf("invalid postfix expression")
	}

	return s[0], nil
}
