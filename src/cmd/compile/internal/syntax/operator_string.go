// Code generated by "stringer -type Operator -linecomment tokens.go"; DO NOT EDIT.

package syntax

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Def-1]
	_ = x[Not-2]
	_ = x[Recv-3]
	_ = x[Tilde-4]
	_ = x[OrOr-5]
	_ = x[AndAnd-6]
	_ = x[Eql-7]
	_ = x[Neq-8]
	_ = x[Neql-9]
	_ = x[Lss-10]
	_ = x[Leq-11]
	_ = x[Gtr-12]
	_ = x[Geq-13]
	_ = x[Add-14]
	_ = x[Sub-15]
	_ = x[Or-16]
	_ = x[Nor-17]
	_ = x[Xor-18]
	_ = x[Mul-19]
	_ = x[Div-20]
	_ = x[Rem-21]
	_ = x[And-22]
	_ = x[Nand-23]
	_ = x[AndNot-24]
	_ = x[Shl-25]
	_ = x[Shr-26]
	_ = x[Pow-27]
}

const _Operator_name = ":!<-~||&&==!=<><<=>>=+-|!|^*/%&!&&^<<>>^^"

var _Operator_index = [...]uint8{0, 1, 2, 4, 5, 7, 9, 11, 13, 15, 16, 18, 19, 21, 22, 23, 24, 26, 27, 28, 29, 30, 31, 33, 35, 37, 39, 41}

func (i Operator) String() string {
	i -= 1
	if i >= Operator(len(_Operator_index)-1) {
		return "Operator(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Operator_name[_Operator_index[i]:_Operator_index[i+1]]
}
