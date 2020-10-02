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
	// 	Result object.Object
	// 	Target T
	// 	Params P
	// }{
	// 	{"gender = :gender and points > :min_points", &object.Boolean{Value: false}, T{
	// 		"gender": 'F',
	// 		"points": 30,
	// 	}, P{
	// 		"gender": 'M',
	// 		"points": 10,
	// 	}},
	// 	{"gender = :gender and points > :min_points", &object.Boolean{Value: true}, T{
	// 		"gender": 'M',
	// 		"points": 11,
	// 	}, P{
	// 		"gender": 'M',
	// 		"points": 10,
	// 	}},
	// }

	// for i, s := range statements {
	// 	interpreter := makeInterpreter(s.Src, s.Target, s.Params)
	// 	r := interpreter.Exec()
	// 	if object.IsNull(r) {
	// 		t.Fatalf("Expression[%d] returns Null", i)
	// 	} else if r.(object.HashTable).HashKey() != s.Result.(object.HashTable).HashKey() {
	// 		t.Fatalf("Expression[%d] '%v' does not match the result: expected '%v', got '%v'", i, s.Src, r, s.Result)
	// 	}
	// }
}
