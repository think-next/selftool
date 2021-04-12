package uniq

/*
	去重，是否已经出现过
*/

func IsDuplicateForStrings(s []string, aim string) bool {
	for _, v := range s {
		if v == aim {
			return true
		}
	}
	return false
}
