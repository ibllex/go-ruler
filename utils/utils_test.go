package utils

import (
	"testing"

	"github.com/ibllex/go-ruler/token"
)

func TestInArray(t *testing.T) {
	arr := []interface{}{"one", "two", "three"}
	var v interface{}
	v = "one"
	if InArray(v, arr) == false {
		t.Error("error: arr contains 'one', but InArray return false")
	}

	v = "four"
	if InArray(v, arr) {
		t.Error("error: arr dose not contains 'four', but InArray return true")
	}

	ops := []interface{}{
		token.EQUAL,
		token.GT,
		token.LT,
		token.NOT_EQUAL,
		token.GT_OR_EQUAL,
		token.LT_OR_EQUAL,
	}

	if InArray(token.EQUAL, ops) == false {
		t.Error("error: ops contains 'token.EQUAL', but InArray return false")
	}
}
