package object

import (
	"testing"
)

type NativeCastData struct {
	Object Object
	Native interface{}
}

func TestToNativeBool(t *testing.T) {
	values := []NativeCastData{
		{&String{"0"}, false},
		{&String{""}, false},
		{&String{"1"}, true},

		{&Integer{1}, true},
		{&Integer{0}, false},
		{&Integer{-1}, true},

		{&Float{1}, true},
		{&Float{0}, false},
		{&Float{-1}, true},

		{&Boolean{true}, true},
		{&Boolean{false}, false},
	}

	for i, v := range values {
		if ToNativeBool(v.Object) != v.Native {
			t.Errorf("values[%d] to native not match expected", i)
		}
	}
}

func TestToNativeInt64(t *testing.T) {
	values := []NativeCastData{
		{&String{"0"}, int64(0)},
		{&String{""}, int64(0)},
		{&String{"1"}, int64(1)},
		{&String{"-1"}, int64(-1)},
		{&String{"1csac"}, int64(0)},

		{&Integer{1}, int64(1)},
		{&Integer{0}, int64(0)},
		{&Integer{-1}, int64(-1)},

		{&Float{1.5}, int64(1)},
		{&Float{1.4}, int64(1)},
		{&Float{0}, int64(0)},

		{&Boolean{true}, int64(1)},
		{&Boolean{false}, int64(0)},
	}

	for i, v := range values {
		if ToNativeInt64(v.Object) != v.Native {
			t.Errorf("values[%d] to native not match expected", i)
		}
	}
}

func TestToNativeFloat64(t *testing.T) {
	values := []NativeCastData{
		{&String{"0.0"}, float64(0)},
		{&String{""}, float64(0)},
		{&String{"1.4"}, float64(1.4)},
		{&String{"-1.4"}, float64(-1.4)},
		{&String{"1.3csac"}, float64(0)},

		{&Integer{1}, float64(1)},
		{&Integer{0}, float64(0)},
		{&Integer{-1}, float64(-1)},

		{&Float{1.5}, float64(1.5)},
		{&Float{1.4}, float64(1.4)},
		{&Float{0}, float64(0)},

		{&Boolean{true}, float64(1)},
		{&Boolean{false}, float64(0)},
	}

	for i, v := range values {
		if ToNativeFloat64(v.Object) != v.Native {
			t.Errorf("values[%d] to native not match expected", i)
		}
	}
}

func TestToNativeString(t *testing.T) {
	values := []NativeCastData{
		{&String{"hello world"}, "hello world"},
		{&String{""}, ""},

		{&Integer{1}, "1"},
		{&Integer{0}, "0"},
		{&Integer{-1}, "-1"},

		{&Float{1.5}, "1.5"},
		{&Float{1.4}, "1.4"},
		{&Float{0}, "0"},

		{&Boolean{true}, "true"},
		{&Boolean{false}, "false"},
	}

	for i, v := range values {
		if ToNativeString(v.Object) != v.Native {
			t.Errorf("values[%d] to native not match expected", i)
		}
	}
}

func TestToNativeArray(t *testing.T) {
	values := []NativeCastData{
		{&Array{[]Object{&String{"hello world"}}}, []Object{&String{"hello world"}}},
	}

	for i, v := range values {
		got := ToNativeArray(v.Object)
		want := v.Native.([]Object)

		if len(got) != len(want) {
			t.Errorf("values[%d] to native not match expected", i)
			continue
		}

		for j, item := range got {
			if !item.Equals(want[j]) {
				t.Errorf("values[%d] to native not match expected", i)
				break
			}
		}
	}
}
