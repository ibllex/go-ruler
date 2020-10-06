package ruler

import (
	"testing"

	"github.com/ibllex/go-ruler/interpreter"
	"github.com/ibllex/go-ruler/operator"
)

func TestSatisfies(t *testing.T) {

	operators := interpreter.O{
		"in_array": operator.InArray,
	}

	ruler := New(operators)
	rule := "(gender = :gender and points > :min_points) or in_array(:users, pseudo)"
	params := interpreter.P{
		"min_points": 30,
		"gender":     "M",
		"users": []interface{}{
			"Birdie", "Diana",
		},
	}

	players := []struct {
		Target interpreter.T
		Result bool
	}{
		{interpreter.T{"pseudo": "Joe", "gender": "M", "points": 40}, true},
		{interpreter.T{"pseudo": "Moe", "gender": "M", "points": 20}, false},
		{interpreter.T{"pseudo": "Alice", "gender": "F", "points": 60}, false},
		{interpreter.T{"pseudo": "Birdie", "gender": "F", "points": 60}, true},
	}

	for i, p := range players {
		ret, err := ruler.Satisfies(p.Target, rule, params)
		if err != nil {
			t.Errorf("players[%d] error: %v", i, err)
		}

		if ret != p.Result {
			t.Errorf("players[%d] satisfies result mismatch expected", i)
		}
	}
}

func TestFilter(t *testing.T) {

	operators := interpreter.O{
		"in_array": operator.InArray,
	}

	ruler := New(operators)
	rule := "(gender = :gender and points > :min_points) or in_array(:users, pseudo)"

	params := interpreter.P{
		"min_points": 30,
		"gender":     "M",
		"users": []interface{}{
			"Birdie", "Diana",
		},
	}

	users := []interpreter.T{
		{"pseudo": "Joe", "gender": "M", "points": 40},
		{"pseudo": "Moe", "gender": "M", "points": 20},
		{"pseudo": "Alice", "gender": "F", "points": 60},
		{"pseudo": "Birdie", "gender": "F", "points": 60},
	}

	remainder, err := ruler.Filter(users, rule, params)
	if err != nil {
		t.Errorf("filter error: %v", err)
	}

	if len(remainder) != 2 {
		t.Errorf("remainder's count greater than one")
	}
}
