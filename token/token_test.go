package token

import "testing"

func TestLookupIdent(t *testing.T) {
	type args struct {
		ident string
	}
	tests := []struct {
		name string
		args args
		want TokenType
	}{
		{ "let", args{ "let" }, LET},
		{ "fun", args{ "fn" }, FUNCTION},
		{ "other", args{ "funnn " }, IDENT},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LookupIdent(tt.args.ident); got != tt.want {
				t.Errorf("LookupIdent() = %v, want %v", got, tt.want)
			}
		})
	}
}
