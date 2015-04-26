// Package token defines constants representing the lexical tokens of SQL
// and basic operations on tokens.
//
package token

// Token is the set of lexical tokens of SQL.
type Token int

// The list of tokens.
const (
	ILLEGAL Token = iota
	EOF
	COMMENT

	operator_beg
	// Operators and delimiters
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND // &
	OR  // |

	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ    // !=
	LEQ    // <=
	GEQ    // >=
	DEFINE // :=

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	operator_end
)
