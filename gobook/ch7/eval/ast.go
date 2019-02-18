package eval

type Env map[Var]float64

// Expr は算術式
type Expr interface {
	// Eval は、環境 env 内でのこの Expr の値を返します。
	Eval(env Env) float64
	// Check は、この Expr 内のエラーを報告し、セットにその Var を追加します。
	Check(vars map[Var]bool) error
}

// Var は変数を特定します。例: x
type Var string

// literal は数値定数。例: 3.141
type literal float64

// unary は単項演算子を表します。例: -x
type unary struct {
	op rune // '+' か '-' のどちらか
	x  Expr
}

// binary は二項演算子を表します。例: x + y
type binary struct {
	op   rune // '+', '-', '*', '/' のどれか
	x, y Expr
}

// call は関数呼び出し式を表します。例: sin(x)
type call struct {
	fn   string // "pow", "sin", "sqrt" のどれか
	args []Expr
}
