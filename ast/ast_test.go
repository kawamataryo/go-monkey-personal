package ast

import (
	"github.com/kawamataryo/go-monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	// LetStatement
	programLet := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if programLet.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", programLet.String())
	}

	// ReturnStatement
	programReturn := &Program{
		Statements: []Statement{
			&ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "returnValue"},
					Value: "returnValue",
				},
			},
		},
	}

	if programReturn.String() != "return returnValue;" {
		t.Errorf("program.String() wrong. got=%q", programReturn.String())
	}

}
