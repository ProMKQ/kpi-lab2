package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/ProMKQ/kpi-lab2"
)

func main() {
	exprFlag := flag.String("e", "", "Expression in postfix notation")
	fileFlag := flag.String("t", "", "Input file containing the expression")
	outFlag := flag.String("o", "", "Output file for the result")
	flag.Parse()

	if *exprFlag != "" && *fileFlag != "" {
		fmt.Fprintln(os.Stderr, "Error: specify either -e or -t, not both")
		os.Exit(1)
	}

	var input io.Reader
	if *exprFlag != "" {
		input = strings.NewReader(*exprFlag)
	} else if *fileFlag != "" {
		f, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to open input file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		input = f
	} else {
		fmt.Fprintln(os.Stderr, "Error: no input provided. Use -e or -t")
		os.Exit(1)
	}

	var output io.Writer = os.Stdout
	if *outFlag != "" {
		f, err := os.Create(*outFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to create output file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		output = f
	}

	handler := lab2.ComputeHandler{Input: input, Output: output}
	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
