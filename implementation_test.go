package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	valid := []struct {
		name  string
		input string
		want  string
	}{
		{"add", "2 3 +", "2 + 3"},
		{"sub", "4 2 -", "4 - 2"},
		{"mul", "6 7 *", "6 * 7"},
		{"div", "8 2 /", "8 / 2"},
		{"pow", "3 2 ^", "3 ^ 2"},

		{"reorder", "4 2 - 3 * 5 +", "5 + 3 * (4 - 2)"},

		{"complex7", "1 2 + 3 * 4 5 - / 6 7 ^ +", "3 * (1 + 2) / (4 - 5) + 6 ^ 7"},
		{"complex10", "1 2 + 3 * 4 5 - 6 / ^ 7 8 + * 9 / 10 -", "((3 * (1 + 2)) ^ ((4 - 5) / 6)) * (7 + 8) / 9 - 10"},
	}

	for _, tc := range valid {
		t.Run(tc.name, func(t *testing.T) {
			got, err := PostfixToInfix(tc.input)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPostfixToInfix_Errors(t *testing.T) {
	bad := []string{
		"",
		"5 +",
		"2 3 4",
		"2 3 + +",
		"2 3 &",
	}

	for _, input := range bad {
		t.Run(fmt.Sprintf("error_%q", input), func(t *testing.T) {
			_, err := PostfixToInfix(input)
			assert.Error(t, err)
		})
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("1 2 + 3 * 4 5 - / 6 7 ^ +")
	fmt.Println(res)
	// Output:
	// 3 * (1 + 2) / (4 - 5) + 6 ^ 7
}
