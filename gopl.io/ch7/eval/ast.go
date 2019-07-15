package eval

type Expr interface {
	// Eval 返回表达式在 env 上下文中的值
	Eval(env Env) float64
	// Check 方法报告表达式中的错误，并把表达式中的值加入Vars中
	Check(vars map[Var]bool) error
}

// Var 表示一个变量，比如 x
type Var string

// literal 是一个数字常量，比如 3.141
type literal float64

// unary 表示一元操作表达式，比如 -x
type unary struct {
	op rune // '+', '-'中的一个
	x  Expr
}

// binary 表示二元操作表达式， 比如 x + y
type binary struct {
	op   rune // '+', '-', '*', '/'中的一个
	x, y Expr
}

// call 表示函数调用表达式，比如 sin(x)
type call struct {
	fn   string // "pow", "sin", "sqrt"中的一个
	args []Expr
}
