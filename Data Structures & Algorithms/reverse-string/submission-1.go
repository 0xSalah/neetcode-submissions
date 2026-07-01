func reverseString(s []byte) {
	f := 0
	l := len(s) -1 

	for f < l {
		if s[f] == s[l]{
			f++
			l--
			continue
		}else {
		s[f],s[l] = s[l],s[f]
		f++
		l--
		}
	}
}
