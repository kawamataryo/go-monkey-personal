package parser

import (
	"fmt"
	"github.com/kawamataryo/go-monkey/ast"
	"github.com/kawamataryo/go-monkey/lexer"
	"github.com/kawamataryo/go-monkey/token"
)

// 式の優先順位の定数。LOWESTから1, 2, 3と順次割り当てられている
const (
	_ int = iota
	LOWEST
	EQUALS // ==
	LESSGREATER // > または <
	SUM // +
	PRODUCT // *
	PREFIX // -X または !X
	CALL // myFunction(X)
)

// 前置構文解析関数
// 中置構文解析関数
type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression)  ast.Expression
)

type Parser struct {
	l *lexer.Lexer // 字句解析器インスタンスへのポインタ
	curToken token.Token // 現在のトークン
	peekToken token.Token // 次のトークン
	errors []string // Error

	prefixParseFns map[token.TokenType]prefixParseFn // 前置構文解析関数のmap
	infixParseFns map[token.TokenType]infixParseFn // 中置構文解析関数のmap
}

func New(l *lexer.Lexer) *Parser  {
	p := &Parser{
		l: l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) Errors() []string {
	return p.errors
}

// errorメッセージを代入する
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s insted", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// トークンを読み進める
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
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
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionsStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// letのASTの雛形を作る
	stmt := &ast.LetStatement{Token: p.curToken}

	// 識別子が次にくることを期待する
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// 識別子のASTを作り、LetStatementのNameに設定する（これが変数名）
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// 次に符号（=）がくること期待する
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// 今のトークンが何なのか判定
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// 次のトークンが何なのか判定
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
		p.peekError(t)
		return false
	}
}

// returnのパーサー
func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{
		Token: p.curToken,
	}

	p.nextToken()

	// TODO: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// 構文解析関数の記録
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn)  {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn)  {
	p.infixParseFns[tokenType] = fn
}

// 式文の構文解析

func (p *Parser) parseExpressionsStatement() *ast.ExpressionStatement  {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}
