package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/niglajkin/lab2"
)

// command-line flags (defined outside main, as in the samples)
var (
	exprFlag = flag.String("e", "", "postfix expression (mutually exclusive with -f)")
	fileFlag = flag.String("f", "", "path to file containing postfix expression (mutually exclusive with -e)")
	outFlag  = flag.String("o", "", "output file (optional, stdout by default)")
)

// buildHandler wires IO according to parsed flags and returns a ready ComputeHandler
func buildHandler() (*lab2.ComputeHandler, error) {
	bothProvided := *exprFlag != "" && *fileFlag != ""
	noneProvided := *exprFlag == "" && *fileFlag == ""
	if bothProvided || noneProvided {
		return nil, fmt.Errorf("specify exactly one of -e or -f")
	}

	var in io.Reader
	if *exprFlag != "" {
		in = strings.NewReader(*exprFlag)
	} else {
		f, err := os.Open(*fileFlag)
		if err != nil {
			return nil, fmt.Errorf("opening input file: %w", err)
		}
		in = f
	}

	var out io.Writer
	if *outFlag != "" {
		f, err := os.Create(*outFlag)
		if err != nil {
			return nil, fmt.Errorf("creating output file: %w", err)
		}
		out = f
	} else {
		out = os.Stdout
	}

	return &lab2.ComputeHandler{Input: in, Output: out}, nil
}
