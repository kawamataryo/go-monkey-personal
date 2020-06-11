package lexer

import "github.com/kawamataryo/go-monkey/token"

// 字句解析器の構造体
type Lexer struct {
	input        string // 初期化時に受け取る文字列
	position     int    // 現在の文字の位置
	readPosition int    // 次に読む文字の位置
	ch           byte   // 現在検査中の文字（バイト）
}

// lexerのコンストラクタ
func New(input string) *Lexer {
	// lにinputをセットしたLexerのポインタを格納
	l := &Lexer{input: input}
	// 一文字の読み込みを実行
	l.readChar()
	// ポインタを返す
	return l
}

// 次の1文字を読んでchに格納して、inputの文字列の現在位置を1文字進める関数
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 入力値の最終文字まで来たらchに0を設定する
		l.ch = 0
	} else {
		// それ以外だったら次に読む文字をchに設定する
		l.ch = l.input[l.readPosition]
	}
	// 次に読む文字位置で現在の文字位置を更新
	l.position = l.readPosition
	// 次に読む文字をインクリメントする
	l.readPosition++
}

// chを読んで対応するtokを返す
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 認識された識別子以外の文字列リテラルをまとめて読み取る
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
