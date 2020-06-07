package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

// トークンと文字列の対応表
const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // 識別子 + リテラル
    IDENT = "IDENT" // add, foobar, x, y, ...
    INT   = "INT"   // 1343456

    // 演算子
    ASSIGN   = "="
    PLUS     = "+"

    // デリミタ
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // キーワード
    FUNCTION = "FUNCTION"
    LET      = "LET"
)

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}