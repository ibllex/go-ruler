package operator

import "github.com/ibllex/go-ruler/object"

// Operator func type
type Operator func(args []object.Object) object.Object

// Empty detects if a given object is empty
func Empty(args []object.Object) object.Object {
	v := true
	if len(args) > 0 {
		v = object.IsEmpty(args[0])
	}
	return &object.Boolean{Value: v}
}

// InArray detects if an array contains a specified object
func InArray(args []object.Object) object.Object {
	v := false
	if len(args) >= 2 {
		arr, _ := args[0].Cast(object.ARRAY).(*object.Array)
		target := args[1]

		for _, i := range arr.Elements {
			if i.Equals(target) {
				v = true
				break
			}
		}
	}

	return &object.Boolean{Value: v}
}
