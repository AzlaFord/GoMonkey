package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type LetStatement struct {
	Token token.Token // token cu tip = let token.LET
	Name  *Identifier // variabila ex: x
	Value Expression  // valaorea ex : 5
}

type Statment interface {
	Node
	StatmentNode()
}
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statments []Statment
}

type Identifier struct { // indetificatorul reprezinta x
	Token token.Token // e token.IDENT
	Value string      // contine valaorea lui x
}

func (ls *LetStatement) StatmentNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (p *Program) TokenLiteral() string {
	if len(p.Statments) > 0 {
		return p.Statments[0].TokenLiteral()
	} else {
		return ""
	}
}
