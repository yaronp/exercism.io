package leap

// IsLeapYear return true if a year is leaf. See https://tinyurl.com/y5m2wa9h for more info
func IsLeapYear(year int) bool {
	if (year%400 == 0) || ((year%4 == 0) && !(year%100 == 0)) {
		return true
	}
	return false
}
