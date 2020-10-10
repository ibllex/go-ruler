package parser

import (
	"testing"

	"github.com/ibllex/go-ruler/lexer"
)

func makeParser(src string) Parser {
	l := lexer.New(src)
	return New(l)
}

func TestParser(t *testing.T) {
	statements := []struct {
		Src string
		Err bool
	}{
		{"gender = :gender and points > :min_points", false},
		{"group.name = \"Players\"", false},
		{"gender = \"F\"", false},
		{"gender != \"F\"", false},
		{"not(gender = \"F\")", false},
		{"male or female", false},
		{"target = find(:data, 'data')", false},
		{"target <= find(:data, 'data')", false},
		{"group = ? or group = ?", false},

		{"gender and", true},
		{"and gender", true},
		{"gender or", true},
		{"or gender", true},
		{"xor gender", true},
		{"gender xor", true},
	}

	for _, s := range statements {
		p := makeParser(s.Src)
		_, err := p.Parse()

		if s.Err && err == nil {
			t.Errorf("'%v' error: expected error but nil got.", s.Src)
		}

		if s.Err == false && err != nil {
			t.Errorf("'%v' error: expected no error but got error: %v.", s.Src, err)
		}
	}
}
