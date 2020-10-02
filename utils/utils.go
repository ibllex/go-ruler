package utils

// IsAlnum determine whether a character is a letter or a digit
func IsAlnum(ch byte) bool {
	return IsAlpha(ch) || IsDigit(ch)
}

// IsAlpha determine whether a character is a letter
func IsAlpha(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

// IsDigit determine whether a character is a digit
func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// InArray returns a slice contains given value or not
func InArray(v interface{}, arr []interface{}) bool {

	for _, item := range arr {
		if v == item {
			return true
		}
	}

	return false
}

// QueryInMapInter query value in map[string]interface{} by path
func QueryInMapInter(data map[string]interface{}, path []string) interface{} {

	var v interface{} = nil

	for _, k := range path {

		if v == nil {
			v = data[k]
		} else if m, ok := v.(map[string]interface{}); ok {
			v = m[k]
		} else {
			v = nil
			break
		}
	}

	return v
}
