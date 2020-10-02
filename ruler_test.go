package ruler

import (
	"testing"

	"github.com/ibllex/go-ruler/interpreter"
)

func TestSatisfies(t *testing.T) {
	ruler := New()
	rule := "gender = :gender and points > :min_points"
	params := interpreter.P{
		"min_points": 30,
		"gender":     "M",
	}

	players := []struct {
		Target interpreter.T
		Result bool
	}{
		{interpreter.T{"pseudo": "Joe", "gender": "M", "points": 40}, true},
		{interpreter.T{"pseudo": "Moe", "gender": "M", "points": 20}, false},
		{interpreter.T{"pseudo": "Alice", "gender": "F", "points": 60}, false},
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
