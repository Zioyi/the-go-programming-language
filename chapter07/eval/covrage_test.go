package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestCoverage(t *testing.T) {
	var tests = []struct {
		input string
		env   Env
		want  string
	}{
		{"5 + 10", nil, "15"},
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
	}

	for _, test := range tests {
		expr, err := Parse(test.input)
		if err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Printf("%+v\n", expr)
		// got := fmt.Sprintf("%.6g", expr.Eval(test.env))

	}
}
