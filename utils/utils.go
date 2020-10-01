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
