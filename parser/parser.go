package parser

import (
	"github.com/kawamataryo/go-monkey/ast"
	"github.com/kawamataryo/go-monkey/lexer"
	"github.com/kawamataryo/go-monkey/token"
)

type Parser struct {
	l *lexer.Lexer // 字句解析器インスタンスへのポインタ
	curToken token.Token // 現在のトークン
	peekToken token.Token // 次のトークン
}

func New(l *lexer.Lexer) *Parser  {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

// トークンを読み進める
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		// パースの実行
		stmt := p.parseStatement()

		if stmt != nil {
			// パース結果がnilでなければ、パース結果をProgramのstatementsに追加している
			program.Statements = append(program.Statements, stmt)
		}
		// 次のトークンへ
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	// 今のトークンのタイプをみてパーサーを実行する
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// letのASTの雛形を作る
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// アサーション関数
// peekTokenの型をチェックして、その型が正しい場合に限ってnextTokenを呼んでトークンを進める
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
