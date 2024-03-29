package ast

import (
	"bytes"
	"github.com/kawamataryo/go-monkey/token"
)

// 全てのノードが実装する
type Node interface {
	TokenLiteral() string // デバックとテストのために用いる
	String() string
}

// 一部のノードが実装する
// Statement = 文。
type Statement interface {
	Node
	statementNode() // ダミーメソッド？
}

// 一部のノードが実装する
// Expression = 式。値を生成するもの
type Expression interface {
	Node
	expressionNode() // ダミーメソッド？
}

// 最初のノード
// Statementのスライスを持つ
type Program struct {
	Statements []Statement
}

// トークンリテラルを出力するための関数
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Letを表すASTのノードの実装
type LetStatement struct {
	// 字句解析で得たトークン （typeとliteralを持つ）
	Token token.Token // token.LET
	// 束縛する変数名
	Name  *Identifier
	// 束縛される式
	Value Expression
}

// letステートメント構造体が持つダミーメソッド
func (ls *LetStatement) statementNode() {}
// letステートメント構造体が持つトークンリテラルの出力メソッド
// デバックテストに使われる
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// 束縛の識別子を保持するためのもの
// Expressionインターフェイスを実装する。
// なぜなら、let以外の場所では、識別子は値を生成するから。（変数からの値の取り出しのように）
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

// Returnを表すASTのノードの実装
type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

type ExpressionStatement struct {
	Token token.Token // 式の最初のトークン
	Expression Expression // 式を保持する
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string  { return es.Token.Literal }


// 再起的にASTを文字列に変換している
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (ls *LetStatement)  String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
