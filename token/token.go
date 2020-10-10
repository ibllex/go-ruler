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
	FALSE Type = "FALSE"
	TRUE  Type = "TRUE"
	NULL  Type = "NULL"

	// logical operators
	AND Type = "AND"
	OR  Type = "OR"
	XOR Type = "XOR"

	// single-character token types
	LPAREN Type = "("
	RPAREN Type = ")"
	COMMA  Type = ","
	DOT    Type = "."
	COLON  Type = ":"
	EQUAL  Type = "="
	GT     Type = ">"
	LT     Type = "<"
	QM     Type = "?"

	// double-character token types
	NOT_EQUAL   Type = "!="
	GT_OR_EQUAL Type = ">="
	LT_OR_EQUAL Type = "<="

	// misc
	IDENT         Type = "IDENT"
	EOF           Type = "EOF"
	INTEGER_CONST Type = "INTEGER_CONST"
	FLOAT_CONST   Type = "FLOAT_CONST"
	STRING_CONST  Type = "STRING_CONST"
)

var keywords = map[string]Type{
	"false": FALSE,
	"true":  TRUE,
	"null":  NULL,
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
	'?': QM,
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
