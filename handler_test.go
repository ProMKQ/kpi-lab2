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
		{
			name:     "Simple Expression",
			input:    "3 4 +",
			expected: "(3 + 4)\n",
		},
		{
			name:     "Expression with multiple operators",
			input:    "3 4 + 2 *",
			expected: "((3 + 4) * 2)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			handler := ComputeHandler{
				Input:  input,
				Output: &output,
			}

			if err := handler.Compute(); err != nil {
				t.Errorf("Compute() error = %v, wantErr %v", err, nil)
			}

			got := strings.TrimSpace(output.String())
			expected := strings.TrimSpace(tt.expected)
			if got != expected {
				t.Errorf("Compute() = %v, want %v", got, expected)
			}
		})
	}
}

func TestCompute_EmptyInput(t *testing.T) {
	input := strings.NewReader("")
	var output bytes.Buffer

	handler := ComputeHandler{
		Input:  input,
		Output: &output,
	}

	err := handler.Compute()
	if err == nil || err.Error() != "empty input expression" {
		t.Errorf("Compute() error = %v, wantErr %v", err, "empty input expression")
	}
}

func TestCompute_SyntaxError(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Invalid Expression",
			input:    "3 + *",
			expected: "expected at least 2 tokens",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			handler := ComputeHandler{
				Input:  input,
				Output: &output,
			}

			err := handler.Compute()
			if err == nil || !strings.Contains(err.Error(), tt.expected) {
				t.Errorf("Compute() error = %v, wantErr containing %v", err, tt.expected)
			}
		})
	}
}
