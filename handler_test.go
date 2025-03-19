package lab2

import (
	"bytes"
	"strings"
	"testing"
)

func TestCompute_Success(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Simple Expression", "3 4 +", "(3 + 4)\n"},
		{"Expression with multiple operators", "3 4 + 2 *", "((3 + 4) * 2)\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			handler := ComputeHandler{input, &output}

			if err := handler.Compute(); err != nil {
				t.Errorf("Unexpected Compute() error: '%v'", err)
			}

			res := strings.TrimSpace(output.String())
			expected := strings.TrimSpace(tt.expected)
			if res != expected {
				t.Errorf("Compute() = '%v', expected '%v'", res, expected)
			}
		})
	}
}

func TestCompute_EmptyInput(t *testing.T) {
	input := strings.NewReader("")
	var output bytes.Buffer

	handler := ComputeHandler{input, &output}

	err := handler.Compute()
	if err == nil || err.Error() != "empty input expression" {
		t.Errorf("Compute() error: '%v', expected '%v'", err, "empty input expression")
	}
}

func TestCompute_SyntaxError(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Invalid Expression", "3 + *", "expected at least 2 tokens"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			handler := ComputeHandler{input, &output}

			err := handler.Compute()
			if err == nil || !strings.Contains(err.Error(), tt.expected) {
				t.Errorf("Compute() error: '%v', expected an error containing '%v'", err, tt.expected)
			}
		})
	}
}
