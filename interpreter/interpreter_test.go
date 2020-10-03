package interpreter

import (
	"testing"

	"github.com/ibllex/go-ruler/lexer"
	"github.com/ibllex/go-ruler/object"
	"github.com/ibllex/go-ruler/parser"
)

type filterItem struct {
	Result object.Object
	Target T
	Params P
}

func makeInterpreter(src string, ops O) Interpreter {
	p := parser.New(lexer.New(src))
	tree, _ := p.Parse()

	return New(tree, ops)
}

func TestInterpreter(t *testing.T) {

	operators := O{
		"empty": func(args ...object.Object) object.Object {
			return &object.Boolean{Value: true}
		},
		"in_array": func(args ...object.Object) object.Object {
			return &object.Boolean{Value: true}
		},
	}

	statements := make(map[string][]filterItem)
	statements["gender = :gender and points > :min_points"] = []filterItem{
		{
			Result: &object.Boolean{Value: false},
			Target: T{"gender": "F", "points": 30},
			Params: P{"gender": "M", "min_points": 10},
		},
		{
			Result: &object.Boolean{Value: false},
			Target: T{"gender": "M", "points": 9},
			Params: P{"gender": "M", "min_points": 10},
		},
		{
			Result: &object.Boolean{Value: true},
			Target: T{"gender": "M", "points": 11},
			Params: P{"gender": "M", "min_points": 10},
		},
	}

	// function call
	statements["empty(name) or in_array(:users, name)"] = []filterItem{
		{
			Result: &object.Boolean{Value: true},
			Target: T{"name": "Joe"},
			Params: P{"users": []string{"Joe", "Moe"}},
		},
		{
			Result: &object.Boolean{Value: true},
			Target: T{"name": "Moe"},
			Params: P{"users": []string{"Joe", "Moe"}},
		},
		{
			Result: &object.Boolean{Value: true},
			Target: T{"name": ""},
			Params: P{"users": []string{"Joe", "Moe"}},
		},
		{
			Result: &object.Boolean{Value: false},
			Target: T{"name": "Alice"},
			Params: P{"users": []string{"Joe", "Moe"}},
		},
	}

	for rule, items := range statements {
		interpreter := makeInterpreter(rule, operators)
		for i, s := range items {
			r := interpreter.Exec(s.Target, s.Params)
			if object.IsNull(r) {
				t.Fatalf("Rule: %s [%d] returns Null", rule, i)
			} else if r.(object.HashTable).HashKey() != s.Result.(object.HashTable).HashKey() {
				t.Fatalf("Rule: %s [%d] does not match the result: expected '%v', got '%v'", rule, i, r, s.Result)
			}
		}
	}
}
