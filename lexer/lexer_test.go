package lexer

import (
	"os"
	"os/exec"
	"testing"

	"github.com/ibllex/go-ruler/token"
)

func TestNextToken(t *testing.T) {

	input := "() , . : = > < != >= <= false true null not and or xor 10 20.0 'username' \"hello\" :gender female"

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COMMA, ","},
		{token.DOT, "."},
		{token.COLON, ":"},
		{token.EQUAL, "="},
		{token.GT, ">"},
		{token.LT, "<"},
		{token.NOT_EQUAL, "!="},
		{token.GT_OR_EQUAL, ">="},
		{token.LT_OR_EQUAL, "<="},
		{token.FALSE, "false"},
		{token.TRUE, "true"},
		{token.NULL, "null"},
		{token.NOT, "not"},
		{token.AND, "and"},
		{token.OR, "or"},
		{token.XOR, "xor"},
		{token.INTEGER_CONST, "10"},
		{token.FLOAT_CONST, "20.0"},
		{token.STRING_CONST, "username"},
		{token.STRING_CONST, "hello"},
		{token.COLON, ":"},
		{token.IDENT, "gender"},
		{token.IDENT, "female"},
	}

	l := New(input)

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}

func TestNextTokenErr(t *testing.T) {

	t.Run("string type without closing single quotes", func(t *testing.T) {
		assertExit(t, "TestNextTokenErr", func() {
			l := New("'username")
			l.NextToken()
		})
	})

	t.Run("string type without closing double quotes", func(t *testing.T) {
		assertExit(t, "TestNextTokenErr", func() {
			l := New("\"username")
			l.NextToken()
		})
	})

	t.Run("unexpected token", func(t *testing.T) {
		assertExit(t, "TestNextTokenErr", func() {
			l := New("@")
			l.NextToken()
		})
	})

}

func assertExit(t *testing.T, testCaseName string, fn func()) {
	t.Helper()

	if os.Getenv("BE_CRASHER") == "1" {
		fn()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run="+testCaseName)
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("process ran with err %v, want exit status 1", err)
}
