package lab2

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (c *ComputeHandler) Compute() error {
	data, err := io.ReadAll(c.Input)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	expr := strings.TrimSpace(string(data))
	if expr == "" {
		return errors.New("empty input expression")
	}

	infix, err := PostfixToInfix(expr)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(c.Output, infix)
	if err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}
