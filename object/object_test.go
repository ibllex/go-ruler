package object

import "testing"

type TypeCastData struct {
	Value  Object
	Cast   Type
	Result Object
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("trues do not have same hash key")
	}

	if false1.HashKey() != false2.HashKey() {
		t.Errorf("falses do not have same hash key")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("true has same hash key as false")
	}
}

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "Good job"}
	diff2 := &String{Value: "Good job"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	one1 := &Integer{Value: 1}
	one2 := &Integer{Value: 1}
	two1 := &Integer{Value: 2}
	two2 := &Integer{Value: 2}

	if one1.HashKey() != one2.HashKey() {
		t.Errorf("integers with same content have twoerent hash keys")
	}

	if two1.HashKey() != two2.HashKey() {
		t.Errorf("integers with same content have twoerent hash keys")
	}

	if one1.HashKey() == two1.HashKey() {
		t.Errorf("integers with twoerent content have same hash keys")
	}
}

func TestFloatHashKey(t *testing.T) {
	one1 := &Float{Value: 1.4}
	one2 := &Float{Value: 1.4}
	two1 := &Float{Value: 2.4}
	two2 := &Float{Value: 2.4}

	if one1.HashKey() != one2.HashKey() {
		t.Errorf("decimals with same content have twoerent hash keys")
	}

	if two1.HashKey() != two2.HashKey() {
		t.Errorf("decimals with same content have twoerent hash keys")
	}

	if one1.HashKey() == two1.HashKey() {
		t.Errorf("decimals with twoerent content have same hash keys")
	}
}

func TestFloatCast(t *testing.T) {
	values := []TypeCastData{
		{&Float{1.50}, STRING, &String{"1.5"}},
		{&Float{1.5}, STRING, &String{"1.5"}},
		{&Float{0.02}, STRING, &String{"0.02"}},
		{&Float{0.0}, STRING, &String{"0"}},

		{&Float{1.4}, INTEGER, &Integer{1}},
		{&Float{1.8}, INTEGER, &Integer{1}},

		{&Float{1.40}, FLOAT, &Float{1.4}},
		{&Float{1.40}, FLOAT, &Float{1.40}},
		{&Float{1.8}, FLOAT, &Float{1.8}},

		{&Float{0}, BOOLEAN, &Boolean{false}},
		{&Float{10.0}, BOOLEAN, &Boolean{true}},
		{&Float{-10.0}, BOOLEAN, &Boolean{true}},
	}

	assertObjectsCast(t, values)
}

func TestIntegerCast(t *testing.T) {

	values := []TypeCastData{
		{&Integer{10}, STRING, &String{"10"}},
		{&Integer{0}, STRING, &String{"0"}},
		{&Integer{-10}, STRING, &String{"-10"}},

		{&Integer{10}, INTEGER, &Integer{10}},
		{&Integer{0}, INTEGER, &Integer{0}},
		{&Integer{-10}, INTEGER, &Integer{-10}},

		{&Integer{10}, FLOAT, &Float{10.0}},
		{&Integer{0}, FLOAT, &Float{0.0}},
		{&Integer{-10}, FLOAT, &Float{-10.0}},

		{&Integer{10}, BOOLEAN, &Boolean{true}},
		{&Integer{0}, BOOLEAN, &Boolean{false}},
		{&Integer{-10}, BOOLEAN, &Boolean{true}},
	}

	assertObjectsCast(t, values)
}

func TestBooleanCast(t *testing.T) {

	values := []TypeCastData{
		{&Boolean{true}, STRING, &String{"true"}},
		{&Boolean{false}, STRING, &String{"false"}},

		{&Boolean{true}, INTEGER, &Integer{1}},
		{&Boolean{false}, INTEGER, &Integer{0}},

		{&Boolean{true}, FLOAT, &Float{1.0}},
		{&Boolean{false}, FLOAT, &Float{0.0}},

		{&Boolean{true}, BOOLEAN, &Boolean{true}},
		{&Boolean{false}, BOOLEAN, &Boolean{false}},
	}

	assertObjectsCast(t, values)
}

func TestStringCast(t *testing.T) {

	values := []TypeCastData{
		{&String{""}, STRING, &String{""}},
		{&String{"hello world"}, STRING, &String{"hello world"}},

		{&String{"10"}, INTEGER, &Integer{10}},
		{&String{""}, INTEGER, &Integer{0}},
		{&String{"10abc"}, INTEGER, &Integer{0}},
		{&String{"abcdef10"}, INTEGER, &Integer{0}},

		{&String{"10.05"}, FLOAT, &Float{10.05}},
		{&String{""}, FLOAT, &Float{0.0}},
		{&String{"10.05abc"}, FLOAT, &Float{0.0}},
		{&String{"abcdef10.0"}, FLOAT, &Float{0.0}},

		{&String{"abcdef"}, BOOLEAN, &Boolean{true}},
		{&String{"1"}, BOOLEAN, &Boolean{true}},
		{&String{"123"}, BOOLEAN, &Boolean{true}},
		{&String{"null"}, BOOLEAN, &Boolean{true}},
		{&String{""}, BOOLEAN, &Boolean{false}},
		{&String{"0"}, BOOLEAN, &Boolean{false}},
	}

	assertObjectsCast(t, values)
}

func TestNullCast(t *testing.T) {

	values := []TypeCastData{
		{&Null{}, STRING, &String{""}},
		{&Null{}, INTEGER, &Integer{0}},
		{&Null{}, FLOAT, &Float{0}},
		{&Null{}, BOOLEAN, &Boolean{false}},
	}

	assertObjectsCast(t, values)
}

func assertObjectsCast(t *testing.T, values []TypeCastData) {
	t.Helper()

	for i, v := range values {
		newObject := v.Value.Cast(v.Cast)

		if newObject.(HashTable).HashKey() != v.Result.(HashTable).HashKey() {
			t.Errorf("values[%d] cast failed, expected %v but got %v.", i, v.Result, newObject)
		}
	}
}
