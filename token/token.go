package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

const (
	ILLEGAL = "ILLEGAL" // e un token/character care nu stim despre el
	EOF     = "EOF"     // EOF e END OF FILE care explica la parser ca se poate opri

	// indetifiers + literal care sunt valorile adica
	IDENT = "IDENT" //add , foobar, x , y ...
	INT   = "INT"   //123456789

	// Operatori
	ASSIGN = "="
	PLUS   = "+"

	// Delimitatori
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Cuvinte cheie
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

func LookupIdent(ident string) TokenType {
	// verificam daca tok are valoare si daca are returnam tok adica tipul de token
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
