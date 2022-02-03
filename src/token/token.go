package token

func Hello() string {
	return "Hello! This is token.go! in token package"
}

type TokenType string

type Token struct {
	Token   TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":   SFUNCTION,
	"func": BFUNCTION,
	"Main": MAINF,
}

func ThatisIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

const (
	SFUNCTION = "SFUNCTION"
	BFUNCTION = "BFUNCTION"

	IDENT = "IDENT"
	MAINF = "MAINF"

	LBRACE = "{"
	RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"
	EOF    = ""
)
