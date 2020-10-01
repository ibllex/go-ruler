package lexer

import (
	"log"
	"strings"

	"github.com/ibllex/go-ruler/token"
	"github.com/ibllex/go-ruler/utils"
)

// Lexer Lexical analyzer (also know as scanner tokenizer)
type Lexer struct {
	input        string
	position     int  // current position in input (point to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char
}

func (l *Lexer) err() {
	c := string(l.ch)
	if l.ch == 0 {
		c = "EOF"
	}

	log.Fatalf("go-ruler: unexpected %s.", c)
}

// advance the 'position' pointer and set the ch variable
func (l *Lexer) advance() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// peek returns next character without increasing the value of the 'position'.
func (l *Lexer) peek() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.advance()
	}
}

func (l *Lexer) id() token.Token {
	pos := l.position
	for utils.IsAlpha(l.ch) {
		l.advance()
	}

	id := l.input[pos:l.position]
	return token.New(token.LookupIdent(id), id)
}

func (l *Lexer) number() token.Token {
	pos := l.position
	for utils.IsDigit(l.ch) || l.ch == '.' {
		l.advance()
	}

	id := l.input[pos:l.position]

	if strings.Index(id, ".") != -1 {
		return token.New(token.FLOAT_CONST, id)
	}

	return token.New(token.INTEGER_CONST, id)
}

func (l *Lexer) str() token.Token {
	lp := l.position
	// skip left quotation mark
	l.advance()
	for l.input[lp] != l.ch {
		// we can't find close quotation mark until EOF
		if l.ch == 0 {
			l.err()
		}

		l.advance()
	}
	id := l.input[lp+1 : l.position]
	// skip right quotation mark
	l.advance()
	return token.New(token.STRING_CONST, id)
}

// NextToken This method is responsible for breaking sentence
// apart into tokens. One token one time
func (l *Lexer) NextToken() token.Token {
	// if the current character is a whitespace then skip
	// consecutive whitespaces
	l.skipWhitespace()

	if tk, ok := token.LookupDoubleCharacter(l.ch, l.peek()); ok {
		l.advance()
		l.advance()
		return tk
	}

	if tk, ok := token.LookupSingleCharacter(l.ch); ok {
		l.advance()
		return tk
	}

	if utils.IsAlpha(l.ch) {
		return l.id()
	}
	if utils.IsDigit(l.ch) {
		return l.number()
	}

	if l.ch == '\'' || l.ch == '"' {
		return l.str()
	}

	if l.ch != 0 {
		l.err()
	}

	return token.New(token.EOF, "EOF")
}

// New create new lexer by given input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.advance()
	return l
}
