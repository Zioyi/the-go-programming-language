package calc

type Expr interface {
	Eval(env Env) float64
}

// Var identifies a variable, e.g., x.
type Var string

// literal is numeric constant, e.g., 3.141
type literal float64

// unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

type binary struct {
	op   rune // one of '+', '-', '*', '-'
	x, y Expr
}

type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}
