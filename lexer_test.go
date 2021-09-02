package schemalex

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLexToken(t *testing.T) {
	type Spec struct {
		input string
		token Token
	}

	specs := []Spec{
		// number
		{
			input: "123",
			token: Token{Value: "123", Type: NUMBER},
		},
		{
			input: ".2",
			token: Token{Value: ".2", Type: NUMBER},
		},
		{
			input: "3.4",
			token: Token{Value: "3.4", Type: NUMBER},
		},
		{
			input: "-5",
			token: Token{Value: "-5", Type: NUMBER},
		},
		{
			input: "-6.78",
			token: Token{Value: "-6.78", Type: NUMBER},
		},
		{
			input: "+9.10",
			token: Token{Value: "+9.10", Type: NUMBER},
		},
		{
			input: "1.2E3",
			token: Token{Value: "1.2E3", Type: NUMBER},
		},
		{
			input: "1.2E-3",
			token: Token{Value: "1.2E-3", Type: NUMBER},
		},
		{
			input: "-1.2E3",
			token: Token{Value: "-1.2E3", Type: NUMBER},
		},
		{
			input: "-1.2E-3",
			token: Token{Value: "-1.2E-3", Type: NUMBER},
		},
		// SINGLE_QUOTE_IDENT
		{
			input: `'hoge'`,
			token: Token{Value: `hoge`, Type: SINGLE_QUOTE_IDENT},
		},
		{
			input: `'ho''ge'`,
			token: Token{Value: `ho'ge`, Type: SINGLE_QUOTE_IDENT},
		},
		// DOUBLE_QUOTE_IDENT
		{
			input: `"hoge"`,
			token: Token{Value: `hoge`, Type: DOUBLE_QUOTE_IDENT},
		},
		{
			input: `"ho""ge"`,
			token: Token{Value: `ho"ge`, Type: DOUBLE_QUOTE_IDENT},
		},
		// BACKTICK_IDENT
		{
			input: "`hoge`",
			token: Token{Value: "hoge", Type: BACKTICK_IDENT},
		},
		{
			input: "`ho``ge`",
			token: Token{Value: "ho`ge", Type: BACKTICK_IDENT},
		},
		// ESCAPED STRING BY BACKSLASH
		{
			input: `'ho\'ge'`,
			token: Token{Value: `ho'ge`, Type: SINGLE_QUOTE_IDENT},
		},
	}

	for _, spec := range specs {
		t.Run(spec.input, func(t *testing.T) {
			tok := lex([]byte(spec.input))
			spec.token.Line = 1
			spec.token.Col = 1
			if diff := cmp.Diff(&spec.token, tok[0]); diff != "" {
				t.Errorf("tok mismatch: (-want/+got):\n%s", diff)
			}
		})
	}
}
