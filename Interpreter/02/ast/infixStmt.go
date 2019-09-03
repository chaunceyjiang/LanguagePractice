package ast

import (
	"Interpreter/token"
	"bytes"
)

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}
func (oe *InfixExpression) expressionNode() {

}
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" "+oe.Operator+" ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()

}
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}
