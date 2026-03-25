package lexer

import (
	"monkey/token"
	"testing"
)

type Lexer struct {
	input        string
	position     int  // pozitia currenta in input  duce la caracterul curet
	readPosition int  // pozitia currenta de citire din input dupa caracterul current
	ch           byte // caracterul curent in evaluare care pozitia la varaibila position
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // verifica daca a ajuns la sfarsitul liniei
		l.ch = 0 // e ASCII cod pentru null valaore 0 si semnifica ori nu am citit nimic inca ori e sfarsit de file EOF
	} else {
		l.ch = l.input[l.readPosition] // daca nu am ajuns la farsitul input-ului seteaza l.ch ca urmatorul caracter l.ch =l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func TestNextToken(t *testing.T) {
	input := "=+(){},;"
	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.COMMA, ","},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range test {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
