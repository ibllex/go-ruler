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
