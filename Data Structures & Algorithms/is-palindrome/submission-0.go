func isPalindrome(s string) bool {
	f := 0
	l := len(s)-1
	for f < l {
		if !(unicode.IsLetter(rune(s[f])) || unicode.IsNumber(rune(s[f]))){
			f++
			continue
		} 
		if !(unicode.IsLetter(rune(s[l])) || unicode.IsNumber(rune(s[l]))){
			l--
			continue
		} 
		if unicode.ToLower(rune(s[f])) == unicode.ToLower(rune(s[l])) {
			f++
			l--
		}else {
			return false
		}

	}
	return true
}