package token

func Hello() string {
	return "Hello! This is token.go! in token package"
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     SFUNCTION,
	"func":   BFUNCTION,
	"Main":   MAINF,
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"var":    VAR,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"const":  CONST,
}

func IsThatIdent(ident string) TokenType {
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

	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	FOR    = "FOR"
	VAR    = "VAR"
	CONST  = "CONST"

	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	ASSIGN   = "="
	GOASSIGN = ":="

	EQ     = "=="
	NOT_EQ = "!="

	BANG     = "!"
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	DOT      = "."

	LBRACKET = "["
	RBRACKET = "]"
	LBRACE   = "{"
	RBRACE   = "}"
	LPAREN   = "("
	RPAREN   = ")"

	COLON     = ":"
	SEMICOLON = ";"
	COMMA     = ","

	EOF     = ""
	ILLEGAL = "ILLEGAL"
)
