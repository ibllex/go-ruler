package interpreter

import (
	"testing"

	"github.com/ibllex/go-ruler/lexer"
	"github.com/ibllex/go-ruler/parser"
)

func makeInterpreter(src string, target T, params P) Interpreter {
	p := parser.New(lexer.New(src))
	tree, _ := p.Parse()

	return New(tree, target, params)
}

func TestInterpreter(t *testing.T) {
	// statements := []struct {
	// 	Src    string
	// 	Result bool
	// 	Target T
	// 	Params P
	// }{
	// 	{"gender = :gender and points > :min_points", false, T{
	// 		"gender": 'F',
	// 		"points": 30,
	// 	}, P{
	// 		"gender": 'M',
	// 		"points": 10,
	// 	}},
	// }

	// for s := range statements {
	// 	//
	// }
}
