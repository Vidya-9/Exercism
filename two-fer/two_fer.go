package twofer

//ShareWith function returns the name
func ShareWith(x string) string {
	lead := "One for "
	end := ", one for me."
	mid := ""
	if x == "" {
		mid = "you"
	} else {
		mid = x
	}
	return lead + mid + end
}
