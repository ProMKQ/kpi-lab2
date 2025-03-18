package lab2

import (
	"fmt"
	"testing"
)

func TestPostfixToInfix(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"4", "4", false},
		{"2 2 +", "(2 + 2)", false},
		{"10 3 -", "(10 - 3)", false},
		{"3  4  *", "(3 * 4)", false},
		{" 10 5 / ", "(10 / 5)", false},
		{"2 3 + 4 ^", "((2 + 3) ^ 4)", false},

		{"2 3 + 4 *", "((2 + 3) * 4)", false},
		{"5 1 2 + 0 - 4 3 / * + 3 3 ^ 2 ^ -", "((5 + (((1 + 2) - 0) * (4 / 3))) - ((3 ^ 3) ^ 2))", false},
		{"17 6 * 5 * 4 * 3 * 2 3 * - 45 +", "((((((17 * 6) * 5) * 4) * 3) - (2 * 3)) + 45)", false},
		{"10 2 3 * + 4 5 * 6 * + 33 - 4 -", "((((10 + (2 * 3)) + ((4 * 5) * 6)) - 33) - 4)", false},

		{"", "", true},
		{"+ 2 2", "", true},
		{"2 +", "", true},
		{"2 3 &", "", true},
		{"4 5", "", true},
	}

	for _, test := range tests {
		res, err := PostfixToInfix(test.input)
		if test.err {
			if err == nil {
				t.Errorf("expected error for input %q, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error from input %q: %v", test.input, err)
			} else if res != test.expected {
				t.Errorf("for input %q expected %q, but got %q", test.input, test.expected, res)
			}
		}
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("4 2 - 3 2 ^ * 5 +")
	fmt.Println(res)

	// Output:
	// (((4 - 2) * (3 ^ 2)) + 5)
}
