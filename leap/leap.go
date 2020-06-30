package leap

//Definition of the function IsLeapYear
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
	}
	if year%100 != 0 {
		if year%4 == 0 {
			return true
		}
	}
	return false
}
