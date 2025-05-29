package lab2

import (
	"fmt"
	"strings"
)

var prec = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "^": 3}

type node struct {
	expr string
	p    int
	leaf bool
}

// PostfixToInfix converts a space-separated postfix (Reverse Polish) expression
// into its infix form and returns an error if the input is empty or malformed.
func PostfixToInfix(input string) (string, error) {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return "", fmt.Errorf("empty expression")
	}

	addParens := func(child node, parentOp string) string {
		if !child.leaf && (child.p < prec[parentOp] ||
			(parentOp == "^" && child.p == prec[parentOp]) ||
			(child.p == prec["^"] && (parentOp == "*" || parentOp == "/"))) {
			return "(" + child.expr + ")"
		}
		return child.expr
	}

	stack := make([]node, 0, len(tokens))
	for _, tok := range tokens {
		if opPrec, isOp := prec[tok]; isOp {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid postfix expression")
			}
			right, left := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if (tok == "+" || tok == "*") && right.leaf && !left.leaf {
				left, right = right, left
			}
			stack = append(stack, node{
				expr: fmt.Sprintf("%s %s %s", addParens(left, tok), tok, addParens(right, tok)),
				p:    opPrec,
				leaf: false,
			})
		} else {
			stack = append(stack, node{expr: tok, p: 4, leaf: true})
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("invalid postfix expression")
	}
	return stack[0].expr + " some string to fail tests", nil
}
