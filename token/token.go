package token

// Type token's type name
type Type string

// Token token type
type Token struct {
	Type    Type
	Literal string
}

const (
	// scalars
	FALSE = "FALSE"
	TRUE  = "TRUE"
	NULL  = "NULL"

	// logical operators
	NOT = "NOT"
	AND = "AND"
	OR  = "OR"
	XOR = "XOR"

	// single-character token types
	LPAREN = "("
	RPAREN = ")"
	COMMA  = ","
	DOT    = "."
	COLON  = ":"
	EQUAL  = "="
	GT     = ">"
	LT     = "<"

	// double-character token types
	NOT_EQUAL   = "!="
	GT_OR_EQUAL = ">="
	LT_OR_EQUAL = "<="

	// misc
	IDENT         = "IDENT"
	EOF           = "EOF"
	INTEGER_CONST = "INTEGER_CONST"
	FLOAT_CONST   = "FLOAT_CONST"
	STRING_CONST  = "STRING_CONST"
)

var keywords = map[string]Type{
	"false": FALSE,
	"true":  TRUE,
	"null":  NULL,
	"not":   NOT,
	"and":   AND,
	"or":    OR,
	"xor":   XOR,
}

var singleCharacterTypes = map[byte]Type{
	'(': LPAREN,
	')': RPAREN,
	',': COMMA,
	'.': DOT,
	':': COLON,
	'=': EQUAL,
	'>': GT,
	'<': LT,
}

var doubleCharacterTypes = map[string]Type{
	"!=": NOT_EQUAL,
	">=": GT_OR_EQUAL,
	"<=": LT_OR_EQUAL,
}

// LookupDoubleCharacter if the current character and the next character
// are combined into a valid token.
func LookupDoubleCharacter(first byte, second byte) (Token, bool) {
	sign := string([]byte{first, second})
	if tk, ok := doubleCharacterTypes[sign]; ok {
		return New(tk, sign), true
	}

	return Token{}, false
}

// LookupSingleCharacter if a character is a token type then return
func LookupSingleCharacter(ch byte) (Token, bool) {
	if tk, ok := singleCharacterTypes[ch]; ok {
		return New(tk, string(ch)), true
	}

	return Token{}, false
}

// LookupIdent if an ident is a keyword then return keyword token type
// otherwise return IDENT token type
func LookupIdent(ident string) Type {
	if tk, ok := keywords[ident]; ok {
		return tk
	}

	return IDENT
}

// New create Token by given token type and literal
func New(tokenType Type, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}
