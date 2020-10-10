package spec

import (
	"reflect"
	"testing"

	"github.com/ibllex/go-ruler/interpreter"
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

func (s *IsFemale) PositionalParams() interpreter.PP {
	return interpreter.PP{}
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

func (p *PlayerMinScore) PositionalParams() interpreter.PP {
	return interpreter.PP{}
}

func assertCompositeResult(t *testing.T, spec Specification, rule string, params interpreter.P) {

	t.Helper()

	if rule != spec.Rule() {
		t.Fatalf("expected rule: %s, but %v got.", rule, spec.Rule())
	}

	if !reflect.DeepEqual(params, spec.Params()) {
		t.Fatalf("expected params: %s, but %v got.", params, spec.Params())
	}
}

func TestComposite(t *testing.T) {
	isFemale := &IsFemale{}
	minScore := &PlayerMinScore{10}

	params := interpreter.P{
		"min_score": 10,
	}

	spec := AndX([]Specification{
		isFemale, minScore,
	})
	assertCompositeResult(t, spec, "(gender = 'F') and (points > :min_score)", params)

	spec = OrX([]Specification{
		isFemale, minScore,
	})
	assertCompositeResult(t, spec, "(gender = 'F') or (points > :min_score)", params)
}

func TestNotSpec(t *testing.T) {
	minScore := &PlayerMinScore{10}
	params := interpreter.P{
		"min_score": 10,
	}

	spec := Not(minScore)
	assertCompositeResult(t, spec, "not(points > :min_score)", params)
}
