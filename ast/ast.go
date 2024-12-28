package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

// there are two different types of nodes - statement and expression
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// first implementation of the Node
type Program struct {
	Statements []Statement
}

// Program is implementing the node interface
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let x = 5 + 5;
// <token> <identifier> = <expression>
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// IDENT token and literal value
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
