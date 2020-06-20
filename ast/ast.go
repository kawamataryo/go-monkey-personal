package ast

import "github.com/kawamataryo/go-monkey/token"

// 全てのノードが実装する
type Node interface {
	TokenLiteral() string // デバックとテストのために用いる
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
	// 軸解析で得たトークン （typeとliteralを持つ）
	Token token.Token // token.LET
	// 束縛する変数名
	Name  *Identifier
	// 束縛される式
	value Expression
}

// letステートメント構造体が持つダミーメソッド
func (ls *LetStatement) statementNode() {}
// letステートメント構造体が持つトークンリテラルの出力メソッド
// デバックテストに使われる
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// 束縛の識別子を保持するためのもの。
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
