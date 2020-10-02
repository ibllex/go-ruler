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

func TestQueryInMapInter(t *testing.T) {
	mapData := map[string]interface{}{
		"country": "jp",
		"age":     18,
		"friends": map[string]interface{}{
			"mia": map[string]interface{}{
				"country": "us",
				"age":     17,
			},
			"ashley": map[string]interface{}{
				"country": "uk",
				"age":     20,
			},
		},
	}

	paths := []struct {
		Path   []string
		Result interface{}
	}{
		{[]string{"country"}, "jp"},
		{[]string{"age"}, 18},
		{[]string{"friends", "mia", "country"}, "us"},
		{[]string{"friends", "mia", "age"}, 17},
		{[]string{"friends", "ashley", "country"}, "uk"},
		{[]string{"friends", "ashley", "age"}, 20},
		{[]string{"google"}, nil},
		{[]string{"friends.amir"}, nil},
		{[]string{"friends.amir.country"}, nil},
	}

	for i, p := range paths {
		r := QueryInMapInter(mapData, p.Path)
		if r != p.Result {
			t.Fatalf("paths[%d] query not match, expected %v but got %v.", i, p.Result, r)
		}
	}
}
