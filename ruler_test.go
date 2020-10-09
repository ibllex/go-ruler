package ruler

import (
	"testing"

	"github.com/ibllex/go-ruler/interpreter"
	"github.com/ibllex/go-ruler/operator"
	"github.com/ibllex/go-ruler/spec"
)

// IsFemale Mock specification
type IsFemale struct {
	//
}

func (s *IsFemale) Rule() string {
	return "gender = 'F'"
}

func (s *IsFemale) Params() interpreter.P {
	return interpreter.P{}
}

// PlayerMinScore Mock specification
type PlayerMinScore struct {
	minScore int
}

func (p *PlayerMinScore) Rule() string {
	return "points > :min_score"
}

func (p *PlayerMinScore) Params() interpreter.P {
	return interpreter.P{
		"min_score": p.minScore,
	}
}

// GroupSpec Mock specification for positional param
type GroupSpec struct {
	group string
}

func (g *GroupSpec) Rule() string {
	return "group = :group"
}

func (g *GroupSpec) Params() interpreter.P {
	return interpreter.P{
		"group": g.group,
	}
}

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

func TestSatisfiesSpec(t *testing.T) {

	ruler := New(interpreter.O{})
	isFemale := &IsFemale{}
	minScore := &PlayerMinScore{30}

	spec := spec.AndX([]spec.Specification{
		isFemale, minScore,
	})

	players := []struct {
		Target interpreter.T
		Result bool
	}{
		{interpreter.T{"pseudo": "Joe", "gender": "M", "points": 40}, false},
		{interpreter.T{"pseudo": "Moe", "gender": "M", "points": 20}, false},
		{interpreter.T{"pseudo": "Alice", "gender": "F", "points": 20}, false},
		{interpreter.T{"pseudo": "Birdie", "gender": "F", "points": 60}, true},
	}

	for i, p := range players {
		ret, err := ruler.SatisfiesSpec(p.Target, spec)
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

func TestFilterSpec(t *testing.T) {

	ruler := New(interpreter.O{})
	isFemale := &IsFemale{}
	minScore := &PlayerMinScore{30}

	spec := spec.OrX([]spec.Specification{
		isFemale, minScore,
	})

	users := []interpreter.T{
		{"pseudo": "Joe", "gender": "M", "points": 40},
		{"pseudo": "Moe", "gender": "M", "points": 20},
		{"pseudo": "Alice", "gender": "F", "points": 60},
		{"pseudo": "Birdie", "gender": "F", "points": 20},
	}

	remainder, err := ruler.FilterSpec(users, spec)
	if err != nil {
		t.Errorf("filter error: %v", err)
	}

	if len(remainder) != 3 {
		t.Errorf("remainder's count is not 3, actually, it's %d.", len(remainder))
	}
}

func TestPositionalParam(t *testing.T) {
	ruler := New(interpreter.O{})
	groupOneSpec := &GroupSpec{group: "group one"}
	groupTwoSpec := &GroupSpec{group: "group two"}

	spec := spec.OrX([]spec.Specification{
		groupOneSpec, groupTwoSpec,
	})

	data := []interpreter.T{
		{"name": "test 01", "group": "group one"},
		{"name": "test 02", "group": "group two"},
		{"name": "test 03", "group": "group three"},
	}

	remainder, err := ruler.FilterSpec(data, spec)
	if err != nil {
		t.Errorf("filter error: %v", err)
	}

	if len(remainder) != 2 {
		t.Errorf("remainder's count is not 2, actually, it's %d.", len(remainder))
	}
}
