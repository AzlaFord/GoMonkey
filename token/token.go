package token
type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // e un token/character care nu stim despre el 
	EOF = "EOF" // EOF e END OF FILE care explica la parser ca se poate opri
	
	// indetifiers + literal care sunt valorile adica 
	IDENT = "IDENT" //add , foobar, x , y ...
	INT = "INT" //123456789

	// Operatori
	ASSIGN = "="
	PLUS = "+"

	// Delimitatori
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Cuvinte cheie
	FUNCTION = "FUNCTION"
	LET = "LET"
)