package object

// ToNativeBool convert object to native bool value
func ToNativeBool(o Object) bool {
	return o.Cast(BOOLEAN).(*Boolean).Value
}

// ToNativeInt64 convert object to native int64 value
func ToNativeInt64(o Object) int64 {
	return o.Cast(INTEGER).(*Integer).Value
}

// ToNativeFloat64 convert object to native float64 value
func ToNativeFloat64(o Object) float64 {
	return o.Cast(FLOAT).(*Float).Value
}

// ToNativeString convert object to native string value
func ToNativeString(o Object) string {
	return o.Cast(STRING).(*String).Value
}

// ToNativeArray convert object to native Object slice
func ToNativeArray(o Object) []Object {
	return o.Cast(ARRAY).(*Array).Elements
}

// NativeToObject convert native interface to Object
func NativeToObject(v interface{}) Object {

	switch v := v.(type) {
	case int:
		return &Integer{Value: int64(v)}
	case int8:
		return &Integer{Value: int64(v)}
	case int16:
		return &Integer{Value: int64(v)}
	case int32:
		return &Integer{Value: int64(v)}
	case int64:
		return &Integer{Value: v}
	case float32:
		return &Float{Value: float64(v)}
	case float64:
		return &Float{Value: v}
	case string:
		return &String{Value: v}
	case bool:
		return &Boolean{Value: v}
	case []interface{}:
		els := []Object{}
		for _, e := range v {
			els = append(els, NativeToObject(e))
		}
		return &Array{Elements: els}
	}

	return &Null{}
}
